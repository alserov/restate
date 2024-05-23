package posgtres

import (
	"context"
	"github.com/alserov/restate/meetings/internal/db"
	"github.com/alserov/restate/meetings/internal/service/models"
	"github.com/jackc/pgx/v5"
	"time"
)

var _ db.Repository = &repo{}

func NewRepository(conn *pgx.Conn) *repo {
	return &repo{
		conn,
	}
}

type repo struct {
	*pgx.Conn
}

func (r repo) ArrangeMeeting(ctx context.Context, m models.Meeting) error {
	//TODO implement me
	panic("implement me")
}

func (r repo) CancelMeeting(ctx context.Context, parameter models.CancelMeetingParameter) error {
	//TODO implement me
	panic("implement me")
}

func (r repo) GetAvailableTimeForMeeting(ctx context.Context, estateID string) ([]time.Time, error) {
	//TODO implement me
	panic("implement me")
}
