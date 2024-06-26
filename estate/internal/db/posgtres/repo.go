package posgtres

import (
	"context"
	"github.com/alserov/restate/estate/internal/db"
	"github.com/alserov/restate/estate/internal/service/models"
	"github.com/alserov/restate/estate/internal/utils"
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

func (r *repo) GetEstateList(ctx context.Context, param models.GetEstateListParameters) ([]models.EstateMainInfo, error) {
	q := `SELECT * FROM estate WHERE (
    		price >= $1 OR $1 = 0 AND
    		price <= $2 OR $2 = 0 AND
    		square >= $3 OR $3 = 0 AND
    		country = $4 OR $4 = '' AND
    		city = $5 OR $5 = '' AND
    		floor = $6 OR $6 = 0
		) LIMIT $7 OFFSET $8`

	rows, err := r.Query(ctx, q,
		param.MinPrice,
		param.MaxPrice,
		param.Square,
		param.Country,
		param.City,
		param.Floor,
		param.Limit,
		param.Offset)
	if err != nil {
		return nil, utils.NewError(err.Error(), utils.Internal)
	}

	var infos []models.EstateMainInfo
	for rows.Next() {
		var info models.EstateMainInfo
		if err = rows.Scan(&info); err != nil {
			return nil, utils.NewError(err.Error(), utils.Internal)
		}

		infos = append(infos, info)
	}

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "passed GetEstateList repo layer")

	return infos, nil
}

func (r *repo) GetEstateInfo(ctx context.Context, estateID string) (models.Estate, error) {
	q := `SELECT * FROM estate WHERE id = $1`

	var estate models.Estate
	if err := r.QueryRow(ctx, q, estateID).Scan(&estate); err != nil {
		return models.Estate{}, utils.NewError(err.Error(), utils.Internal)
	}

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "passed GetEstateInfo repo layer")

	return estate, nil
}

func (r *repo) CreateEstate(ctx context.Context, estate models.Estate) error {
	q := `INSERT INTO estate (
        id,
		title,
		description,
		price,
		country,
		city,
		street,
		images,
		mainImage,
		square,
		floor            
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := r.Exec(ctx, q,
		estate.ID,
		estate.Title,
		estate.Description,
		estate.Price,
		estate.Country,
		estate.City,
		estate.Street,
		estate.MainImage,
		estate.Square,
		estate.Floor)
	if err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "passed CreateEstate repo layer")

	return nil
}

func (r *repo) DeleteEstate(ctx context.Context, estateID string) error {
	q := `DELETE FROM estate WHERE id = $1`

	_, err := r.Exec(ctx, q, estateID)
	if err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "passed DeleteEstate repo layer")

	return nil
}
