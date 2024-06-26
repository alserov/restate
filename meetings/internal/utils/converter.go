package utils

import (
	"github.com/alserov/restate/meetings/internal/service/models"
	meetings "github.com/alserov/restate/meetings/pkg/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Converter struct {
}

func (Converter) ToMeeting(in *meetings.Meeting) models.Meeting {
	return models.Meeting{
		ID:           in.Id,
		Timestamp:    in.Timestamp.AsTime(),
		EstateID:     in.EstateID,
		VisitorPhone: in.VisitorPhone,
	}
}

func (Converter) ToCancelMeetingParameter(in *meetings.CancelMeetingParameter) models.CancelMeetingParameter {
	return models.CancelMeetingParameter{
		ID:           in.Id,
		VisitorPhone: in.VisitorPhone,
	}
}

func (Converter) FromTimestamps(in []time.Time) *meetings.AvailableTimeList {
	tStamps := make([]*timestamppb.Timestamp, 0, len(in))

	for _, t := range in {
		tStamps = append(tStamps, timestamppb.New(t))
	}

	return &meetings.AvailableTimeList{Timestamps: tStamps}
}

func (Converter) FromMeetings(in []models.Meeting) *meetings.Meetings {
	mtngs := make([]*meetings.Meeting, 0, len(in))

	for _, mtng := range in {
		mtngs = append(mtngs, &meetings.Meeting{
			Id:           mtng.ID,
			Timestamp:    timestamppb.New(mtng.Timestamp),
			EstateID:     mtng.EstateID,
			VisitorPhone: mtng.VisitorPhone,
		})
	}

	return &meetings.Meetings{Meetings: mtngs}
}
