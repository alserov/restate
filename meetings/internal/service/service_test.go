package service

import (
	"context"
	"fmt"
	"github.com/alserov/restate/meetings/internal/mocks"
	"github.com/alserov/restate/meetings/internal/service/models"
	"github.com/alserov/restate/meetings/internal/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

func TestService(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

type ServiceTestSuite struct {
	suite.Suite

	ctrl *gomock.Controller
}

func (s *ServiceTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
}

func (s *ServiceTestSuite) TeardownTest() {
	s.ctrl.Finish()
}

func (s *ServiceTestSuite) TestGetMeetingsByEstateID() {
	id := "id"
	expectedMtngs := []models.Meeting{
		{
			ID:           "id",
			Timestamp:    time.Now(),
			EstateID:     id,
			VisitorPhone: "12345",
		},
	}

	repo := mocks.NewMockRepository(s.ctrl)
	repo.EXPECT().
		GetMeetingsByEstateID(gomock.Any(), gomock.Eq(id)).
		Return(expectedMtngs, nil).
		Times(1)

	l := mocks.NewMockLogger(s.ctrl)
	l.EXPECT().
		Trace(gomock.Any(), gomock.Any()).AnyTimes()

	ctx := utils.WithLogger(l)(utils.WithIdempotencyKey(context.Background()))

	mtngs, err := NewService(repo).GetMeetingsByEstateID(ctx, id)
	s.Require().Nil(err)
	s.Require().Equal(expectedMtngs, mtngs)
}

func (s *ServiceTestSuite) TestGetMeetingsByPhoneNumber() {
	phoneNumber := "1"
	expectedMtngs := []models.Meeting{
		{
			ID:           "id",
			Timestamp:    time.Now(),
			EstateID:     "id",
			VisitorPhone: phoneNumber,
		},
	}

	repo := mocks.NewMockRepository(s.ctrl)
	repo.EXPECT().
		GetMeetingsByPhoneNumber(gomock.Any(), gomock.Eq(phoneNumber)).
		Return(expectedMtngs, nil).
		Times(1)

	l := mocks.NewMockLogger(s.ctrl)
	l.EXPECT().
		Trace(gomock.Any(), gomock.Any()).AnyTimes()

	ctx := utils.WithLogger(l)(utils.WithIdempotencyKey(context.Background()))

	mtngs, err := NewService(repo).GetMeetingsByPhoneNumber(ctx, phoneNumber)
	s.Require().Nil(err)
	s.Require().Equal(expectedMtngs, mtngs)
}

func (s *ServiceTestSuite) TestArrangeMeeting() {
	mtng := models.Meeting{
		ID:           "id",
		Timestamp:    time.Now(),
		EstateID:     "id",
		VisitorPhone: "1",
	}

	repo := mocks.NewMockRepository(s.ctrl)
	repo.EXPECT().
		ArrangeMeeting(gomock.Any(), gomock.Any()).
		Return(nil).
		Times(1)

	l := mocks.NewMockLogger(s.ctrl)
	l.EXPECT().
		Trace(gomock.Any(), gomock.Any()).AnyTimes()

	ctx := utils.WithLogger(l)(utils.WithIdempotencyKey(context.Background()))

	err := NewService(repo).ArrangeMeeting(ctx, mtng)
	s.Require().Nil(err)
}

func (s *ServiceTestSuite) TestCancelMeeting() {
	param := models.CancelMeetingParameter{
		ID:           "id",
		VisitorPhone: "1",
	}

	repo := mocks.NewMockRepository(s.ctrl)
	repo.EXPECT().
		CancelMeeting(gomock.Any(), gomock.Eq(param)).
		Return(nil).
		Times(1)

	l := mocks.NewMockLogger(s.ctrl)
	l.EXPECT().
		Trace(gomock.Any(), gomock.Any()).AnyTimes()

	ctx := utils.WithLogger(l)(utils.WithIdempotencyKey(context.Background()))

	err := NewService(repo).CancelMeeting(ctx, param)
	s.Require().Nil(err)
}

func (s *ServiceTestSuite) TestGetAvailableTimeForMeeting() {
	estateID := "id"

	repo := mocks.NewMockRepository(s.ctrl)
	repo.EXPECT().
		GetMeetingTimestamps(gomock.Any(), gomock.Eq(estateID)).
		Return([]time.Time{
			time.Date(2000, 9, 9, 9, 0, 0, 0, time.UTC),
			time.Date(2000, 9, 9, 10, 30, 0, 0, time.UTC),
			time.Date(2000, 9, 9, 12, 0, 0, 0, time.UTC),
			time.Date(2000, 9, 9, 13, 30, 0, 0, time.UTC),
		}, nil).
		Times(1)

	l := mocks.NewMockLogger(s.ctrl)
	l.EXPECT().
		Trace(gomock.Any(), gomock.Any()).AnyTimes()

	ctx := utils.WithLogger(l)(utils.WithIdempotencyKey(context.Background()))

	tStamps, err := NewService(repo).GetAvailableTimeForMeeting(ctx, estateID)
	s.Require().Nil(err)
	s.Require().NotNil(tStamps)
}

func TestSelectAvailableTStampsForMeeting(t *testing.T) {
	tests := []struct {
		Stamps []time.Time

		BeforeStamps []time.Time
		InnerStamps  []time.Time
		OuterStamps  []time.Time
	}{
		{
			Stamps: []time.Time{
				time.Date(2000, 9, 9, 12, 0, 0, 0, time.UTC),
				time.Date(2000, 9, 9, 13, 30, 0, 0, time.UTC)},
			BeforeStamps: []time.Time{
				time.Date(2000, 9, 9, 10, 30, 0, 0, time.UTC),
				time.Date(2000, 9, 9, 9, 0, 0, 0, time.UTC),
			},
			InnerStamps: []time.Time{},
			OuterStamps: []time.Time{
				time.Date(2000, 9, 9, 15, 0, 0, 0, time.UTC),
				time.Date(2000, 9, 9, 16, 30, 0, 0, time.UTC),
			},
		},
		{
			Stamps: []time.Time{
				time.Date(2000, 9, 9, 12, 0, 0, 0, time.UTC),
				time.Date(2000, 9, 9, 15, 00, 0, 0, time.UTC)},
			BeforeStamps: []time.Time{
				time.Date(2000, 9, 9, 10, 30, 0, 0, time.UTC),
				time.Date(2000, 9, 9, 9, 0, 0, 0, time.UTC),
			},
			InnerStamps: []time.Time{time.Date(2000, 9, 9, 13, 30, 0, 0, time.UTC)},
			OuterStamps: []time.Time{time.Date(2000, 9, 9, 16, 30, 0, 0, time.UTC)},
		},
		{
			Stamps: []time.Time{
				time.Date(2000, 9, 9, 12, 0, 0, 0, time.UTC),
				time.Date(2000, 9, 9, 15, 00, 0, 0, time.UTC)},
			BeforeStamps: []time.Time{
				time.Date(2000, 9, 9, 10, 30, 0, 0, time.UTC),
				time.Date(2000, 9, 9, 9, 0, 0, 0, time.UTC),
			},
			InnerStamps: []time.Time{time.Date(2000, 9, 9, 13, 30, 0, 0, time.UTC)},
			OuterStamps: []time.Time{time.Date(2000, 9, 9, 16, 30, 0, 0, time.UTC)},
		},
	}

	for idx, tc := range tests {
		t.Run(fmt.Sprintf("tc: %d", idx), func(t *testing.T) {
			stamps := selectAvailableTStampsForMeeting(tc.Stamps)

			for i, befTStamp := range tc.BeforeStamps {
				require.Equal(t, befTStamp, stamps[i])
			}

			for i, inTStamp := range tc.InnerStamps {
				require.Equal(t, inTStamp, stamps[i+len(tc.BeforeStamps)])
			}

			for i, outTStamp := range tc.OuterStamps {
				require.Equal(t, outTStamp, stamps[i+len(tc.InnerStamps)+len(tc.BeforeStamps)])
			}
		})
	}
}
