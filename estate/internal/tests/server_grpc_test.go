package tests

import (
	"context"
	"fmt"
	"github.com/alserov/restate/estate/internal/cache/redis"
	"github.com/alserov/restate/estate/internal/db"
	"github.com/alserov/restate/estate/internal/db/posgtres"
	"github.com/alserov/restate/estate/internal/server/grpc"
	"github.com/alserov/restate/estate/internal/service"
	"github.com/alserov/restate/estate/internal/tests/mocks"
	estate "github.com/alserov/restate/estate/pkg/grpc"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	rd "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	gRPC "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"testing"
)

func TestServerGRPCSuite(t *testing.T) {
	suite.Run(t, new(serverGRPCSuite))
}

type serverGRPCSuite struct {
	suite.Suite

	srvr *gRPC.Server
	srvc service.Service
	repo db.Repository

	cl     estate.EstateServiceClient
	clConn *gRPC.ClientConn

	postgresConn       *sqlx.DB
	postgresContainer  testcontainers.Container
	user, password, db string

	redisConn      *rd.Client
	redisContainer testcontainers.Container

	ctrl *gomock.Controller
}

func (s *serverGRPCSuite) SetupTest() {
	// mocks
	ctrl := gomock.NewController(s.T())

	metr := mocks.NewMockMetrics(ctrl)
	metr.EXPECT().
		ObserveRequest(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		AnyTimes()

	logger := mocks.NewMockLogger(ctrl)
	logger.EXPECT().
		Trace(gomock.Any(), gomock.Any()).
		AnyTimes()

	// redis
	redisContainer := s.newRedisInstance()
	port, err := redisContainer.MappedPort(context.Background(), "6379")
	s.Require().NoError(err)

	redisConn := rd.NewClient(&rd.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,
	})

	cache := redis.NewCache(redisConn)

	// ==================

	// postgres
	s.password = "postgres"
	s.user = "postgres"
	s.db = "estate"

	postgresContainer := s.newPostgresInstance()
	port, err = postgresContainer.MappedPort(context.Background(), "5432")
	s.Require().NoError(err)

	postgresConn, err := sqlx.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		s.user, s.password, "127.0.0.1", port.Port(), s.db))
	s.Require().NoError(err)
	s.Require().NoError(postgresConn.Ping())

	s.Require().NoError(goose.SetDialect("postgres"))
	s.Require().NoError(goose.Up(postgresConn.DB, "../db/migrations"))

	// ==================

	repo := posgtres.NewRepository(postgresConn)
	srvc := service.NewService(repo)
	srvr := grpc.RegisterHandler(srvc, cache, metr, logger)

	s.ctrl = ctrl

	s.postgresContainer = postgresContainer
	s.redisContainer = redisContainer

	s.postgresConn = postgresConn
	s.redisConn = redisConn

	s.repo = repo
	s.srvc = srvc
	s.srvr = srvr

	// ====================

	go func() {
		l, err := net.Listen("tcp", ":5000")
		s.Require().NoError(err)
		s.Require().NoError(srvr.Serve(l))
	}()

	// ====================

	cc, err := gRPC.NewClient("localhost:5000", gRPC.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)

	s.clConn = cc
	s.cl = estate.NewEstateServiceClient(cc)
}

func (s *serverGRPCSuite) TeardownTest() {
	s.ctrl.Finish()
	s.Require().NoError(s.clConn.Close())
	s.Require().NoError(s.redisConn.Close())
	s.Require().NoError(s.postgresConn.Close())
	s.Require().NoError(s.postgresContainer.Terminate(context.Background()))
	s.Require().NoError(s.redisContainer.Terminate(context.Background()))
}

func (s *serverGRPCSuite) TestCreateEstate() {
	_, err := s.cl.CreateEstate(context.Background(), &estate.Estate{
		Title:       "title",
		Description: "desc",
		Price:       100,
		Country:     "uk",
		City:        "london",
		Street:      "baker",
		Images:      []string{"img"},
		MainImage:   "img",
		Square:      100,
		Floor:       5,
	})
	s.Require().NoError(err)
}

func (s *serverGRPCSuite) newRedisInstance() testcontainers.Container {
	container, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "redis",
			ExposedPorts: []string{"6379/tcp"},
			WaitingFor: wait.ForAll(
				wait.ForListeningPort("6379/tcp"),
			),
		},
		Started: true,
	})
	s.Require().NoError(err)

	fmt.Println()

	return container
}

func (s *serverGRPCSuite) newPostgresInstance() testcontainers.Container {
	container, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres",
			ExposedPorts: []string{"5432/tcp"},
			Env: map[string]string{
				"POSTGRES_USER":     s.user,
				"POSTGRES_PASSWORD": s.password,
				"POSTGRES_DB":       s.db,
			},
			WaitingFor: wait.ForAll(
				wait.ForListeningPort("5432/tcp"),
			),
		},
		Started: true,
	})
	s.Require().NoError(err)

	return container
}
