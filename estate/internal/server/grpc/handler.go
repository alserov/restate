package grpc

import (
	"context"
	"github.com/alserov/restate/estate/internal/cache"
	"github.com/alserov/restate/estate/internal/log"
	"github.com/alserov/restate/estate/internal/metrics"
	middleware "github.com/alserov/restate/estate/internal/middleware/grpc"
	"github.com/alserov/restate/estate/internal/service"
	"github.com/alserov/restate/estate/internal/service/models"
	"github.com/alserov/restate/estate/internal/utils"
	estate "github.com/alserov/restate/estate/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func RegisterHandler(srvc service.Service, cache cache.Cache, metr metrics.Metrics, l log.Logger) *grpc.Server {
	srvr := grpc.NewServer(
		middleware.ChainUnaryServer(
			middleware.WithWrappers(utils.WithLogger(l), utils.WithIdempotencyKey),
			middleware.WithRecover(),
			middleware.WithRequestObserver(metr),
			middleware.WithErrorHandler(),
		),
	)
	estate.RegisterEstateServiceServer(srvr, &handler{srvc: srvc, metr: metr, cache: cache})
	return srvr
}

type handler struct {
	estate.UnimplementedEstateServiceServer

	srvc service.Service

	conv utils.Converter

	metr metrics.Metrics

	cache cache.Cache
}

func (h *handler) GetEstateList(ctx context.Context, parameters *estate.GetListParameters) (*estate.EstateList, error) {
	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "passed GetEstateList server layer")

	list, err := h.srvc.GetEstateList(ctx, h.conv.ToGetEstateListParameters(parameters))
	if err != nil {
		return nil, err
	}

	return h.conv.FromEstateList(list), nil
}

func (h *handler) GetEstateInfo(ctx context.Context, parameter *estate.GetEstateInfoParameter) (*estate.Estate, error) {
	var cached models.Estate
	if h.cache.Get(ctx, parameter.Id, &cached) {
		return h.conv.FromEstate(cached), nil
	}

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "passed GetEstateInfo server layer")

	info, err := h.srvc.GetEstateInfo(ctx, parameter.Id)
	if err != nil {
		return nil, err
	}

	return h.conv.FromEstate(info), nil
}

func (h *handler) CreateEstate(ctx context.Context, e *estate.Estate) (*emptypb.Empty, error) {
	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "passed CreateEstate server layer")

	err := h.srvc.CreateEstate(ctx, h.conv.ToEstate(e))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *handler) DeleteEstate(ctx context.Context, parameter *estate.DeleteEstateParameter) (*emptypb.Empty, error) {
	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "passed DeleteEstate server layer")

	err := h.srvc.DeleteEstate(ctx, parameter.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
