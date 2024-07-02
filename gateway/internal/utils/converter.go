package utils

import (
	estate "github.com/alserov/restate/estate/pkg/grpc"
	"github.com/alserov/restate/gateway/internal/models"
	meetings "github.com/alserov/restate/meetings/pkg/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Converter struct{}

func (Converter) FromEstateList(in *estate.EstateList) models.EstateList {
	res := make(models.EstateList, 0, len(in.List))

	for _, est := range in.List {
		res = append(res, models.EstateInfo{
			Id:        est.Id,
			Title:     est.Title,
			Country:   est.Country,
			City:      est.City,
			Price:     est.Price,
			MainImage: est.MainImage,
		})
	}

	return res
}

func (Converter) ToGetEstateInfoParameter(in string) *estate.GetEstateInfoParameter {
	return &estate.GetEstateInfoParameter{
		Id: in,
	}
}

func (Converter) ToDeleteEstateParameter(in string) *estate.DeleteEstateParameter {
	return &estate.DeleteEstateParameter{Id: in}
}

func (Converter) FromEstate(in *estate.Estate) models.Estate {
	return models.Estate{
		Id:          in.Id,
		Title:       in.Title,
		Country:     in.Country,
		City:        in.City,
		Price:       in.Price,
		MainImage:   in.MainImage,
		Description: in.Description,
		Street:      in.Street,
		Images:      in.Images,
		Square:      in.Square,
		Floor:       in.Floor,
	}
}

func (Converter) ToEstate(in models.Estate) *estate.Estate {
	return &estate.Estate{
		Id:          in.Id,
		Title:       in.Title,
		Country:     in.Country,
		City:        in.City,
		Price:       in.Price,
		MainImage:   in.MainImage,
		Description: in.Description,
		Street:      in.Street,
		Images:      in.Images,
		Square:      in.Square,
		Floor:       in.Floor,
	}
}

func (Converter) ToGetEstateListParameter(in models.GetEstateListParameter) *estate.GetListParameters {
	return &estate.GetListParameters{
		MinPrice: in.MinPrice,
		MaxPrice: in.MaxPrice,
		Square:   in.Square,
		Country:  in.Country,
		City:     in.City,
		Floor:    in.Floor,
	}
}

func (Converter) ToGetMeetingsByPhoneNumberParameter(in string) *meetings.GetMeetingsByPhoneNumberParameter {
	return &meetings.GetMeetingsByPhoneNumberParameter{PhoneNumber: in}
}

func (Converter) ToGetMeetingsByEstateIDParameter(in string) *meetings.GetMeetingsByEstateIDParameter {
	return &meetings.GetMeetingsByEstateIDParameter{Id: in}
}

func (Converter) FromMeetings(in *meetings.Meetings) models.Meetings {
	mtngs := make([]models.Meeting, 0, len(in.Meetings))
	for _, mtng := range in.Meetings {
		mtngs = append(mtngs, fromMeeting(mtng))
	}

	return mtngs
}

func (Converter) FromAvailableTimeList(in *meetings.AvailableTimeList) []time.Time {
	tStamps := make([]time.Time, 0, len(in.Timestamps))
	for _, tStmap := range in.Timestamps {
		tStamps = append(tStamps, tStmap.AsTime())
	}

	return tStamps
}

func (Converter) ToGetAvailableTimeParameter(in string) *meetings.GetAvailableTimeForMeetingParameter {
	return &meetings.GetAvailableTimeForMeetingParameter{EstateID: in}
}

func (Converter) ToMeeting(in models.Meeting) *meetings.Meeting {
	return &meetings.Meeting{
		Timestamp:    timestamppb.New(in.Timestamp),
		EstateID:     in.EstateID,
		VisitorPhone: in.VisitorPhone,
	}
}

func (Converter) ToCancelMeetingParameter(in models.CancelMeetingParameter) *meetings.CancelMeetingParameter {
	return &meetings.CancelMeetingParameter{
		Id:           in.ID,
		VisitorPhone: in.VisitorPhone,
	}
}

func fromMeeting(in *meetings.Meeting) models.Meeting {
	return models.Meeting{
		Id:           in.Id,
		Timestamp:    in.Timestamp.AsTime(),
		EstateID:     in.EstateID,
		VisitorPhone: in.VisitorPhone,
	}
}
