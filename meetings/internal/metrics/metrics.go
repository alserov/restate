package metrics

import (
	"context"
	"time"
)

type Metrics interface {
	TimePerRequest(ctx context.Context, duration time.Duration, handlerName string)
}
