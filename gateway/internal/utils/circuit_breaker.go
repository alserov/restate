package utils

import (
	"github.com/alserov/restate/gateway/internal/log"
	"sync"
	"sync/atomic"
	"time"
)

func NewBreaker(lim uint32, retry time.Duration, log log.Logger) *Breaker {
	return &Breaker{status: ok, lim: lim, retryPeriod: retry, log: log}
}

const (
	ok uint = iota
	check
	closed
)

type Breaker struct {
	log log.Logger

	status uint

	retryPeriod time.Duration
	lastAttempt time.Time
	lim         uint32
	curr        atomic.Uint32

	mu sync.Mutex
}

// Execute
// argument function should return true if it is an internal error, otherwise false
func (b *Breaker) Execute(fn func() (bool, error)) error {
	go func() {
		for range time.Tick(b.retryPeriod) {
			if b.status == closed {
				b.mu.Lock()
				b.status = check
				b.mu.Unlock()

				b.log.Info("updated breaker status", log.WithData("status", check))
			}
		}
	}()

	switch b.status {
	case closed:
		return NewError("breaker timeout", Internal)
	case check:
		isInternalErr, fnErr := fn()

		if isInternalErr {
			b.mu.Lock()
			b.lastAttempt = time.Now()
			b.status = closed
			b.mu.Unlock()

			b.log.Info("updated breaker status", log.WithData("status", closed))

			return fnErr
		} else {
			b.mu.Lock()
			b.status = ok
			b.mu.Unlock()

			b.log.Info("updated breaker status", log.WithData("status", ok))

			return fnErr
		}
	default:
		isInternalErr, fnErr := fn()
		if isInternalErr {
			b.curr.Add(1)
			if b.curr.Load() == b.lim {
				b.mu.Lock()
				b.lastAttempt = time.Now()
				b.status = closed
				b.mu.Unlock()

				b.log.Info("updated breaker status", log.WithData("status", closed))
			}
		}

		return fnErr
	}
}
