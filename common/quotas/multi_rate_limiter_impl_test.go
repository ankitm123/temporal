package quotas

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type (
	multiStageRateLimiterSuite struct {
		suite.Suite
		*require.Assertions

		controller        *gomock.Controller
		firstRateLimiter  *MockRateLimiter
		secondRateLimiter *MockRateLimiter
		firstReservation  *MockReservation
		secondReservation *MockReservation

		rateLimiter *MultiRateLimiterImpl
	}
)

func TestMultiStageRateLimiterSuite(t *testing.T) {
	s := new(multiStageRateLimiterSuite)
	suite.Run(t, s)
}

func (s *multiStageRateLimiterSuite) SetupSuite() {

}

func (s *multiStageRateLimiterSuite) TearDownSuite() {

}

func (s *multiStageRateLimiterSuite) SetupTest() {
	s.Assertions = require.New(s.T())

	s.controller = gomock.NewController(s.T())
	s.firstRateLimiter = NewMockRateLimiter(s.controller)
	s.secondRateLimiter = NewMockRateLimiter(s.controller)
	s.firstReservation = NewMockReservation(s.controller)
	s.secondReservation = NewMockReservation(s.controller)

	s.rateLimiter = NewMultiRateLimiter([]RateLimiter{s.firstRateLimiter, s.secondRateLimiter})
}

func (s *multiStageRateLimiterSuite) TearDownTest() {
	s.controller.Finish()
}

func (s *multiStageRateLimiterSuite) TestAllowN_NonSuccess() {
	now := time.Now()
	numToken := 2

	s.firstReservation.EXPECT().OK().Return(false).AnyTimes()
	s.firstRateLimiter.EXPECT().ReserveN(now, numToken).Return(s.firstReservation)

	result := s.rateLimiter.AllowN(now, numToken)
	s.False(result)
}

func (s *multiStageRateLimiterSuite) TestAllowN_SomeSuccess_Case1() {
	now := time.Now()
	numToken := 2

	s.firstReservation.EXPECT().OK().Return(true).AnyTimes()
	s.firstReservation.EXPECT().DelayFrom(now).Return(time.Duration(0)).AnyTimes()
	s.firstReservation.EXPECT().CancelAt(now)
	s.firstRateLimiter.EXPECT().ReserveN(now, numToken).Return(s.firstReservation)

	s.secondReservation.EXPECT().OK().Return(false).AnyTimes()
	s.secondReservation.EXPECT().DelayFrom(now).Return(time.Duration(0)).AnyTimes()
	s.secondRateLimiter.EXPECT().ReserveN(now, numToken).Return(s.secondReservation)

	result := s.rateLimiter.AllowN(now, numToken)
	s.False(result)
}

func (s *multiStageRateLimiterSuite) TestAllowN_SomeSuccess_Case2() {
	now := time.Now()
	numToken := 2

	s.firstReservation.EXPECT().OK().Return(true).AnyTimes()
	s.firstReservation.EXPECT().DelayFrom(now).Return(time.Duration(0)).AnyTimes()
	s.firstReservation.EXPECT().CancelAt(now)
	s.firstRateLimiter.EXPECT().ReserveN(now, numToken).Return(s.firstReservation)

	s.secondReservation.EXPECT().OK().Return(true).AnyTimes()
	s.secondReservation.EXPECT().DelayFrom(now).Return(time.Duration(1)).AnyTimes()
	s.secondReservation.EXPECT().CancelAt(now)
	s.secondRateLimiter.EXPECT().ReserveN(now, numToken).Return(s.secondReservation)

	result := s.rateLimiter.AllowN(now, numToken)
	s.False(result)
}

