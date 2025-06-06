package scheduler

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/dgryski/go-farm"
	schedulepb "go.temporal.io/api/schedule/v1"
	"go.temporal.io/server/common"
	"go.temporal.io/server/common/cache"
	"go.temporal.io/server/common/primitives/timestamp"
	"go.temporal.io/server/common/util"
)

type (
	CompiledSpec struct {
		spec     *schedulepb.ScheduleSpec
		tz       *time.Location
		calendar []*compiledCalendar
		excludes []*compiledCalendar
	}

	GetNextTimeResult struct {
		Nominal time.Time // scheduled time before adding jitter
		Next    time.Time // scheduled time after adding jitter
	}

	SpecBuilder struct {
		// locationCache is a cache for the results of time.LoadLocation. That function accesses
		// the filesystem and is relatively slow. We assume that it returns a semantically
		// equivalent value for the same location name. This isn't strictly true, for example if
		// the time zone database is changed while the process is running. To handle that, we
		// expire entries after a day. Note that we cache negative results also.
		locationCache cache.Cache
	}

	locationAndError struct {
		loc *time.Location
		err error
	}
)

func NewSpecBuilder() *SpecBuilder {
	return &SpecBuilder{
		locationCache: cache.New(1000,
			&cache.Options{
				TTL: 24 * time.Hour,
			},
		),
	}
}

func (b *SpecBuilder) NewCompiledSpec(spec *schedulepb.ScheduleSpec) (*CompiledSpec, error) {
	spec, err := canonicalizeSpec(spec)
	if err != nil {
		return nil, err
	}

	// load timezone
	tz, err := b.loadTimezone(spec)
	if err != nil {
		return nil, err
	}

	// compile StructuredCalendarSpecs
	ccs := make([]*compiledCalendar, len(spec.StructuredCalendar))
	for i, structured := range spec.StructuredCalendar {
		ccs[i] = newCompiledCalendar(structured, tz)
	}

	// compile excludes
	excludes := make([]*compiledCalendar, len(spec.ExcludeStructuredCalendar))
	for i, excal := range spec.ExcludeStructuredCalendar {
		excludes[i] = newCompiledCalendar(excal, tz)
	}

	cspec := &CompiledSpec{
		spec:     spec,
		tz:       tz,
		calendar: ccs,
		excludes: excludes,
	}

	return cspec, nil
}

// CleanSpec sets default values in ranges.
func CleanSpec(spec *schedulepb.ScheduleSpec) {
	cleanRanges := func(ranges []*schedulepb.Range) {
		for _, r := range ranges {
			if r.End < r.Start {
				r.End = r.Start
			}
			if r.Step == 0 {
				r.Step = 1
			}
		}
	}
	cleanCal := func(structured *schedulepb.StructuredCalendarSpec) {
		cleanRanges(structured.Second)
		cleanRanges(structured.Minute)
		cleanRanges(structured.Hour)
		cleanRanges(structured.DayOfMonth)
		cleanRanges(structured.Month)
		cleanRanges(structured.Year)
		cleanRanges(structured.DayOfWeek)
	}
	for _, structured := range spec.StructuredCalendar {
		cleanCal(structured)
	}
	for _, structured := range spec.ExcludeStructuredCalendar {
		cleanCal(structured)
	}
}

