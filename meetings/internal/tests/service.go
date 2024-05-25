package tests

import (
	"context"
	"github.com/alserov/restate/meetings/internal/service"
	"github.com/alserov/restate/meetings/internal/service/models"
	"github.com/alserov/restate/meetings/internal/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"time"
)

type ServiceTestSuite struct {
	suite.Suite

	ctrl *gomock.Controller
}

func (s *ServiceTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
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