func (s *multiStageRateLimiterSuite) TestAllowN_AllSuccess() {
	now := time.Now()
	numToken := 2

	s.firstReservation.EXPECT().OK().Return(true).AnyTimes()
	s.firstReservation.EXPECT().DelayFrom(now).Return(time.Duration(0)).AnyTimes()
	s.firstRateLimiter.EXPECT().ReserveN(now, numToken).Return(s.firstReservation)

	s.secondReservation.EXPECT().OK().Return(true).AnyTimes()
	s.secondReservation.EXPECT().DelayFrom(now).Return(time.Duration(0)).AnyTimes()
	s.secondRateLimiter.EXPECT().ReserveN(now, numToken).Return(s.secondReservation)

	result := s.rateLimiter.AllowN(now, numToken)
	s.True(result)
}

func (s *multiStageRateLimiterSuite) TestReserveN_NonSuccess() {
	now := time.Now()
	numToken := 4

	s.firstReservation.EXPECT().OK().Return(false).AnyTimes()
	s.firstRateLimiter.EXPECT().ReserveN(now, numToken).Return(s.firstReservation)

	result := s.rateLimiter.ReserveN(now, numToken)
	s.Equal(&MultiReservationImpl{
		ok:           false,
		reservations: nil,
	}, result)
}

func (s *multiStageRateLimiterSuite) TestReserveN_SomeSuccess() {
	now := time.Now()
	numToken := 4

	s.firstReservation.EXPECT().OK().Return(true).AnyTimes()
	s.firstReservation.EXPECT().CancelAt(now)
	s.firstRateLimiter.EXPECT().ReserveN(now, numToken).Return(s.firstReservation)

	s.secondReservation.EXPECT().OK().Return(false).AnyTimes()
	s.secondRateLimiter.EXPECT().ReserveN(now, numToken).Return(s.secondReservation)

	result := s.rateLimiter.ReserveN(now, numToken)
	s.Equal(&MultiReservationImpl{
		ok:           false,
		reservations: nil,
	}, result)
}

func (s *multiStageRateLimiterSuite) TestReserveN_AllSuccess() {
	now := time.Now()
	numToken := 4

	s.firstReservation.EXPECT().OK().Return(true).AnyTimes()
	s.firstRateLimiter.EXPECT().ReserveN(now, numToken).Return(s.firstReservation)

	s.secondReservation.EXPECT().OK().Return(true).AnyTimes()
	s.secondRateLimiter.EXPECT().ReserveN(now, numToken).Return(s.secondReservation)

	result := s.rateLimiter.ReserveN(now, numToken)
	s.Equal(&MultiReservationImpl{
		ok:           true,
		reservations: []Reservation{s.firstReservation, s.secondReservation},
	}, result)
}

func (s *multiStageRateLimiterSuite) TestWaitN_AlreadyExpired() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	numToken := 4

	result := s.rateLimiter.WaitN(ctx, numToken)
	s.Error(result)
}

func (s *multiStageRateLimiterSuite) TestWaitN_NotExpired_WithExpiration_Error() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	numToken := 4

	firstReservationDelay := 2 * time.Second
	secondReservationDelay := 3 * time.Second
	s.firstReservation.EXPECT().DelayFrom(gomock.Any()).Return(firstReservationDelay).AnyTimes()
	s.secondReservation.EXPECT().DelayFrom(gomock.Any()).Return(secondReservationDelay).AnyTimes()
	s.firstReservation.EXPECT().CancelAt(gomock.Any())
	s.secondReservation.EXPECT().CancelAt(gomock.Any())

	s.firstReservation.EXPECT().OK().Return(true).AnyTimes()
	s.firstRateLimiter.EXPECT().ReserveN(gomock.Any(), numToken).Return(s.firstReservation)
	s.secondReservation.EXPECT().OK().Return(true).AnyTimes()
	s.secondRateLimiter.EXPECT().ReserveN(gomock.Any(), numToken).Return(s.secondReservation)

	result := s.rateLimiter.WaitN(ctx, numToken)
	s.Error(result)
}

