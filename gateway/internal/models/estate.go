package models

type (
	GetEstateListParameter struct {
		MinPrice float32
		MaxPrice float32
		Square   float32
		Country  string
		City     string
		Floor    int32
	}

	Estate struct {
		Id          *string
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
		Id        string
		Title     string
		Country   string
		City      string
		Price     float32
		MainImage string
	}

	EstateList []EstateInfo
)
