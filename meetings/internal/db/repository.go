package db

import (
	"context"
	"github.com/alserov/restate/meetings/internal/service/models"
	"time"
)

type Repository interface {
	ArrangeMeeting(ctx context.Context, m models.Meeting) error
	CancelMeeting(ctx context.Context, parameter models.CancelMeetingParameter) error
	GetMeetingTimestamps(ctx context.Context, estateID string) ([]time.Time, error)

	GetMeetingsByEstateID(ctx context.Context, estateID string) ([]models.Meeting, error)
	GetMeetingsByPhoneNumber(ctx context.Context, phoneNumber string) ([]models.Meeting, error)
}
