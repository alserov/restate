package service

import (
	"context"
	"fmt"
	"github.com/alserov/restate/estate/internal/db"
	"github.com/alserov/restate/estate/internal/middleware/grpc/wrappers"
	"github.com/alserov/restate/estate/internal/service/models"
	"github.com/google/uuid"
)

type Service interface {
	GetEstateList(ctx context.Context, param models.GetEstateListParameters) ([]models.EstateMainInfo, error)
	GetEstateInfo(ctx context.Context, estateID string) (models.Estate, error)

	CreateEstate(ctx context.Context, estate models.Estate) error
	DeleteEstate(ctx context.Context, estateID string) error
}

var _ Service = &service{}

func NewService(repo db.Repository) *service {
	return &service{repo: repo}
}

type service struct {
	repo db.Repository
}

func (s *service) GetEstateInfo(ctx context.Context, estateID string) (models.Estate, error) {
	wrappers.ExtractLogger(ctx).Trace(wrappers.ExtractIdempotencyKey(ctx), "passed GetEstateInfo service layer")

	estate, err := s.repo.GetEstateInfo(ctx, estateID)
	if err != nil {
		return models.Estate{}, fmt.Errorf("failed to get estate info: %w", err)
	}

	return estate, nil
}

func (s *service) CreateEstate(ctx context.Context, estate models.Estate) error {
	estate.ID = uuid.NewString()

	wrappers.ExtractLogger(ctx).Trace(wrappers.ExtractIdempotencyKey(ctx), "passed CreateEstate service layer")

	err := s.repo.CreateEstate(ctx, estate)
	if err != nil {
		return fmt.Errorf("failed to insert estate: %w", err)
	}

	return nil
}

func (s *service) DeleteEstate(ctx context.Context, estateID string) error {
	wrappers.ExtractLogger(ctx).Trace(wrappers.ExtractIdempotencyKey(ctx), "passed DeleteEstate service layer")

	err := s.repo.DeleteEstate(ctx, estateID)
	if err != nil {
		return fmt.Errorf("failed to delete estate: %w", err)
	}

	return nil
}

func (s *service) GetEstateList(ctx context.Context, param models.GetEstateListParameters) ([]models.EstateMainInfo, error) {
	wrappers.ExtractLogger(ctx).Trace(wrappers.ExtractIdempotencyKey(ctx), "passed GetEstateList service layer")

	list, err := s.repo.GetEstateList(ctx, param)
	if err != nil {
		return nil, fmt.Errorf("failed to get estate list: %w", err)
	}

	return list, nil
}
