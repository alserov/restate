package tests

import (
	"context"
	"fmt"
	"github.com/alserov/restate/estate/internal/db"
	"github.com/alserov/restate/estate/internal/db/posgtres"
	"github.com/alserov/restate/estate/internal/service/models"
	"github.com/alserov/restate/estate/internal/tests/mocks"
	"github.com/alserov/restate/estate/internal/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"
)

func TestPostgresRepo(t *testing.T) {
	suite.Run(t, new(postgresSuite))
}

type postgresSuite struct {
	suite.Suite
	ctrl *gomock.Controller

	repo               db.Repository
	conn               *sqlx.DB
	container          testcontainers.Container
	user, password, db string

	ctx context.Context

	estateModels []models.Estate
}

func (p *postgresSuite) SetupTest() {
	p.user = "postgres"
	p.password = "postgres"
	p.db = "postgres"

	// starting container
	container := p.newPostgresInstance()
	port, err := container.MappedPort(context.Background(), "5432")
	p.Require().NoError(err)

	// connecting to container
	conn, err := sqlx.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.user, p.password, "127.0.0.1", port.Port(), p.db))
	p.Require().NoError(err)
	p.Require().NoError(conn.Ping())

	// migrations
	p.Require().NoError(goose.SetDialect("postgres"))
	p.Require().NoError(goose.Up(conn.DB, "../db/migrations"))

	p.conn = conn
	p.container = container
	p.repo = posgtres.NewRepository(conn)

	// prepare db data
	p.estateModels = []models.Estate{
		{
			ID:          uuid.NewString(),
			Title:       "title",
			Description: "abc",
			Price:       10,
			Country:     "uk",
			City:        "london",
			Street:      "street",
			MainImage:   "img",
			Images:      pq.StringArray{"img1", "img2"},
			Square:      100,
			Floor:       5,
		},
		{
			ID:          uuid.NewString(),
			Title:       "title 1",
			Description: "abc",
			Price:       10,
			Country:     "uk",
			City:        "london",
			Street:      "street",
			MainImage:   "img",
			Images:      pq.StringArray{"img1", "img2"},
			Square:      100,
			Floor:       5,
		},
	}

	// filling db with prepared data
	p.conn.MustExec(
		`INSERT INTO estate (id,title,description,price,country,city,street,images,main_image,square,floor)
					VALUES
    				($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11), 
                    ($12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)`,
		p.estateModels[0].ID,
		p.estateModels[0].Title,
		p.estateModels[0].Description,
		p.estateModels[0].Price,
		p.estateModels[0].Country,
		p.estateModels[0].City,
		p.estateModels[0].Street,
		p.estateModels[0].Images,
		p.estateModels[0].MainImage,
		p.estateModels[0].Square,
		p.estateModels[0].Floor,

		p.estateModels[1].ID,
		p.estateModels[1].Title,
		p.estateModels[1].Description,
		p.estateModels[1].Price,
		p.estateModels[1].Country,
		p.estateModels[1].City,
		p.estateModels[1].Street,
		p.estateModels[1].Images,
		p.estateModels[1].MainImage,
		p.estateModels[1].Square,
		p.estateModels[1].Floor,
	)

	// mocks
	p.ctrl = gomock.NewController(p.T())

	logger := mocks.NewMockLogger(p.ctrl)
	logger.EXPECT().Trace(gomock.Any(), gomock.Any()).AnyTimes()

	// ctx
	p.ctx = context.WithValue(context.Background(), utils.ContextLogger, logger)
	p.ctx = context.WithValue(p.ctx, utils.ContextIdempotencyKey, uuid.NewString())
}

func (p *postgresSuite) TeardownTest() {
	p.Require().NoError(p.container.Terminate(context.Background()))
	p.ctrl.Finish()
}

func (p *postgresSuite) TestGetEstateList() {
	estate, err := p.repo.GetEstateList(p.ctx, models.GetEstateListParameters{
		MinPrice: p.estateModels[0].Price,
		MaxPrice: p.estateModels[0].Price,
		Square:   p.estateModels[0].Square,
		Country:  p.estateModels[0].Country,
		City:     p.estateModels[0].City,
		Floor:    p.estateModels[0].Floor,
		Limit:    1,
		Offset:   0,
	})
	p.Require().NoError(err)
	p.Require().Greater(len(estate), 0)
	p.Require().Equal(p.estateModels[0].ID, estate[0].ID)

	// =========================

	estate, err = p.repo.GetEstateList(p.ctx, models.GetEstateListParameters{
		MinPrice: p.estateModels[1].Price,
		MaxPrice: p.estateModels[1].Price,
		Square:   p.estateModels[1].Square,
		Country:  p.estateModels[1].Country,
		City:     p.estateModels[1].City,
		Floor:    5,
		Limit:    1,
		Offset:   1,
	})

	p.Require().NoError(err)
	p.Require().Greater(len(estate), 0)
	p.Require().Equal(p.estateModels[1].ID, estate[0].ID)

	// =======================

	estate, err = p.repo.GetEstateList(p.ctx, models.GetEstateListParameters{
		MinPrice: 0,
		MaxPrice: 0,
		Square:   0,
		Country:  "",
		City:     "",
		Floor:    0,
		Limit:    uint32(len(p.estateModels)),
		Offset:   0,
	})

	p.Require().NoError(err)
	p.Require().Greater(len(estate), 0)
	for idx, est := range p.estateModels {
		p.Require().Equal(est.ID, estate[idx].ID)
	}
}

func (p *postgresSuite) TestGetEstateInfo() {
	info, err := p.repo.GetEstateInfo(p.ctx, p.estateModels[0].ID)
	p.Require().NoError(err)
	p.Require().Equal(p.estateModels[0], info)

	// ====================================

	info, err = p.repo.GetEstateInfo(p.ctx, "some random id")
	p.Require().Error(err)
	p.Require().Equal(models.Estate{}, info)
}

func (p *postgresSuite) newPostgresInstance() testcontainers.Container {
	container, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres",
			ExposedPorts: []string{"5432/tcp"},
			Env: map[string]string{
				"POSTGRES_USER":     p.user,
				"POSTGRES_PASSWORD": p.password,
				"POSTGRES_DB":       p.db,
			},
			WaitingFor: wait.ForAll(
				wait.ForLog("database system is ready to accept connections"),
				wait.ForListeningPort("5432/tcp"),
			),
		},
		Started: true,
	})
	require.NoError(p.T(), err)

	return container
}
