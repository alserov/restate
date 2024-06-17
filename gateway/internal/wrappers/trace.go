package wrappers

import "context"

func ExtractIdempotencyKey(ctx context.Context) string {
	key, ok := ctx.Value(ContextIdempotencyKey).(string)
	if !ok {
		panic("can not get idempotency key from context")
	}

	return key
}
