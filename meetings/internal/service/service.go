package service

import (
	"context"
	"fmt"
	"github.com/alserov/restate/meetings/internal/db"
	"github.com/alserov/restate/meetings/internal/service/models"
	"github.com/google/uuid"
	"time"
)

type Service interface {
	ArrangeMeeting(ctx context.Context, m models.Meeting) error
	CancelMeeting(ctx context.Context, parameter models.CancelMeetingParameter) error
	GetAvailableTimeForMeeting(ctx context.Context, estateID string) ([]time.Time, error)
}

var _ Service = &service{}

func NewService(repo db.Repository) *service {
	return &service{repo: repo}
}

type service struct {
	repo db.Repository
}

func (s *service) ArrangeMeeting(ctx context.Context, m models.Meeting) error {
	m.ID = uuid.NewString()

	err := s.repo.ArrangeMeeting(ctx, m)
	if err != nil {
		return fmt.Errorf("failed to arrange meeting: %w", err)
	}

	return nil
}

func (s *service) CancelMeeting(ctx context.Context, parameter models.CancelMeetingParameter) error {
	err := s.repo.CancelMeeting(ctx, parameter)
	if err != nil {
		return fmt.Errorf("failed to cancel meeting: %w", err)
	}

	return nil
}

func (s *service) GetAvailableTimeForMeeting(ctx context.Context, estateID string) ([]time.Time, error) {
	tStamps, err := s.repo.GetAvailableTimeForMeeting(ctx, estateID)
	if err != nil {
		return nil, fmt.Errorf("failed to get available time for meeting: %w", err)
	}

	return tStamps, nil
}
