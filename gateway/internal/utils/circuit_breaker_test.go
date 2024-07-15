package utils

import (
	"fmt"
	"github.com/alserov/restate/gateway/internal/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

func TestCircuitBreaker(t *testing.T) {
	suite.Run(t, new(circuitBreakerSuite))
}

type circuitBreakerSuite struct {
	suite.Suite
}

const (
	isInternalErr    = 0
	isNotInternalErr = 1
)

func (cb *circuitBreakerSuite) TestDefault() {
	lg := mocks.NewMockLogger(gomock.NewController(cb.T()))
	lg.EXPECT().
		Info(gomock.Any(), gomock.Any()).
		AnyTimes()

	n := 3
	t := time.Millisecond * 500
	b := NewBreaker(uint32(n), t, lg)

	// check if 'ok' status will work
	err := b.Execute(func() (bool, error) {
		return isNotInternalErr == Internal, nil
	})
	cb.Require().Nil(err)

	// check if change status to 'closed' on n failures
	for i := 0; i < n; i++ {
		err = b.Execute(func() (bool, error) {
			return isInternalErr == Internal, fmt.Errorf("some internal error")
		})
		cb.Require().NotNil(err)
	}

	// check if breaker will stop incoming request
	err = b.Execute(func() (bool, error) {
		return isNotInternalErr == Internal, nil
	})
	cb.Require().NotNil(err)

	// wait for breaker switch its status to 'check', then request will fail again and breaker will set status to 'closed'
	time.Sleep(t)
	err = b.Execute(func() (bool, error) {
		return isInternalErr == Internal, fmt.Errorf("some internal error")
	})
	cb.Require().NotNil(err)

	// wait for breaker to switch its status to 'check', then incoming request will be successful
	time.Sleep(t)
	err = b.Execute(func() (bool, error) {
		return isNotInternalErr == Internal, nil
	})
	cb.Require().Nil(err)
}
