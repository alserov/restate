package grpc

import (
	"context"
	"github.com/alserov/restate/estate/internal/log"
	"github.com/alserov/restate/estate/internal/service"
	"github.com/alserov/restate/estate/internal/utils"
	estate "github.com/alserov/restate/estate/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func RegisterHandler(srvc service.Service, l log.Logger) *grpc.Server {
	srvr := grpc.NewServer()
	estate.RegisterEstateServiceServer(srvr, &handler{srvc: srvc, logger: l})
	return srvr
}

type handler struct {
	estate.UnimplementedEstateServiceServer

	srvc   service.Service
	logger log.Logger

	conv utils.Converter
}

func (h *handler) GetEstateList(ctx context.Context, parameters *estate.GetListParameters) (*estate.EstateList, error) {
	ctx = log.WithLogger(ctx, h.logger)

	h.logger.Trace(ctx, "passed GetEstateList server layer")

	list, err := h.srvc.GetEstateList(ctx, h.conv.ToGetEstateListParameters(parameters))
	if err != nil {
		msg, st := utils.FromError(h.logger, err)
		return nil, status.Error(st, msg)
	}

	return h.conv.FromEstateList(list), nil
}

func (h *handler) GetEstateInfo(ctx context.Context, parameter *estate.GetEstateInfoParameter) (*estate.Estate, error) {
	ctx = log.WithLogger(ctx, h.logger)

	h.logger.Trace(ctx, "passed GetEstateInfo server layer")

	info, err := h.srvc.GetEstateInfo(ctx, parameter.Id)
	if err != nil {
		msg, st := utils.FromError(h.logger, err)
		return nil, status.Error(st, msg)
	}

	return h.conv.FromEstate(info), nil
}

func (h *handler) CreateEstate(ctx context.Context, e *estate.Estate) (*emptypb.Empty, error) {
	ctx = log.WithLogger(ctx, h.logger)

	h.logger.Trace(ctx, "passed CreateEstate server layer")

	err := h.srvc.CreateEstate(ctx, h.conv.ToEstate(e))
	if err != nil {
		msg, st := utils.FromError(h.logger, err)
		return nil, status.Error(st, msg)
	}

	return &emptypb.Empty{}, nil
}

func (h *handler) DeleteEstate(ctx context.Context, parameter *estate.DeleteEstateParameter) (*emptypb.Empty, error) {
	ctx = log.WithLogger(ctx, h.logger)

	h.logger.Trace(ctx, "passed DeleteEstate server layer")

	err := h.srvc.DeleteEstate(ctx, parameter.Id)
	if err != nil {
		msg, st := utils.FromError(h.logger, err)
		return nil, status.Error(st, msg)
	}

	return &emptypb.Empty{}, nil
}
