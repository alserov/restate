package clients

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
	Addr() string
	SetClient(cl any)
	Type() uint
}

type client[T any] struct {
	cl     T
	clType uint

	addr string
}

func (c *client[T]) Type() uint {
	return c.clType
}

func (c *client[T]) Addr() string {
	return c.addr
}

func (c *client[T]) SetClient(cl any) {
	cln, ok := cl.(T)
	if !ok {
		panic("invalid client")
	}

	c.cl = cln
}

const (
	GRPCClient = iota
)

func NewClient[T any](addr string, clientType uint) Client {
	return &client[T]{
		addr:   addr,
		clType: clientType,
	}
}

func Dial(cls ...Client) {
	for _, cl := range cls {
		switch cl.Type() {
		case GRPCClient:
			cln, err := grpc.NewClient(cl.Addr(), grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				panic("failed to dial to client: " + err.Error())
			}

			cl.SetClient(cln)
		default:
			panic("invalid client type")
		}
	}
}
