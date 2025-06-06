package replication

import (
	"container/list"
	"fmt"
	"sync"
	"time"

	replicationspb "go.temporal.io/server/api/replication/v1"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/log/tag"
	"go.temporal.io/server/common/metrics"
	ctasks "go.temporal.io/server/common/tasks"
)

//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination executable_task_tracker_mock.go

const MarkPoisonPillMaxAttempts = 3

type (
	TrackableExecutableTask interface {
		ctasks.Task
		QueueID() interface{}
		TaskID() int64
		TaskCreationTime() time.Time
		MarkPoisonPill() error
		SourceClusterName() string
		ReplicationTask() *replicationspb.ReplicationTask
	}
	WatermarkInfo struct {
		Watermark int64
		Timestamp time.Time
	}
	ExecutableTaskTracker interface {
		TrackTasks(exclusiveHighWatermarkInfo WatermarkInfo, tasks ...TrackableExecutableTask) []TrackableExecutableTask
		LowWatermark() *WatermarkInfo
		Size() int
		Cancel()
	}
	ExecutableTaskTrackerImpl struct {
		logger         log.Logger
		metricsHandler metrics.Handler

		sync.Mutex
		cancelled                  bool
		exclusiveHighWatermarkInfo *WatermarkInfo // this is exclusive, i.e. source need to resend with this watermark / task ID
		taskQueue                  *list.List     // sorted by task ID
	}
)

var _ ExecutableTaskTracker = (*ExecutableTaskTrackerImpl)(nil)

func NewExecutableTaskTracker(
	logger log.Logger,
	metricsHandler metrics.Handler,
) *ExecutableTaskTrackerImpl {
	return &ExecutableTaskTrackerImpl{
		logger:         logger,
		metricsHandler: metricsHandler,

		exclusiveHighWatermarkInfo: nil,
		taskQueue:                  list.New(),
	}
}

// TrackTasks add tasks for tracking, return valid tasks (dedup)
// if task tracker is cancelled, then newly added tasks will also be cancelled
// tasks should be sorted by task ID, all task IDs < exclusiveHighWatermarkInfo
func (t *ExecutableTaskTrackerImpl) TrackTasks(
	exclusiveHighWatermarkInfo WatermarkInfo,
	tasks ...TrackableExecutableTask,
) []TrackableExecutableTask {
	filteredTasks := make([]TrackableExecutableTask, 0, len(tasks))

	t.Lock()
	defer t.Unlock()

	// need to assume source side send replication tasks in order
	if t.exclusiveHighWatermarkInfo != nil && exclusiveHighWatermarkInfo.Watermark <= t.exclusiveHighWatermarkInfo.Watermark {
		return filteredTasks
	}

	lastTaskID := int64(-1)
	if item := t.taskQueue.Back(); item != nil {
		lastTaskID = item.Value.(TrackableExecutableTask).TaskID()
	}
Loop:
	for _, task := range tasks {
		if lastTaskID >= task.TaskID() {
			// need to assume source side send replication tasks in order
			continue Loop
		}
		t.taskQueue.PushBack(task)
		filteredTasks = append(filteredTasks, task)
		lastTaskID = task.TaskID()
	}

	if exclusiveHighWatermarkInfo.Watermark <= lastTaskID {
		panic(fmt.Sprintf(
			"ExecutableTaskTracker encountered lower high watermark: %v < %v",
			exclusiveHighWatermarkInfo.Watermark,
			lastTaskID,
		))
	}
	t.exclusiveHighWatermarkInfo = &exclusiveHighWatermarkInfo

	if t.cancelled {
		t.cancelLocked()
	}
	return filteredTasks
}

func (t *ExecutableTaskTrackerImpl) LowWatermark() *WatermarkInfo {
	t.Lock()
	defer t.Unlock()

	element := t.taskQueue.Front()
Loop:
	for element != nil {
		task := element.Value.(TrackableExecutableTask)
		taskState := task.State()
		switch taskState {
		case ctasks.TaskStateAcked:
			nextElement := element.Next()
			t.taskQueue.Remove(element)
			element = nextElement
		case ctasks.TaskStateNacked:
			if err := task.MarkPoisonPill(); err != nil {
				t.logger.Error("unable to save poison pill", tag.Error(err), tag.TaskID(task.TaskID()))
				metrics.ReplicationDLQFailed.With(t.metricsHandler).Record(
					1,
					metrics.OperationTag(metrics.ReplicationTaskTrackerScope),
				)
				// unable to save poison pill, retry later
				element = element.Next()
				continue Loop
			}
			nextElement := element.Next()
			t.taskQueue.Remove(element)
			element = nextElement
		case ctasks.TaskStateAborted:
			// noop, do not remove from queue, let it block low watermark
			element = element.Next()
		case ctasks.TaskStateCancelled:
			// noop, do not remove from queue, let it block low watermark
			element = element.Next()
		case ctasks.TaskStatePending:
			// noop, do not remove from queue, let it block low watermark
			element = element.Next()
		default:
			panic(fmt.Sprintf(
				"ExecutableTaskTracker encountered unknown task state: %v",
				taskState,
			))
		}
	}

	if element := t.taskQueue.Front(); element != nil {
		inclusiveLowWatermarkInfo := WatermarkInfo{
			Watermark: element.Value.(TrackableExecutableTask).TaskID(),
			Timestamp: element.Value.(TrackableExecutableTask).TaskCreationTime(),
		}
		return &inclusiveLowWatermarkInfo
	} else if t.exclusiveHighWatermarkInfo != nil {
		inclusiveLowWatermarkInfo := *t.exclusiveHighWatermarkInfo
		return &inclusiveLowWatermarkInfo
	} else {
		return nil
	}
}

func (t *ExecutableTaskTrackerImpl) Size() int {
	t.Lock()
	defer t.Unlock()

	return t.taskQueue.Len()
}

func (t *ExecutableTaskTrackerImpl) Cancel() {
	t.Lock()
	defer t.Unlock()

	t.cancelled = true
	t.cancelLocked()
}

func (t *ExecutableTaskTrackerImpl) cancelLocked() {
	for element := t.taskQueue.Front(); element != nil; element = element.Next() {
		task := element.Value.(TrackableExecutableTask)
		task.Cancel()
	}
}
