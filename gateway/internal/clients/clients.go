package clients

import (
	"context"
	estate "github.com/alserov/restate/estate/pkg/grpc"
	"github.com/alserov/restate/gateway/internal/models"
	"github.com/alserov/restate/gateway/internal/utils"
	meetings "github.com/alserov/restate/meetings/pkg/grpc"
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

func (e *estateClient) GetList(ctx context.Context, param models.GetEstateListParameter) (models.EstateList, error) {
	list, err := e.cl.GetEstateList(ctx, e.conv.ToGetEstateListParameter(param))
	if err != nil {
		return nil, utils.FromGRPCError(err)
	}

	return e.conv.FromEstateList(list), nil
}

func (e *estateClient) GetInfo(ctx context.Context, id string) (models.Estate, error) {
	info, err := e.cl.GetEstateInfo(ctx, e.conv.ToGetEstateInfoParameter(id))
	if err != nil {
		return models.Estate{}, utils.FromGRPCError(err)
	}

	return e.conv.FromEstate(info), nil
}

func (e *estateClient) CreateEstate(ctx context.Context, est models.Estate) error {
	_, err := e.cl.CreateEstate(ctx, e.conv.ToEstate(est))
	if err != nil {
		return utils.FromGRPCError(err)
	}

	return nil
}

func (e *estateClient) DeleteEstate(ctx context.Context, id string) error {
	_, err := e.cl.DeleteEstate(ctx, e.conv.ToDeleteEstateParameter(id))
	if err != nil {
		return utils.FromGRPCError(err)
	}

	return nil
}

type MeetingsClient interface {
	GetMeetings()
	GetAvailableTime()
	ArrangeMeeting()
	CancelMeeting()
}

func NewMeetingsClient(cl meetings.MeetingsServiceClient) MeetingsClient {
	return nil
}
