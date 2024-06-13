package clients

import (
	estate "github.com/alserov/restate/estate/pkg/grpc"
	meetings "github.com/alserov/restate/meetings/pkg/grpc"
)

type EstateClient interface {
	GetList()
	GetInfo()
	CreateEstate()
	DeleteEstate()
}

func NewEstateClient(cl estate.EstateServiceClient) EstateClient {
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