func (s *multiStageRateLimiterSuite) TestWaitN_NotExpired_WithExpiration_Cancelled() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	numToken := 4

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	firstReservationDelay := 20 * time.Second
	secondReservationDelay := 30 * time.Second
	s.firstReservation.EXPECT().DelayFrom(gomock.Any()).Return(firstReservationDelay).AnyTimes()
	s.secondReservation.EXPECT().DelayFrom(gomock.Any()).Return(secondReservationDelay).AnyTimes()
	s.firstReservation.EXPECT().CancelAt(gomock.Any())
	s.secondReservation.EXPECT().CancelAt(gomock.Any())

	s.firstReservation.EXPECT().OK().Return(true).AnyTimes()
	s.firstRateLimiter.EXPECT().ReserveN(gomock.Any(), numToken).Return(s.firstReservation)
	s.secondReservation.EXPECT().OK().Return(true).AnyTimes()
	s.secondRateLimiter.EXPECT().ReserveN(gomock.Any(), numToken).Return(s.secondReservation)

	result := s.rateLimiter.WaitN(ctx, numToken)
	s.Error(result)
}

func (s *multiStageRateLimiterSuite) TestWaitN_NotExpired_WithExpiration_NoError() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	numToken := 4

	firstReservationDelay := 2 * time.Second
	secondReservationDelay := 3 * time.Second
	s.firstReservation.EXPECT().DelayFrom(gomock.Any()).Return(firstReservationDelay).AnyTimes()
	s.secondReservation.EXPECT().DelayFrom(gomock.Any()).Return(secondReservationDelay).AnyTimes()

	s.firstReservation.EXPECT().OK().Return(true).AnyTimes()
	s.firstRateLimiter.EXPECT().ReserveN(gomock.Any(), numToken).Return(s.firstReservation)
	s.secondReservation.EXPECT().OK().Return(true).AnyTimes()
	s.secondRateLimiter.EXPECT().ReserveN(gomock.Any(), numToken).Return(s.secondReservation)

	result := s.rateLimiter.WaitN(ctx, numToken)
	s.NoError(result)
}

func (s *multiStageRateLimiterSuite) TestWaitN_NotExpired_WithoutExpiration() {
	ctx := context.Background()
	numToken := 4

	firstReservationDelay := 2 * time.Second
	secondReservationDelay := 3 * time.Second
	s.firstReservation.EXPECT().DelayFrom(gomock.Any()).Return(firstReservationDelay).AnyTimes()
	s.secondReservation.EXPECT().DelayFrom(gomock.Any()).Return(secondReservationDelay).AnyTimes()

	s.firstReservation.EXPECT().OK().Return(true).AnyTimes()
	s.firstRateLimiter.EXPECT().ReserveN(gomock.Any(), numToken).Return(s.firstReservation)
	s.secondReservation.EXPECT().OK().Return(true).AnyTimes()
	s.secondRateLimiter.EXPECT().ReserveN(gomock.Any(), numToken).Return(s.secondReservation)

	result := s.rateLimiter.WaitN(ctx, numToken)
	s.NoError(result)
}

func (s *multiStageRateLimiterSuite) TestRate() {
	firstRateLimiterRate := float64(10)
	secondRateLimiterRate := float64(5)

	s.firstRateLimiter.EXPECT().Rate().Return(firstRateLimiterRate).AnyTimes()
	s.secondRateLimiter.EXPECT().Rate().Return(secondRateLimiterRate).AnyTimes()

	result := s.rateLimiter.Rate()
	s.Equal(secondRateLimiterRate, result)
}

func (s *multiStageRateLimiterSuite) TestBurst() {
	firstRateLimiterBurst := 5
	secondRateLimiterBurst := 10

	s.firstRateLimiter.EXPECT().Burst().Return(firstRateLimiterBurst).AnyTimes()
	s.secondRateLimiter.EXPECT().Burst().Return(secondRateLimiterBurst).AnyTimes()

	result := s.rateLimiter.Burst()
	s.Equal(firstRateLimiterBurst, result)
}
