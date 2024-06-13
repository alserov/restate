package log

import (
	"context"
)

type Logger interface {
	Info(s string, data *Data)
	Warn(s string, data *Data)
	Error(s string, data *Data)
	Debug(s string, data *Data)

	Trace(ctx context.Context, msg string)
}

type Data struct {
	key string
	val any
}

type ContextKey string

func FromCtx(ctx context.Context) Logger {
	l, ok := ctx.Value(ContextLogger).(Logger)
	if !ok {
		panic("can not get logger from context")
	}

	return l
}

func WithData(key string, val any) *Data {
	return &Data{key: key, val: val}
}

const (
	ContextLogger         ContextKey = "log"
	ContextIdempotencyKey ContextKey = "ikey"

	EnvLocal = "local"
	EnvProd  = "prod"

	KindZap = "zap"
)

func NewLogger(env, kind string) Logger {
	var l Logger

	switch kind {
	case KindZap:
		l = NewZap(env)
	default:
		panic("invalid kind: " + kind)
	}

	return l
}
