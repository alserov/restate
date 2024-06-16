package service

import (
	"context"
	"fmt"
	"github.com/alserov/restate/estate/internal/db"
	"github.com/alserov/restate/estate/internal/service/models"
	"github.com/alserov/restate/estate/internal/wrappers"
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
	l := wrappers.ExtractLogger(ctx)
	key := wrappers.ExtractIdempotencyKey(ctx)

	l.Trace(key, "passed GetEstateInfo service layer")

	estate, err := s.repo.GetEstateInfo(ctx, estateID)
	if err != nil {
		return models.Estate{}, fmt.Errorf("failed to get estate info: %w", err)
	}

	return estate, nil
}

func (s *service) CreateEstate(ctx context.Context, estate models.Estate) error {
	l := wrappers.ExtractLogger(ctx)
	key := wrappers.ExtractIdempotencyKey(ctx)

	estate.ID = uuid.NewString()

	l.Trace(key, "passed CreateEstate service layer")

	err := s.repo.CreateEstate(ctx, estate)
	if err != nil {
		return fmt.Errorf("failed to insert estate: %w", err)
	}

	return nil
}

func (s *service) DeleteEstate(ctx context.Context, estateID string) error {
	l := wrappers.ExtractLogger(ctx)
	key := wrappers.ExtractIdempotencyKey(ctx)

	l.Trace(key, "passed DeleteEstate service layer")

	err := s.repo.DeleteEstate(ctx, estateID)
	if err != nil {
		return fmt.Errorf("failed to delete estate: %w", err)
	}

	return nil
}

func (s *service) GetEstateList(ctx context.Context, param models.GetEstateListParameters) ([]models.EstateMainInfo, error) {
	l := wrappers.ExtractLogger(ctx)
	key := wrappers.ExtractIdempotencyKey(ctx)

	l.Trace(key, "passed GetEstateLis service layer")

	list, err := s.repo.GetEstateList(ctx, param)
	if err != nil {
		return nil, fmt.Errorf("failed to get estate list: %w", err)
	}

	return list, nil
}