//revive:disable-next-line:cognitive-complexity
func canonicalizeSpec(spec *schedulepb.ScheduleSpec) (*schedulepb.ScheduleSpec, error) {
	// make copy so we can change some fields
	spec = common.CloneProto(spec)

	// parse CalendarSpecs to StructuredCalendarSpecs
	for _, cal := range spec.Calendar {
		structured, err := parseCalendarToStructured(cal)
		if err != nil {
			return nil, err
		}
		spec.StructuredCalendar = append(spec.StructuredCalendar, structured)
	}
	spec.Calendar = nil

	// parse ExcludeCalendars
	for _, cal := range spec.ExcludeCalendar {
		structured, err := parseCalendarToStructured(cal)
		if err != nil {
			return nil, err
		}
		spec.ExcludeStructuredCalendar = append(spec.ExcludeStructuredCalendar, structured)
	}
	spec.ExcludeCalendar = nil

	// parse CronStrings
	const unset = "__unset__"
	cronTZ := unset
	for _, cs := range spec.CronString {
		structured, interval, tz, err := parseCronString(cs)
		if err != nil {
			return nil, err
		}
		if cronTZ != unset && tz != cronTZ {
			// all cron strings must agree on timezone (whether present or not)
			return nil, errConflictingTimezoneNames
		}
		cronTZ = tz
		if structured != nil {
			spec.StructuredCalendar = append(spec.StructuredCalendar, structured)
		}
		if interval != nil {
			spec.Interval = append(spec.Interval, interval)
		}
	}
	spec.CronString = nil

	// if we have cron string(s), copy the timezone to spec, checking for conflict first.
	// if cron string timezone is empty string, don't copy, let the one in spec be used.
	if cronTZ != unset && cronTZ != "" {
		if spec.TimezoneName != "" && spec.TimezoneName != cronTZ || spec.TimezoneData != nil {
			return nil, errConflictingTimezoneNames
		} else if spec.TimezoneName == "" {
			spec.TimezoneName = cronTZ
		}
	}

	// validate structured calendar
	for _, structured := range spec.StructuredCalendar {
		if err := validateStructuredCalendar(structured); err != nil {
			return nil, err
		}
	}

	// validate intervals
	for _, interval := range spec.Interval {
		if err := validateInterval(interval); err != nil {
			return nil, err
		}
	}

	return spec, nil
}

func validateStructuredCalendar(scs *schedulepb.StructuredCalendarSpec) error {
	var errs []string

	checkRanges := func(ranges []*schedulepb.Range, field string, minVal, maxVal int32) {
		for _, r := range ranges {
			if r == nil { // shouldn't happen
				errs = append(errs, "range is nil")
				continue
			}
			if r.Start < minVal || r.Start > maxVal {
				errs = append(errs, fmt.Sprintf("%s Start is not in range [%d-%d]", field, minVal, maxVal))
			}
			if r.End != 0 && (r.End < r.Start || r.End > maxVal) {
				errs = append(errs, fmt.Sprintf("%s End is before Start or not in range [%d-%d]", field, minVal, maxVal))
			}
			if r.Step < 0 {
				errs = append(errs, fmt.Sprintf("%s has invalid Step", field))
			}
		}
	}

	checkRanges(scs.Second, "Second", 0, 59)
	checkRanges(scs.Minute, "Minute", 0, 59)
	checkRanges(scs.Hour, "Hour", 0, 23)
	checkRanges(scs.DayOfMonth, "DayOfMonth", 1, 31)
	checkRanges(scs.Month, "Month", 1, 12)
	checkRanges(scs.Year, "Year", minCalendarYear, maxCalendarYear)
	checkRanges(scs.DayOfWeek, "DayOfWeek", 0, 6)

	if len(scs.Comment) > maxCommentLen {
		errs = append(errs, "comment is too long")
	}

	if len(errs) > 0 {
		return errors.New("invalid calendar spec: " + strings.Join(errs, ", "))
	}
	return nil
}

func validateInterval(i *schedulepb.IntervalSpec) error {
	if i == nil {
		return errors.New("interval is nil")
	}
	// TODO: use timestamp.ValidateAndCapProtoDuration after switching to state machine based implementation.
	// 	Not adding it to workflow based implementation to avoid potential non-determinism errors.
	iv, phase := timestamp.DurationValue(i.Interval), timestamp.DurationValue(i.Phase)
	if iv < time.Second {
		return errors.New("interval is too small")
	} else if phase < 0 {
		return errors.New("phase is negative")
	} else if phase >= iv {
		return errors.New("phase cannot be greater than Interval")
	}
	return nil
}

func (b *SpecBuilder) loadTimezone(spec *schedulepb.ScheduleSpec) (*time.Location, error) {
	if spec.TimezoneData != nil {
		return time.LoadLocationFromTZData(spec.TimezoneName, spec.TimezoneData)
	}

	if cached, ok := b.locationCache.Get(spec.TimezoneName).(*locationAndError); ok {
		return cached.loc, cached.err
	}
	loc, err := time.LoadLocation(spec.TimezoneName)
	b.locationCache.Put(spec.TimezoneName, &locationAndError{
		loc: loc,
		err: err,
	})
	return loc, err
}

