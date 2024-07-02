package models

import (
	"github.com/lib/pq"
)

type Estate struct {
	ID          string         `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Price       float32        `json:"price"`
	Country     string         `json:"country"`
	City        string         `json:"city"`
	Street      string         `json:"street"`
	Images      pq.StringArray `json:"images"`
	MainImage   string         `json:"mainImage" db:"main_image"`
	Square      float32        `json:"square"`
	Floor       int32          `json:"floor"`
}

type EstateMainInfo struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Country   string  `json:"country"`
	City      string  `json:"city"`
	Price     float32 `json:"price"`
	MainImage string  `json:"mainImage" db:"main_image"`
}
