package log

import (
	"go.uber.org/zap"
)

var _ Logger = &zapLogger{}

func NewZap(env string) *zapLogger {
	l, err := zap.NewProduction()
	if err != nil {
		panic("failed to init zapLogger: " + err.Error())
	}

	return &zapLogger{l.Sugar()}
}

type zapLogger struct {
	*zap.SugaredLogger
}

func (l zapLogger) Trace(key, msg string) {
	l.Infow(msg, "key", key)
}

func (l zapLogger) Debug(s string, data *Data) {
	if data == nil {
		l.Debugw(s)
	} else {
		l.Debugw(s, data.key, data.val)
	}
}

func (l zapLogger) Info(s string, data *Data) {
	if data == nil {
		l.Infow(s)
	} else {
		l.Infow(s, data.key, data.val)
	}
}

func (l zapLogger) Warn(s string, data *Data) {
	if data == nil {
		l.Warnw(s)
	} else {
		l.Warnw(s, data.key, data.val)
	}
}

func (l zapLogger) Error(s string, data *Data) {
	if data == nil {
		l.Errorw(s)
	} else {
		l.Errorw(s, data.key, data.val)
	}
}
