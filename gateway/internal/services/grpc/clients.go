package grpc

import (
	estate "github.com/alserov/restate/estate/pkg/grpc"
	meetings "github.com/alserov/restate/meetings/pkg/grpc"
	"google.golang.org/grpc"
)

func NewEstateClient(cc *grpc.ClientConn) any {
	return estate.NewEstateServiceClient(cc)
}

func NewMeetingsClient(cc *grpc.ClientConn) any {
	return meetings.NewMeetingsServiceClient(cc)
}
