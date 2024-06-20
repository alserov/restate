package grpc

import (
	"context"
	"github.com/alserov/restate/meetings/internal/log"
	"github.com/alserov/restate/meetings/internal/metrics"
	"github.com/alserov/restate/meetings/internal/middleware"
	"github.com/alserov/restate/meetings/internal/service"
	"github.com/alserov/restate/meetings/internal/utils"
	meetings "github.com/alserov/restate/meetings/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

func RegisterHandler(srvc service.Service, m metrics.Metrics, l log.Logger) *grpc.Server {
	srvr := grpc.NewServer(
		middleware.WithWrapper(
			middleware.WithLogger,
			middleware.WithIdempotencyKey,
		),
		middleware.WithRequestObserver(m),
	)
	meetings.RegisterMeetingsServiceServer(srvr, &handler{srvc: srvc, logger: l, metr: m})
	return srvr
}

type handler struct {
	meetings.UnimplementedMeetingsServiceServer

	srvc   service.Service
	logger log.Logger

	conv utils.Converter

	metr metrics.Metrics
}

func (h *handler) GetMeetingsByEstateID(ctx context.Context, parameter *meetings.GetAvailableTimeForMeetingParameter) (*meetings.Meetings, error) {
	var (
		st    *int
		start = time.Now()
	)
	defer func() {
		if err := h.metr.ObserveRequest(ctx, *st, time.Since(start), "GetMeetingsByEstateID"); err != nil {
			h.logger.Warn("failed to observe request", log.WithData("warn", err.Error()))
		}
	}()

	h.logger.Trace(middleware.ExtractIdempotencyKey(ctx), "passed GetEstateList server layer")

	mtngs, err := h.srvc.GetMeetingsByEstateID(ctx, parameter.EstateID)
	if err != nil {
		msg, st := utils.FromError(h.logger, err)
		return nil, status.Error(st, msg)
	}

	return h.conv.FromMeetings(mtngs), nil
}

func (h *handler) GetMeetingsByPhoneNumber(ctx context.Context, parameter *meetings.GetMeetingsByPhoneNumberParameter) (*meetings.Meetings, error) {
	var (
		st    *int
		start = time.Now()
	)
	defer func() {
		if err := h.metr.ObserveRequest(ctx, *st, time.Since(start), "GetMeetingsByPhoneNumber"); err != nil {
			h.logger.Warn("failed to observe request", log.WithData("warn", err.Error()))
		}
	}()

	h.logger.Trace(middleware.ExtractIdempotencyKey(ctx), "passed GetMeetingsByPhoneNumber server layer")

	mtngs, err := h.srvc.GetMeetingsByPhoneNumber(ctx, parameter.PhoneNumber)
	if err != nil {
		msg, st := utils.FromError(h.logger, err)
		return nil, status.Error(st, msg)
	}

	return h.conv.FromMeetings(mtngs), nil
}

func (h *handler) ArrangeMeeting(ctx context.Context, meeting *meetings.Meeting) (*emptypb.Empty, error) {
	var (
		st    *int
		start = time.Now()
	)
	defer func() {
		if err := h.metr.ObserveRequest(ctx, *st, time.Since(start), "ArrangeMeeting"); err != nil {
			h.logger.Warn("failed to observe request", log.WithData("warn", err.Error()))
		}
	}()

	h.logger.Trace(middleware.ExtractIdempotencyKey(ctx), "passed ArrangeMeeting server layer")

	err := h.srvc.ArrangeMeeting(ctx, h.conv.ToMeeting(meeting))
	if err != nil {
		msg, st := utils.FromError(h.logger, err)
		return nil, status.Error(st, msg)
	}

	return &emptypb.Empty{}, nil
}

func (h *handler) CancelMeeting(ctx context.Context, parameter *meetings.CancelMeetingParameter) (*emptypb.Empty, error) {
	var (
		st    *int
		start = time.Now()
	)
	defer func() {
		if err := h.metr.ObserveRequest(ctx, *st, time.Since(start), "CancelMeeting"); err != nil {
			h.logger.Warn("failed to observe request", log.WithData("warn", err.Error()))
		}
	}()

	h.logger.Trace(middleware.ExtractIdempotencyKey(ctx), "passed CancelMeeting server layer")

	err := h.srvc.CancelMeeting(ctx, h.conv.ToCancelMeetingParameter(parameter))
	if err != nil {
		msg, st := utils.FromError(h.logger, err)
		return nil, status.Error(st, msg)
	}

	return &emptypb.Empty{}, nil
}

func (h *handler) GetAvailableTimeForMeeting(ctx context.Context, parameter *meetings.GetAvailableTimeForMeetingParameter) (*meetings.AvailableTimeList, error) {
	var (
		st    *int
		start = time.Now()
	)
	defer func() {
		if err := h.metr.ObserveRequest(ctx, *st, time.Since(start), "GetAvailableTimeForMeeting"); err != nil {
			h.logger.Warn("failed to observe request", log.WithData("warn", err.Error()))
		}
	}()

	h.logger.Trace(middleware.ExtractIdempotencyKey(ctx), "passed GetAvailableTimeForMeeting server layer")

	tStamps, err := h.srvc.GetAvailableTimeForMeeting(ctx, parameter.EstateID)
	if err != nil {
		msg, st := utils.FromError(h.logger, err)
		return nil, status.Error(st, msg)
	}

	return h.conv.FromTimestamps(tStamps), nil
}
