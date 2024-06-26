package posgtres

import (
	"context"
	"fmt"
	"github.com/alserov/restate/meetings/internal/db"
	"github.com/alserov/restate/meetings/internal/service/models"
	"github.com/alserov/restate/meetings/internal/utils"
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

func (r *repo) GetMeetingsByEstateID(ctx context.Context, estateID string) ([]models.Meeting, error) {
	l := utils.ExtractLogger(ctx)
	key := utils.ExtractIdempotencyKey(ctx)

	q := `SELECT * FROM meetings WHERE timestamp > $1 AND estate_id = $2`

	rows, err := r.Query(ctx, q, time.Now(), time.Now(), estateID)
	if err != nil {
		return nil, fmt.Errorf("failed to select: %w", err)
	}

	var meetings []models.Meeting
	for rows.Next() {
		var meeting models.Meeting
		if err = rows.Scan(&meeting); err != nil {
			return nil, fmt.Errorf("failed to scan: %w", err)
		}

		meetings = append(meetings, meeting)
	}

	l.Trace(key, "passed GetMeetingsByEstateID repo layer")

	return meetings, nil
}

func (r *repo) GetMeetingsByPhoneNumber(ctx context.Context, phoneNumber string) ([]models.Meeting, error) {
	l := utils.ExtractLogger(ctx)
	key := utils.ExtractIdempotencyKey(ctx)

	q := `SELECT * FROM meetings WHERE timestamp > $1 AND visitor_phone = $2`

	rows, err := r.Query(ctx, q, time.Now(), time.Now(), phoneNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to select: %w", err)
	}

	var meetings []models.Meeting
	for rows.Next() {
		var meeting models.Meeting
		if err = rows.Scan(&meeting); err != nil {
			return nil, fmt.Errorf("failed to scan: %w", err)
		}

		meetings = append(meetings, meeting)
	}

	l.Trace(key, "passed GetMeetingsByPhoneNumber repo layer")

	return meetings, nil
}

func (r *repo) ArrangeMeeting(ctx context.Context, m models.Meeting) error {
	l := utils.ExtractLogger(ctx)
	key := utils.ExtractIdempotencyKey(ctx)

	q := `INSERT INTO meetings (id,timestamp,estate_id,visitor_phone) VALUES ($1,$2,$3,$4)`

	_, err := r.Exec(ctx, q, m.ID, m.Timestamp, m.EstateID, m.VisitorPhone)
	if err != nil {
		return fmt.Errorf("failed to insert: %w", err)
	}

	l.Trace(key, "passed ArrangeMeeting repo layer")

	return nil
}

func (r *repo) CancelMeeting(ctx context.Context, parameter models.CancelMeetingParameter) error {
	l := utils.ExtractLogger(ctx)
	key := utils.ExtractIdempotencyKey(ctx)

	q := `DELETE FROM meetings WHERE id = $1 AND visitor_phone = $2`

	_, err := r.Exec(ctx, q, parameter.ID, parameter.VisitorPhone)
	if err != nil {
		return fmt.Errorf("failed to delete: %w", err)
	}

	l.Trace(key, "passed CancelMeeting repo layer")

	return nil
}

func (r *repo) GetMeetingTimestamps(ctx context.Context, estateID string) ([]time.Time, error) {
	l := utils.ExtractLogger(ctx)
	key := utils.ExtractIdempotencyKey(ctx)

	q := `SELECT timestamp FROM meetings WHERE timestamp > $1 AND estate_id = $2 ORDER BY timestamp`

	rows, err := r.Query(ctx, q, time.Now(), estateID)
	if err != nil {
		return nil, fmt.Errorf("failed to select: %w", err)
	}

	var tStamps []time.Time
	for rows.Next() {
		var tStamp time.Time
		if err = rows.Scan(&tStamp); err != nil {
			return nil, fmt.Errorf("failed to scan: %w", err)
		}

		tStamps = append(tStamps, tStamp)
	}

	l.Trace(key, "passed GetMeetingTimestamps repo layer")

	return tStamps, nil
}
