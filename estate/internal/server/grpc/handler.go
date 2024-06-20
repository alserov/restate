package grpc

import (
	"context"
	"github.com/alserov/restate/estate/internal/metrics"
	"github.com/alserov/restate/estate/internal/middleware"
	"github.com/alserov/restate/estate/internal/middleware/wrappers"
	"github.com/alserov/restate/estate/internal/service"
	"github.com/alserov/restate/estate/internal/utils"
	estate "github.com/alserov/restate/estate/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func RegisterHandler(srvc service.Service, metr metrics.Metrics) *grpc.Server {
	srvr := grpc.NewServer(
		middleware.WithWrappers(wrappers.WithLogger, wrappers.WithIdempotencyKey),
		middleware.WithRequestObserver(metr),
		middleware.WithErrorHandler(),
	)
	estate.RegisterEstateServiceServer(srvr, &handler{srvc: srvc, metr: metr})
	return srvr
}

type handler struct {
	estate.UnimplementedEstateServiceServer

	srvc service.Service

	conv utils.Converter

	metr metrics.Metrics
}

func (h *handler) GetEstateList(ctx context.Context, parameters *estate.GetListParameters) (*estate.EstateList, error) {
	list, err := h.srvc.GetEstateList(ctx, h.conv.ToGetEstateListParameters(parameters))
	if err != nil {
		return nil, err
	}

	return h.conv.FromEstateList(list), nil
}

func (h *handler) GetEstateInfo(ctx context.Context, parameter *estate.GetEstateInfoParameter) (*estate.Estate, error) {
	info, err := h.srvc.GetEstateInfo(ctx, parameter.Id)
	if err != nil {
		return nil, err
	}

	return h.conv.FromEstate(info), nil
}

func (h *handler) CreateEstate(ctx context.Context, e *estate.Estate) (*emptypb.Empty, error) {
	err := h.srvc.CreateEstate(ctx, h.conv.ToEstate(e))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *handler) DeleteEstate(ctx context.Context, parameter *estate.DeleteEstateParameter) (*emptypb.Empty, error) {
	err := h.srvc.DeleteEstate(ctx, parameter.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
