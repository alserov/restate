package clients

import (
	"context"
	estate "github.com/alserov/restate/estate/pkg/grpc"
	"github.com/alserov/restate/gateway/internal/models"
	"github.com/alserov/restate/gateway/internal/utils"
	meetings "github.com/alserov/restate/meetings/pkg/grpc"
	"time"
)

type EstateClient interface {
	GetList(ctx context.Context, param models.GetEstateListParameter) (models.EstateList, error)
	GetInfo(ctx context.Context, estateID string) (models.Estate, error)
	CreateEstate(ctx context.Context, est models.Estate) error
	DeleteEstate(ctx context.Context, id string) error
}

func NewEstateClient(cl estate.EstateServiceClient) EstateClient {
	return &estateClient{cl: cl}
}

type estateClient struct {
	cl estate.EstateServiceClient

	conv utils.Converter
}

func (e estateClient) GetList(ctx context.Context, param models.GetEstateListParameter) (models.EstateList, error) {
	list, err := e.cl.GetEstateList(ctx, e.conv.ToGetEstateListParameter(param))
	if err != nil {
		return nil, utils.FromGRPCError(err)
	}

	return e.conv.FromEstateList(list), nil
}

func (e estateClient) GetInfo(ctx context.Context, id string) (models.Estate, error) {
	info, err := e.cl.GetEstateInfo(ctx, e.conv.ToGetEstateInfoParameter(id))
	if err != nil {
		return models.Estate{}, utils.FromGRPCError(err)
	}

	return e.conv.FromEstate(info), nil
}

func (e estateClient) CreateEstate(ctx context.Context, est models.Estate) error {
	_, err := e.cl.CreateEstate(ctx, e.conv.ToEstate(est))
	if err != nil {
		return utils.FromGRPCError(err)
	}

	return nil
}

func (e estateClient) DeleteEstate(ctx context.Context, id string) error {
	_, err := e.cl.DeleteEstate(ctx, e.conv.ToDeleteEstateParameter(id))
	if err != nil {
		return utils.FromGRPCError(err)
	}

	return nil
}

type MeetingsClient interface {
	GetMeetingsByPhoneNumber(ctx context.Context, phone string) (models.Meetings, error)
	GetMeetingsByEstateID(ctx context.Context, estateID string) (models.Meetings, error)
	GetAvailableTime(ctx context.Context, estateID string) ([]time.Time, error)
	ArrangeMeeting(ctx context.Context, mtng models.Meeting) error
	CancelMeeting(ctx context.Context, par models.CancelMeetingParameter) error
}

func NewMeetingsClient(cl meetings.MeetingsServiceClient) MeetingsClient {
	return &meetingsClient{
		cl: cl,
	}
}

type meetingsClient struct {
	cl meetings.MeetingsServiceClient

	conv utils.Converter
}

func (m meetingsClient) GetMeetingsByPhoneNumber(ctx context.Context, phone string) (models.Meetings, error) {
	mtngs, err := m.cl.GetMeetingsByPhoneNumber(ctx, m.conv.ToGetMeetingsByPhoneNumberParameter(phone))
	if err != nil {
		return nil, utils.FromGRPCError(err)
	}

	return m.conv.FromMeetings(mtngs), nil
}

func (m meetingsClient) GetMeetingsByEstateID(ctx context.Context, estateID string) (models.Meetings, error) {
	mtngs, err := m.cl.GetMeetingsByEstateID(ctx, m.conv.ToGetMeetingsByEstateIDParameter(estateID))
	if err != nil {
		return nil, utils.FromGRPCError(err)
	}

	return m.conv.FromMeetings(mtngs), nil
}

func (m meetingsClient) GetAvailableTime(ctx context.Context, estateID string) ([]time.Time, error) {
	tStamps, err := m.cl.GetAvailableTimeForMeeting(ctx, m.conv.ToGetAvailableTimeParameter(estateID))
	if err != nil {
		return nil, utils.FromGRPCError(err)
	}

	return m.conv.FromAvailableTimeList(tStamps), nil
}

func (m meetingsClient) ArrangeMeeting(ctx context.Context, mtng models.Meeting) error {
	_, err := m.cl.ArrangeMeeting(ctx, m.conv.ToMeeting(mtng))
	if err != nil {
		return utils.FromGRPCError(err)
	}

	return nil
}

func (m meetingsClient) CancelMeeting(ctx context.Context, par models.CancelMeetingParameter) error {
	_, err := m.cl.CancelMeeting(ctx, m.conv.ToCancelMeetingParameter(par))
	if err != nil {
		return utils.FromGRPCError(err)
	}

	return nil
}
