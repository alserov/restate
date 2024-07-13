package redis

import (
	"context"
	rd "github.com/redis/go-redis/v9"
)

func mustConnect(addr string) *rd.Client {
	cl := rd.NewClient(&rd.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,
	})
	cl.Ping(context.TODO())

	return cl
}
