package wrappers

import (
	"context"
	"github.com/alserov/restate/meetings/internal/log"
)

func WithLogger(ctx context.Context, args ...any) context.Context {
	l, ok := args[0].(log.Logger)
	if !ok {
		panic("invalid argument")
	}

	ctx = context.WithValue(ctx, ContextLogger, l)
	return context.WithValue(ctx, ContextLogger, l)
}

func ExtractLogger(ctx context.Context) log.Logger {
	l, ok := ctx.Value(ContextLogger).(log.Logger)
	if !ok {
		panic("can not get logger from context")
	}

	return l
}
