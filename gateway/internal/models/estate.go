package models

type (
	GetEstateListParameter struct {
		MinPrice float32
		MaxPrice float32
		Square   float32
		Country  string
		City     string
		Floor    int32

		Limit  int
		Offset int
	}

	Estate struct {
		Id          *string `swaggeringore:"true"`
		Title       string
		Description string
		Price       float32
		Country     string
		City        string
		Street      string
		Images      []string
		MainImage   string
		Square      float32
		Floor       int32
	}

	EstateInfo struct {
		Id        string  `json:"id"`
		Title     string  `json:"title"`
		Country   string  `json:"country"`
		City      string  `json:"city"`
		Price     float32 `json:"price"`
		MainImage string  `json:"mainImage"`
	}

	EstateList []EstateInfo
)
