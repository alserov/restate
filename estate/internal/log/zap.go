package log

import "go.uber.org/zap"

func NewZap(env string) *logger {
	l, err := zap.NewProduction()
	if err != nil {
		panic("failed to init logger: " + err.Error())
	}

	return &logger{l.Sugar()}
}

type logger struct {
	*zap.SugaredLogger
}

func (l logger) Debug(s string, data *Data) {
	if data == nil {
		l.Debugw(s)
	} else {
		l.Debugw(s, data.key, data.val)
	}
}

func (l logger) Info(s string, data *Data) {
	if data == nil {
		l.Infow(s)
	} else {
		l.Infow(s, data.key, data.val)
	}
}

func (l logger) Warn(s string, data *Data) {
	if data == nil {
		l.Warnw(s)
	} else {
		l.Warnw(s, data.key, data.val)
	}
}

func (l logger) Error(s string, data *Data) {
	if data == nil {
		l.Errorw(s)
	} else {
		l.Errorw(s, data.key, data.val)
	}
}
