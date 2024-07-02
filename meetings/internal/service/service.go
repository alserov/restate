package service

import (
	"context"
	"fmt"
	"github.com/alserov/restate/meetings/internal/db"
	"github.com/alserov/restate/meetings/internal/service/models"
	"github.com/alserov/restate/meetings/internal/utils"
	"github.com/google/uuid"
	"time"
)

type Service interface {
	ArrangeMeeting(ctx context.Context, m models.Meeting) error
	CancelMeeting(ctx context.Context, parameter models.CancelMeetingParameter) error
	GetAvailableTimeForMeeting(ctx context.Context, estateID string) ([]time.Time, error)

	GetMeetingsByEstateID(ctx context.Context, estateID string) ([]models.Meeting, error)
	GetMeetingsByPhoneNumber(ctx context.Context, phoneNumber string) ([]models.Meeting, error)
}

var _ Service = &service{}

func NewService(repo db.Repository) *service {
	return &service{repo: repo}
}

type service struct {
	repo db.Repository
}

func (s *service) GetMeetingsByEstateID(ctx context.Context, estateID string) ([]models.Meeting, error) {
	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "passed GetMeetingsByEstateID service layer")

	mtngs, err := s.repo.GetMeetingsByEstateID(ctx, estateID)
	if err != nil {
		return nil, fmt.Errorf("failed to get meeting by estate id: %w", err)
	}

	return mtngs, nil
}

func (s *service) GetMeetingsByPhoneNumber(ctx context.Context, phoneNumber string) ([]models.Meeting, error) {
	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "passed GetMeetingsByPhoneNumber service layer")

	mtngs, err := s.repo.GetMeetingsByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get meeting by phone number: %w", err)
	}

	return mtngs, nil
}

func (s *service) ArrangeMeeting(ctx context.Context, m models.Meeting) error {
	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "passed ArrangeMeeting service layer")

	m.ID = uuid.NewString()

	err := s.repo.ArrangeMeeting(ctx, m)
	if err != nil {
		return fmt.Errorf("failed to arrange meeting: %w", err)
	}

	return nil
}

func (s *service) CancelMeeting(ctx context.Context, parameter models.CancelMeetingParameter) error {
	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "passed CancelMeeting service layer")

	err := s.repo.CancelMeeting(ctx, parameter)
	if err != nil {
		return fmt.Errorf("failed to cancel meeting: %w", err)
	}

	return nil
}

func (s *service) GetAvailableTimeForMeeting(ctx context.Context, estateID string) ([]time.Time, error) {
	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "passed GetAvailableTimeForMeeting service layer")

	tStamps, err := s.repo.GetMeetingTimestamps(ctx, estateID)
	if err != nil {
		return nil, fmt.Errorf("failed to get available time for meeting: %w", err)
	}

	availableTStamps := selectAvailableTStampsForMeeting(tStamps)

	return availableTStamps, nil
}

func selectAvailableTStampsForMeeting(tStamps []time.Time) []time.Time {
	var availableTStamps []time.Time

	// filling timestamps before first
	t := time.Now()

	if len(tStamps) > 0 {
		t = tStamps[0].Add(-90 * time.Minute)
	}

	for t.Hour() >= models.MinMeetingTimestamp {
		availableTStamps = append(availableTStamps, t)
		t = t.Add(-90 * time.Minute)
	}

	// filling timestamps between
	for i := 0; i < len(tStamps)-1; i++ {
		if tStamps[i+1].Sub(tStamps[i]) > time.Minute*90 {
			t = tStamps[i+1].Add(-90 * time.Minute)
			if t.Hour() >= models.MinMeetingTimestamp {
				availableTStamps = append(availableTStamps, t)
			}
		}
	}

	// filling timestamps after last (1 month)
	lastMeetingTStamp := time.Now()

	if len(tStamps) > 0 {
		lastMeetingTStamp = tStamps[len(tStamps)-1]
	}
	for i := 0; i < 30; i++ {
		t = lastMeetingTStamp.Add(90 * time.Minute)

		if t.Hour() >= models.MaxMeetingTimestamp {
			now := time.Now()
			t = time.Date(now.Year(), now.Month(), now.Day()+1, models.MinMeetingTimestamp, 0, 0, 0, time.UTC)
		}

		availableTStamps = append(availableTStamps, t)
		lastMeetingTStamp = t
	}

	return availableTStamps
}
