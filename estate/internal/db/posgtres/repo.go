package posgtres

import (
	"context"
	"github.com/alserov/restate/estate/internal/db"
	"github.com/alserov/restate/estate/internal/service/models"
	"github.com/jackc/pgx/v5"
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

func (r repo) GetEstateList(ctx context.Context, param models.GetEstateListParameters) ([]models.EstateMainInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (r repo) GetEstateInfo(ctx context.Context, estateID string) (models.Estate, error) {
	//TODO implement me
	panic("implement me")
}

func (r repo) CreateEstate(ctx context.Context, estate models.Estate) error {
	//TODO implement me
	panic("implement me")
}

func (r repo) DeleteEstate(ctx context.Context, estateID string) error {
	//TODO implement me
	panic("implement me")
}
