package tests

import (
	"context"
	"fmt"
	"github.com/alserov/restate/estate/internal/db"
	"github.com/alserov/restate/estate/internal/db/posgtres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

	repo db.Repository

	container testcontainers.Container

	user, password, db string
}

func (p *postgresSuite) SetupTest() {
	p.user = "postgres"
	p.password = "postgres"
	p.db = "postgres"

	container := p.newTestDatabase(p.T())
	conn, err := sqlx.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.user, p.password, "127.0.0.1", "5432", p.db))
	p.Require().NoError(err)

	p.container = container
	p.repo = posgtres.NewRepository(conn)
}

func (p *postgresSuite) TeardownTest() {
	p.Require().NoError(p.container.Terminate(context.Background()))
}

func (p *postgresSuite) TestGetEstateList() {

}

func (p *postgresSuite) newTestDatabase(t *testing.T) testcontainers.Container {
	db, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
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
	require.NoError(t, err)

	return db
}
