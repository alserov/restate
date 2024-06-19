package tests

import (
	"context"
	"github.com/alserov/restate/meetings/internal/service"
	"github.com/alserov/restate/meetings/internal/service/models"
	"github.com/alserov/restate/meetings/internal/tests/mocks"
	"github.com/golang/mock/gomock"
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
	repo := mocks.NewMockRepository(s.ctrl)
	id := "id"
	expectedMtngs := []models.Meeting{
		{
			ID:           "id",
			Timestamp:    time.Now(),
			EstateID:     id,
			VisitorPhone: "12345",
		},
	}

	repo.EXPECT().
		GetMeetingsByEstateID(gomock.Any(), gomock.Eq(id)).
		Return(expectedMtngs, nil).
		Times(1)

	srvc := service.NewService(repo)

	mtngs, err := srvc.GetMeetingsByEstateID(context.Background(), id)
	s.Require().Nil(err)
	s.Require().Equal(expectedMtngs, mtngs)
}

func (s *ServiceTestSuite) TestGetMeetingsByPhoneNumber() {
	repo := mocks.NewMockRepository(s.ctrl)
	phoneNumber := "1"
	expectedMtngs := []models.Meeting{
		{
			ID:           "id",
			Timestamp:    time.Now(),
			EstateID:     "id",
			VisitorPhone: phoneNumber,
		},
	}

	repo.EXPECT().
		GetMeetingsByPhoneNumber(gomock.Any(), gomock.Eq(phoneNumber)).
		Return(expectedMtngs, nil).
		Times(1)

	srvc := service.NewService(repo)

	mtngs, err := srvc.GetMeetingsByPhoneNumber(context.Background(), phoneNumber)
	s.Require().Nil(err)
	s.Require().Equal(expectedMtngs, mtngs)
}

func (s *ServiceTestSuite) TestArrangeMeeting() {
	repo := mocks.NewMockRepository(s.ctrl)
	mtng := models.Meeting{
		ID:           "id",
		Timestamp:    time.Now(),
		EstateID:     "id",
		VisitorPhone: "1",
	}

	repo.EXPECT().
		ArrangeMeeting(gomock.Any(), gomock.Eq(phoneNumber)).
		Return(expectedMtngs, nil).
		Times(1)

	srvc := service.NewService(repo)

	mtngs, err := srvc.GetMeetingsByPhoneNumber(context.Background(), phoneNumber)
	s.Require().Nil(err)
	s.Require().Equal(expectedMtngs, mtngs)
}
