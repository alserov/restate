package service

import (
	"context"
	"github.com/alserov/restate/estate/internal/service/models"
)

type Service interface {
	GetEstateList(ctx context.Context, param models.GetEstateListParameters) ([]models.EstateMainInfo, error)
	GetEstateInfo(ctx context.Context, estateID string) (models.Estate, error)

	CreateEstate(ctx context.Context, estate models.Estate) error
	DeleteEstate(ctx context.Context, estateID string) error
}

var _ Service = &service{}

func NewService() *service {
	return &service{}
}

type service struct {
}

func (s service) GetEstateInfo(ctx context.Context, estateID string) (models.Estate, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) CreateEstate(ctx context.Context, estate models.Estate) error {
	//TODO implement me
	panic("implement me")
}

func (s service) DeleteEstate(ctx context.Context, estateID string) error {
	//TODO implement me
	panic("implement me")
}

func (s service) GetEstateList(ctx context.Context, param models.GetEstateListParameters) ([]models.EstateMainInfo, error) {
	//TODO implement me
	panic("implement me")
}