func (cs *CompiledSpec) CanonicalForm() *schedulepb.ScheduleSpec {
	return cs.spec
}

// Returns the earliest time that matches the schedule spec that is after the given time.
// Returns: Nominal is the time that matches, pre-jitter. Next is the nominal time with
// jitter applied. If there is no matching time, Nominal and Next will be the zero time.
func (cs *CompiledSpec) GetNextTime(jitterSeed string, after time.Time) GetNextTimeResult {
	// If we're starting before the schedule's allowed time range, jump up to right before
	// it (so that we can still return the first second of the range if it happens to match).
	// note: AsTime returns unix epoch on nil StartTime
	after = util.MaxTime(after, cs.spec.StartTime.AsTime().Add(-time.Second))

	pastEndTime := func(t time.Time) bool {
		return cs.spec.EndTime != nil && t.After(cs.spec.EndTime.AsTime()) || t.Year() > maxCalendarYear
	}
	var nominal time.Time
	for nominal.IsZero() || cs.excluded(nominal) {
		nominal = cs.rawNextTime(after)
		after = nominal

		if nominal.IsZero() || pastEndTime(nominal) {
			return GetNextTimeResult{}
		}
	}

	maxJitter := timestamp.DurationValue(cs.spec.Jitter)
	// Ensure that jitter doesn't push this time past the _next_ nominal start time
	if following := cs.rawNextTime(nominal); !following.IsZero() {
		maxJitter = min(maxJitter, following.Sub(nominal))
	}
	next := cs.addJitter(jitterSeed, nominal, maxJitter)

	return GetNextTimeResult{Nominal: nominal, Next: next}
}

// Returns the next matching time (without jitter), or the zero value if no time matches.
func (cs *CompiledSpec) rawNextTime(after time.Time) (nominal time.Time) {
	var minTimestamp int64 = math.MaxInt64 // unix seconds-since-epoch as int64

	for _, cal := range cs.calendar {
		if next := cal.next(after); !next.IsZero() {
			nextTs := next.Unix()
			if nextTs < minTimestamp {
				minTimestamp = nextTs
			}
		}
	}

	ts := after.Unix()
	for _, iv := range cs.spec.Interval {
		next := cs.nextIntervalTime(iv, ts)
		if next < minTimestamp {
			minTimestamp = next
		}
	}

	if minTimestamp == math.MaxInt64 {
		return time.Time{}
	}
	return time.Unix(minTimestamp, 0).UTC()
}

// Returns the next matching time for a single interval spec.
func (cs *CompiledSpec) nextIntervalTime(iv *schedulepb.IntervalSpec, ts int64) int64 {
	interval := int64(timestamp.DurationValue(iv.Interval) / time.Second)
	if interval < 1 {
		interval = 1
	}
	phase := int64(timestamp.DurationValue(iv.Phase) / time.Second)
	if phase < 0 {
		phase = 0
	}
	return (((ts-phase)/interval)+1)*interval + phase
}

// Returns true if any exclude spec matches the time.
func (cs *CompiledSpec) excluded(nominal time.Time) bool {
	for _, excal := range cs.excludes {
		if excal.matches(nominal) {
			return true
		}
	}
	return false
}

// Adds jitter to a nominal time, deterministically (by hashing the given time and a seed).
func (cs *CompiledSpec) addJitter(seed string, nominal time.Time, maxJitter time.Duration) time.Time {
	if maxJitter < 0 {
		maxJitter = 0
	}

	bin, err := nominal.MarshalBinary()
	if err != nil {
		return nominal
	}

	bin = append(bin, []byte(seed)...)

	// we want to fit the result of a multiply in 64 bits, and use 32 bits of hash, which
	// leaves 32 bits for the range. if we use nanoseconds or microseconds, our range is
	// limited to only a few seconds or hours. using milliseconds supports up to 49 days.
	fp := uint64(farm.Fingerprint32(bin))
	ms := uint64(maxJitter.Milliseconds())
	if ms > math.MaxUint32 {
		ms = math.MaxUint32
	}
	jitter := time.Duration((fp*ms)>>32) * time.Millisecond
	return nominal.Add(jitter)
}
