package utils

import (
	estate "github.com/alserov/restate/estate/pkg/grpc"
	"github.com/alserov/restate/gateway/internal/models"
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
