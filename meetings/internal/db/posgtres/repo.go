package posgtres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/alserov/restate/meetings/internal/db"
	"github.com/alserov/restate/meetings/internal/service/models"
	"github.com/alserov/restate/meetings/internal/utils"
	"github.com/jmoiron/sqlx"
	"time"
)

var _ db.Repository = &repo{}

func NewRepository(conn *sqlx.DB) *repo {
	return &repo{
		conn,
	}
}

type repo struct {
	*sqlx.DB
}

func (r *repo) GetMeetingsByEstateID(ctx context.Context, estateID string) ([]models.Meeting, error) {
	q := `SELECT * FROM meetings WHERE timestamp < $1 AND estate_id = $2`

	rows, err := r.Queryx(q, time.Now(), estateID)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, utils.NewError(err.Error(), utils.NotFound)
		}
		return nil, utils.NewError(err.Error(), utils.Internal)
	}

	var meetings []models.Meeting
	for rows.Next() {
		var meeting models.Meeting
		if err = rows.StructScan(&meeting); err != nil {
			return nil, utils.NewError(err.Error(), utils.Internal)
		}

		meetings = append(meetings, meeting)
	}

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "passed GetMeetingsByEstateID repo layer")

	return meetings, nil
}

func (r *repo) GetMeetingsByPhoneNumber(ctx context.Context, phoneNumber string) ([]models.Meeting, error) {
	q := `SELECT * FROM meetings WHERE timestamp < $1 AND visitor_phone = $2`

	rows, err := r.Queryx(q, time.Now(), phoneNumber)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, utils.NewError(err.Error(), utils.NotFound)
		}
		return nil, utils.NewError(err.Error(), utils.Internal)
	}

	var meetings []models.Meeting
	for rows.Next() {
		var meeting models.Meeting
		if err = rows.StructScan(&meeting); err != nil {
			return nil, utils.NewError(err.Error(), utils.Internal)
		}

		meetings = append(meetings, meeting)
	}

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "passed GetMeetingsByPhoneNumber repo layer")

	return meetings, nil
}

func (r *repo) ArrangeMeeting(ctx context.Context, m models.Meeting) error {
	q := `INSERT INTO meetings (id,timestamp,estate_id,visitor_phone) VALUES ($1,$2,$3,$4)`

	_, err := r.Exec(q, m.ID, m.Timestamp, m.EstateID, m.VisitorPhone)
	if err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "passed ArrangeMeeting repo layer")

	return nil
}

func (r *repo) CancelMeeting(ctx context.Context, parameter models.CancelMeetingParameter) error {
	q := `DELETE FROM meetings WHERE id = $1 AND visitor_phone = $2`

	_, err := r.Exec(q, parameter.ID, parameter.VisitorPhone)
	if err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "passed CancelMeeting repo layer")

	return nil
}

func (r *repo) GetMeetingTimestamps(ctx context.Context, estateID string) ([]time.Time, error) {
	q := `SELECT timestamp FROM meetings WHERE timestamp > $1 AND estate_id = $2 ORDER BY timestamp`

	rows, err := r.Queryx(q, time.Now(), estateID)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, utils.NewError(err.Error(), utils.NotFound)
		}
		return nil, utils.NewError(err.Error(), utils.Internal)
	}

	var tStamps []time.Time
	for rows.Next() {
		var tStamp time.Time
		if err = rows.StructScan(&tStamp); err != nil {
			return nil, utils.NewError(err.Error(), utils.Internal)
		}

		tStamps = append(tStamps, tStamp)
	}

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "passed GetMeetingTimestamps repo layer")

	return tStamps, nil
}
