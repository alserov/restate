package grpc

import (
	"context"
	"github.com/alserov/restate/meetings/internal/log"
	"github.com/alserov/restate/meetings/internal/metrics"
	"github.com/alserov/restate/meetings/internal/middleware"
	"github.com/alserov/restate/meetings/internal/middleware/wrappers"
	"github.com/alserov/restate/meetings/internal/service"
	"github.com/alserov/restate/meetings/internal/utils"
	meetings "github.com/alserov/restate/meetings/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func RegisterHandler(srvc service.Service, m metrics.Metrics, l log.Logger) *grpc.Server {
	srvr := grpc.NewServer(
		middleware.WithWrappers(
			wrappers.WithLogger,
			wrappers.WithIdempotencyKey,
		),
		middleware.WithRequestObserver(m),
		middleware.WithErrorHandler(),
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
	h.logger.Trace(wrappers.ExtractIdempotencyKey(ctx), "passed GetEstateList server layer")

	mtngs, err := h.srvc.GetMeetingsByEstateID(ctx, parameter.EstateID)
	if err != nil {
		return nil, err
	}

	return h.conv.FromMeetings(mtngs), nil
}

func (h *handler) GetMeetingsByPhoneNumber(ctx context.Context, parameter *meetings.GetMeetingsByPhoneNumberParameter) (*meetings.Meetings, error) {
	h.logger.Trace(wrappers.ExtractIdempotencyKey(ctx), "passed GetMeetingsByPhoneNumber server layer")

	mtngs, err := h.srvc.GetMeetingsByPhoneNumber(ctx, parameter.PhoneNumber)
	if err != nil {
		return nil, err
	}

	return h.conv.FromMeetings(mtngs), nil
}

func (h *handler) ArrangeMeeting(ctx context.Context, meeting *meetings.Meeting) (*emptypb.Empty, error) {
	h.logger.Trace(wrappers.ExtractIdempotencyKey(ctx), "passed ArrangeMeeting server layer")

	err := h.srvc.ArrangeMeeting(ctx, h.conv.ToMeeting(meeting))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *handler) CancelMeeting(ctx context.Context, parameter *meetings.CancelMeetingParameter) (*emptypb.Empty, error) {
	h.logger.Trace(wrappers.ExtractIdempotencyKey(ctx), "passed CancelMeeting server layer")

	err := h.srvc.CancelMeeting(ctx, h.conv.ToCancelMeetingParameter(parameter))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *handler) GetAvailableTimeForMeeting(ctx context.Context, parameter *meetings.GetAvailableTimeForMeetingParameter) (*meetings.AvailableTimeList, error) {
	h.logger.Trace(wrappers.ExtractIdempotencyKey(ctx), "passed GetAvailableTimeForMeeting server layer")

	tStamps, err := h.srvc.GetAvailableTimeForMeeting(ctx, parameter.EstateID)
	if err != nil {
		return nil, err
	}

	return h.conv.FromTimestamps(tStamps), nil
}
