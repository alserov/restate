package wrappers

type ContextKey string

const (
	ContextIdempotencyKey ContextKey = "key"
	ContextLogger         ContextKey = "log"
)
