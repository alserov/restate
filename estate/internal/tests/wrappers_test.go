package tests

import (
	"context"
	"github.com/alserov/restate/estate/internal/log"
	"github.com/alserov/restate/estate/internal/wrappers"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestWrapperSuite(t *testing.T) {
	suite.Run(t, new(WrappersSuite))
}

type WrappersSuite struct {
	suite.Suite

	ctx context.Context
}

func (wp *WrappersSuite) SetupTest() {
	wp.ctx = context.Background()
}

func (wp *WrappersSuite) TestIdempotencyKey() {
	wp.ctx = wrappers.WithIdempotencyKey(wp.ctx)

	key := wrappers.ExtractIdempotencyKey(wp.ctx)

	wp.Require().NotEqual("", key)
}

func (wp *WrappersSuite) TestLogger() {
	l := log.NewLogger(log.EnvLocal, log.KindZap)

	wp.ctx = wrappers.WithLogger(wp.ctx, l)

	extractedL := wrappers.ExtractLogger(wp.ctx)

	wp.Require().Equal(l, extractedL)
}
