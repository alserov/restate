package log

type Logger interface {
	Info(s string, data *Data)
	Warn(s string, data *Data)
	Error(s string, data *Data)
	Debug(s string, data *Data)

	Trace(key string, msg string)
}

type Data struct {
	key string
	val any
}

func WithData(key string, val any) *Data {
	return &Data{key: key, val: val}
}

const (
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
