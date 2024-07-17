package tests

import (
	"context"
	"github.com/alserov/restate/estate/internal/service"
	"github.com/alserov/restate/estate/internal/service/models"
	"github.com/alserov/restate/estate/internal/tests/mocks"
	"github.com/alserov/restate/estate/internal/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(serviceSuite))
}

type serviceSuite struct {
	suite.Suite

	ctrl *gomock.Controller

	ctx context.Context

	repo *mocks.MockRepository
}

func (s *serviceSuite) SetupTest() {
	// mocks
	s.ctrl = gomock.NewController(s.T())

	s.repo = mocks.NewMockRepository(s.ctrl)

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
	param := models.GetEstateListParameters{}

	// default
	expect := []models.EstateMainInfo{
		{
			ID:        "1",
			Title:     "title",
			Country:   "uk",
			City:      "london",
			Price:     1000,
			MainImage: "image",
		},
	}

	s.repo.EXPECT().
		GetEstateList(gomock.Any(), gomock.Eq(param)).
		Times(1).
		Return(expect, nil)

	list, err := service.NewService(s.repo).GetEstateList(s.ctx, param)
	s.Require().NoError(err)
	s.Require().Equal(expect, list)

	// no rows error
	s.repo.EXPECT().
		GetEstateList(gomock.Any(), gomock.Eq(param)).
		Times(1).
		Return(nil, utils.NewError("not found", utils.NotFound))

	list, err = service.NewService(s.repo).GetEstateList(s.ctx, param)
	s.Require().Error(err)
	s.Require().Nil(list)
}

func (s *serviceSuite) CreateEstate() {
	estate := models.Estate{
		ID:          "1",
		Title:       "title",
		Description: "desc",
		Price:       100,
		Country:     "uk",
		City:        "london",
		Street:      "baker's street",
		Images:      []string{"img"},
		MainImage:   "img",
		Square:      100,
		Floor:       5,
	}

	// default
	s.repo.EXPECT().
		CreateEstate(gomock.Any(), gomock.AssignableToTypeOf(estate)).
		Times(1).
		Return(nil)

	err := service.NewService(s.repo).CreateEstate(s.ctx, estate)
	s.Require().NoError(err)

	// insert error
	s.repo.EXPECT().
		CreateEstate(gomock.Any(), gomock.AssignableToTypeOf(estate)).
		Times(1).
		Return(utils.NewError("failed to insert", utils.Internal))

	err = service.NewService(s.repo).CreateEstate(s.ctx, estate)
	s.Require().Error(err)
}
