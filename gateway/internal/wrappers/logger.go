package wrappers

import (
	"context"
	"github.com/alserov/restate/gateway/internal/log"
)

func ExtractLogger(ctx context.Context) log.Logger {
	l, ok := ctx.Value(ContextLogger).(log.Logger)
	if !ok {
		panic("can not get logger from context")
	}

	return l
}
