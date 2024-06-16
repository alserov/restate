package grpc

import (
	"context"
	"github.com/alserov/restate/estate/internal/log"
	"github.com/alserov/restate/estate/internal/metrics"
	"github.com/alserov/restate/estate/internal/service"
	"github.com/alserov/restate/estate/internal/utils"
	"github.com/alserov/restate/estate/internal/wrappers"
	estate "github.com/alserov/restate/estate/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

func RegisterHandler(srvc service.Service, metr metrics.Metrics, l log.Logger) *grpc.Server {
	srvr := grpc.NewServer(wrappers.WithGRPC(wrappers.WithLogger, wrappers.WithIdempotencyKey))

	estate.RegisterEstateServiceServer(srvr, &handler{srvc: srvc, logger: l, metr: metr})

	return srvr
}

type handler struct {
	estate.UnimplementedEstateServiceServer

	srvc   service.Service
	logger log.Logger

	conv utils.Converter

	metr metrics.Metrics
}

func (h *handler) GetEstateList(ctx context.Context, parameters *estate.GetListParameters) (*estate.EstateList, error) {
	var (
		st    *int
		start = time.Now()
	)
	defer func() {
		if err := h.metr.ObserveRequest(ctx, *st, time.Since(start), "GetEstateList"); err != nil {
			h.logger.Warn("failed to observe request", log.WithData("warn", err.Error()))
		}
	}()

	h.logger.Trace(wrappers.ExtractIdempotencyKey(ctx), "passed GetEstateList server layer")

	list, err := h.srvc.GetEstateList(ctx, h.conv.ToGetEstateListParameters(parameters))
	if err != nil {
		msg, st := utils.FromError(h.logger, err)
		return nil, status.Error(st, msg)
	}

	return h.conv.FromEstateList(list), nil
}

func (h *handler) GetEstateInfo(ctx context.Context, parameter *estate.GetEstateInfoParameter) (*estate.Estate, error) {
	var (
		st    *int
		start = time.Now()
	)
	defer func() {
		if err := h.metr.ObserveRequest(ctx, *st, time.Since(start), "GetEstateInfo"); err != nil {
			h.logger.Warn("failed to observe request", log.WithData("warn", err.Error()))
		}
	}()

	h.logger.Trace(wrappers.ExtractIdempotencyKey(ctx), "passed GetEstateInfo server layer")

	info, err := h.srvc.GetEstateInfo(ctx, parameter.Id)
	if err != nil {
		msg, st := utils.FromError(h.logger, err)
		return nil, status.Error(st, msg)
	}

	return h.conv.FromEstate(info), nil
}

func (h *handler) CreateEstate(ctx context.Context, e *estate.Estate) (*emptypb.Empty, error) {
	var (
		st    *int
		start = time.Now()
	)
	defer func() {
		if err := h.metr.ObserveRequest(ctx, *st, time.Since(start), "CreateEstate"); err != nil {
			h.logger.Warn("failed to observe request", log.WithData("warn", err.Error()))
		}
	}()

	h.logger.Trace(wrappers.ExtractIdempotencyKey(ctx), "passed CreateEstate server layer")

	err := h.srvc.CreateEstate(ctx, h.conv.ToEstate(e))
	if err != nil {
		msg, st := utils.FromError(h.logger, err)
		return nil, status.Error(st, msg)
	}

	return &emptypb.Empty{}, nil
}

func (h *handler) DeleteEstate(ctx context.Context, parameter *estate.DeleteEstateParameter) (*emptypb.Empty, error) {
	var (
		st    *int
		start = time.Now()
	)

	defer func() {
		if err := h.metr.ObserveRequest(ctx, *st, time.Since(start), "DeleteEstate"); err != nil {
			h.logger.Warn("failed to observe request", log.WithData("warn", err.Error()))
		}
	}()

	h.logger.Trace(wrappers.ExtractIdempotencyKey(ctx), "passed DeleteEstate server layer")

	err := h.srvc.DeleteEstate(ctx, parameter.Id)
	if err != nil {
		msg, st := utils.FromError(h.logger, err)
		return nil, status.Error(st, msg)
	}

	return &emptypb.Empty{}, nil
}
