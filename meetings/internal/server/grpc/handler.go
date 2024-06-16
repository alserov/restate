package grpc

import (
	"context"
	"github.com/alserov/restate/meetings/internal/log"
	"github.com/alserov/restate/meetings/internal/service"
	"github.com/alserov/restate/meetings/internal/utils"
	meetings "github.com/alserov/restate/meetings/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func RegisterHandler(srvc service.Service, l log.Logger) *grpc.Server {
	srvr := grpc.NewServer()
	meetings.RegisterMeetingsServiceServer(srvr, &handler{srvc: srvc, logger: l})
	return srvr
}

type handler struct {
	meetings.UnimplementedMeetingsServiceServer

	srvc   service.Service
	logger log.Logger

	conv utils.Converter
}

func (h *handler) GetMeetingsByEstateID(ctx context.Context, parameter *meetings.GetAvailableTimeForMeetingParameter) (*meetings.Meetings, error) {
	mtngs, err := h.srvc.GetMeetingsByEstateID(ctx, parameter.EstateID)
	if err != nil {
		msg, st := utils.FromError(h.logger, err)
		return nil, status.Error(st, msg)
	}

	return h.conv.FromMeetings(mtngs), nil
}

func (h *handler) GetMeetingsByPhoneNumber(ctx context.Context, parameter *meetings.GetMeetingsByPhoneNumberParameter) (*meetings.Meetings, error) {
	mtngs, err := h.srvc.GetMeetingsByPhoneNumber(ctx, parameter.PhoneNumber)
	if err != nil {
		msg, st := utils.FromError(h.logger, err)
		return nil, status.Error(st, msg)
	}

	return h.conv.FromMeetings(mtngs), nil
}

func (h *handler) ArrangeMeeting(ctx context.Context, meeting *meetings.Meeting) (*emptypb.Empty, error) {
	err := h.srvc.ArrangeMeeting(ctx, h.conv.ToMeeting(meeting))
	if err != nil {
		msg, st := utils.FromError(h.logger, err)
		return nil, status.Error(st, msg)
	}

	return &emptypb.Empty{}, nil
}

func (h *handler) CancelMeeting(ctx context.Context, parameter *meetings.CancelMeetingParameter) (*emptypb.Empty, error) {
	err := h.srvc.CancelMeeting(ctx, h.conv.ToCancelMeetingParameter(parameter))
	if err != nil {
		msg, st := utils.FromError(h.logger, err)
		return nil, status.Error(st, msg)
	}

	return &emptypb.Empty{}, nil
}

func (h *handler) GetAvailableTimeForMeeting(ctx context.Context, parameter *meetings.GetAvailableTimeForMeetingParameter) (*meetings.AvailableTimeList, error) {
	tStamps, err := h.srvc.GetAvailableTimeForMeeting(ctx, parameter.EstateID)
	if err != nil {
		msg, st := utils.FromError(h.logger, err)
		return nil, status.Error(st, msg)
	}

	return h.conv.FromTimestamps(tStamps), nil
}
