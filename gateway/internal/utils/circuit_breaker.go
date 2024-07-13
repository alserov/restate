package utils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
	"sync/atomic"
	"time"
)

type Breaker interface {
	Execute(fn func() (any, error), resp *any, err *error)
}

func NewBreaker(lim uint, clear time.Duration) Breaker {
	return &breaker{}
}

const (
	ok uint = iota
	check
	closed
)

type breaker struct {
	status uint

	lastAttempt time.Time
	lim         uint32
	curr        atomic.Uint32

	mu sync.Mutex
}

func (b *breaker) Execute(fn func() (any, error), resp *any, err *error) {
	switch b.status {
	case closed:
		*resp = nil
		*err = NewError("breaker timeout", Internal)
	case check:
		res, respError := fn()
		st, _ := status.FromError(respError)

		if st.Code() == codes.Internal {
			b.mu.Lock()
			b.lastAttempt = time.Now()
			b.status = closed
			b.mu.Unlock()

			*resp = nil
			*err = NewError(respError.Error(), Internal)
		} else {
			b.mu.Lock()
			b.status = ok
			b.mu.Unlock()

			*resp = res
			*err = respError
		}
	default:
		res, respError := fn()
		st, _ := status.FromError(respError)

		if st.Code() == codes.Internal {
			b.curr.Add(1)
			if b.curr.Load() == b.lim {
				b.mu.Lock()
				b.lastAttempt = time.Now()
				b.status = closed
				b.mu.Unlock()
			}
		}

		*resp = res
		*err = respError
	}
}
