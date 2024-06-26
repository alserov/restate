package models

type GetEstateListParameters struct {
	MinPrice float32 `json:"minPrice"`
	MaxPrice float32 `json:"maxPrice"`
	Square   float32 `json:"square"`
	Country  string  `json:"country"`
	City     string  `json:"city"`
	Floor    int32   `json:"floor"`

	Limit  uint32
	Offset uint32
}
