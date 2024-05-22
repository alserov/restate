package service

type Service interface {
}

var _ Service = &service{}

func NewService() *service {
	return &service{}
}

type service struct {
}
