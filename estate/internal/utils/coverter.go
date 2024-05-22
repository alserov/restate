package utils

import (
	"github.com/alserov/restate/estate/internal/service/models"
	estate "github.com/alserov/restate/estate/pkg/grpc"
)

type Converter struct {
}

func (Converter) ToGetEstateListParameters(in *estate.GetListParameters) models.GetEstateListParameters {
	return models.GetEstateListParameters{}
}

func (Converter) FromEstate(in models.Estate) *estate.Estate {
	return &estate.Estate{
		Id:          &in.ID,
		Title:       in.Title,
		Description: in.Description,
		Price:       in.Price,
		Country:     in.Country,
		City:        in.City,
		Street:      in.Street,
		Images:      in.Images,
		MainImage:   in.MainImage,
		Square:      in.Square,
		Floor:       in.Floor,
	}
}

func (Converter) ToEstate(in *estate.Estate) models.Estate {
	return models.Estate{
		Title:       in.Title,
		Description: in.Description,
		Price:       in.Price,
		Country:     in.Country,
		City:        in.City,
		Street:      in.Street,
		Images:      in.Images,
		MainImage:   in.MainImage,
		Square:      in.Square,
		Floor:       in.Floor,
	}
}

func (Converter) FromEstateMainInfo(in models.EstateMainInfo) *estate.EstateMainInfo {
	return &estate.EstateMainInfo{
		Id:        in.ID,
		Title:     in.Title,
		Country:   in.Country,
		City:      in.City,
		Price:     in.Price,
		MainImage: in.MainImage,
	}
}

func (c Converter) FromEstateList(in []models.EstateMainInfo) *estate.EstateList {
	res := estate.EstateList{List: make([]*estate.EstateMainInfo, 0, len(in))}

	for _, e := range in {
		res.List = append(res.List, c.FromEstateMainInfo(e))
	}

	return &res
}
