package redis

import (
	"context"
	"fmt"
	"github.com/alserov/restate/estate/internal/cache"
	rd "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"testing"
)

func TestRedisSuite(t *testing.T) {
	suite.Run(t, new(redisSuite))
}

type redisSuite struct {
	suite.Suite

	repo      cache.Cache
	container testcontainers.Container
	conn      *rd.Client
}

func (r *redisSuite) SetupTest() {
	container := r.newRedisInstance()
	port, err := container.MappedPort(context.Background(), "6379")
	r.Require().NoError(err)

	cl := MustConnect(fmt.Sprintf("127.0.0.1:%v", port.Port()))
	r.Require().NoError(cl.Ping(context.Background()).Err())

	r.repo = NewCache(cl)
}

func (r *redisSuite) TeardownTest() {
	r.Require().NoError(r.container.Terminate(context.Background()))
	r.Require().NoError(r.conn.Close())
}

func (r *redisSuite) TestDefault() {
	expected := struct {
		Str  string
		Intg int
	}{
		Str:  "a",
		Intg: 10,
	}

	r.repo.Set(context.Background(), "key", expected)

	var actual struct {
		Str  string
		Intg int
	}
	r.repo.Get(context.Background(), "key", &actual)

	r.Require().Equal(expected, actual)
}

func (r *redisSuite) newRedisInstance() testcontainers.Container {
	container, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "redis",
			ExposedPorts: []string{"6379/tcp"},
		},
		Started: true,
	})
	require.NoError(r.T(), err)

	return container
}
