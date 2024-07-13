package cache

import (
	"context"
	"github.com/goccy/go-json"
	"time"

	rd "github.com/redis/go-redis/v9"
)

const exp = time.Minute * 15

func newRedis(addr string) *redis {
	return &redis{mustConnectRedis(addr)}
}

func mustConnectRedis(addr string) *rd.Client {
	cl := rd.NewClient(&rd.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,
	})
	cl.Ping(context.TODO())

	return cl
}

var _ Cache = &redis{}

type redis struct {
	cl *rd.Client
}

func (r *redis) Set(ctx context.Context, key string, val any) {
	r.cl.Set(ctx, key, val, exp)
}

func (r *redis) Get(ctx context.Context, key string, dst any) bool {
	val, err := r.cl.Get(ctx, key).Result()
	if err != nil {
		return false
	}

	if err = json.Unmarshal([]byte(val), &dst); err != nil {
		return false
	}

	return true
}
