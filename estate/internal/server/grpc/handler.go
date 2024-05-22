package grpc

import (
	"context"
	"github.com/alserov/restate/estate/internal/service"
	estate "github.com/alserov/restate/estate/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func RegisterHandler(srvc service.Service) *grpc.Server {
	srvr := grpc.NewServer()
	estate.RegisterEstateServiceServer(srvr, &handler{srvc: srvc})
	return srvr
}

type handler struct {
	estate.UnimplementedEstateServiceServer

	srvc service.Service
}

func (h handler) GetEstateList(ctx context.Context, parameters *estate.GetListParameters) (*estate.EstateMainInfo, error) {
	return nil, nil
}

func (h handler) GetEstateInfo(ctx context.Context, parameter *estate.GetEstateInfoParameter) (*estate.Estate, error) {
	return nil, nil
}

func (h handler) CreateEstate(ctx context.Context, e *estate.Estate) (*emptypb.Empty, error) {
	return nil, nil
}

func (h handler) DeleteEstate(ctx context.Context, parameter *estate.DeleteEstateParameter) (*emptypb.Empty, error) {
	return nil, nil
}
