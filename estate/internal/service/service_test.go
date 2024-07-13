package service

import (
	"context"
	"github.com/alserov/restate/estate/internal/service/models"
	"github.com/alserov/restate/estate/internal/tests/mocks"
	"github.com/alserov/restate/estate/internal/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestService(t *testing.T) {
	suite.Run(t, new(serviceSuite))
}

type serviceSuite struct {
	suite.Suite
	ctrl *gomock.Controller

	ctx context.Context
}

func (s *serviceSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())

	// mocks
	s.ctrl = gomock.NewController(s.T())

	logger := mocks.NewMockLogger(s.ctrl)
	logger.EXPECT().Trace(gomock.Any(), gomock.Any()).AnyTimes()

	// ctx
	s.ctx = context.WithValue(context.Background(), utils.ContextLogger, logger)
	s.ctx = context.WithValue(s.ctx, utils.ContextIdempotencyKey, uuid.NewString())
}

func (s *serviceSuite) TeardownTest() {
	s.ctrl.Finish()
}

func (s *serviceSuite) TestGetEstateList() {
	expectedData := []models.EstateMainInfo{
		{
			ID:        uuid.NewString(),
			Title:     "title",
			Country:   "uk",
			City:      "london",
			Price:     100,
			MainImage: "img",
		},
	}

	repo := mocks.NewMockRepository(s.ctrl)
	repo.EXPECT().
		GetEstateList(gomock.Any(), gomock.Eq(models.GetEstateListParameters{
			MinPrice: expectedData[0].Price,
			MaxPrice: expectedData[0].Price,
			Square:   0,
			Country:  expectedData[0].Country,
			City:     expectedData[0].City,
			Floor:    0,
			Limit:    1,
			Offset:   0,
		})).
		Return(expectedData, nil).
		Times(1)

	srvc := NewService(repo)

	estate, err := srvc.GetEstateList(s.ctx, models.GetEstateListParameters{
		MinPrice: expectedData[0].Price,
		MaxPrice: expectedData[0].Price,
		Square:   0,
		Country:  expectedData[0].Country,
		City:     expectedData[0].City,
		Floor:    0,
		Limit:    1,
		Offset:   0,
	})
	s.Require().NoError(err)
	s.Require().Equal(expectedData, estate)
}
