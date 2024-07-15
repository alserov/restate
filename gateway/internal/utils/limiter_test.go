package utils

import (
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

func TestLimiterSuite(t *testing.T) {
	suite.Run(t, new(limiterSuite))
}

type limiterSuite struct {
	suite.Suite

	lim Limiter
}

func (l *limiterSuite) SetupTest() {
	l.lim = NewLimiter(context.Background(), 100, time.Millisecond*500)
}

func (l *limiterSuite) TestDefault() {
	for i := 0; i < 101; i++ {
		ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*3)
		if i < 100 {
			l.Require().True(l.lim.Allow(ctx))
		} else {
			l.Require().False(l.lim.Allow(ctx))
		}
	}

	time.Sleep(time.Millisecond * 500)

	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*3)
	l.Require().True(l.lim.Allow(ctx))
}
