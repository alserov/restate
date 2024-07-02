package services

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

const (
	GRPCClient = iota
	HTTPClient
)

type GRPCClientInitFn func(cc *grpc.ClientConn) any

func Dial[T any](addr string, clientType uint, fn ...GRPCClientInitFn) T {
	return dial[T](addr, clientType, fn...)
}

func dial[T any](addr string, clType uint, fn ...GRPCClientInitFn) T {
	switch clType {
	case GRPCClient:
		cc, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic("failed to services to client: " + err.Error())
		}

		cl, ok := fn[0](cc).(T)
		if !ok {
			panic("invalid dial func")
		}

		return cl
	case HTTPClient:
		c := func() any {
			return http.DefaultClient
		}()

		cl, ok := c.(T)
		if !ok {
			panic("failed to init http client")
		}

		return cl
	default:
		panic("invalid client type")
	}
}
