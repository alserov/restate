package db

import (
	"context"
	"github.com/alserov/restate/estate/internal/service/models"
)

type Repository interface {
	GetEstateList(ctx context.Context, param models.GetEstateListParameters) ([]models.EstateMainInfo, error)
	GetEstateInfo(ctx context.Context, estateID string) (models.Estate, error)

	CreateEstate(ctx context.Context, estate models.Estate) error
	DeleteEstate(ctx context.Context, estateID string) error
}
