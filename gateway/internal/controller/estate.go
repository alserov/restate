package controller

import (
	"github.com/alserov/restate/gateway/internal/clients"
	"github.com/labstack/echo/v4"
)

type EstateHandler struct {
	estateClient clients.EstateClient
}

func (eh *EstateHandler) GetList(c echo.Context) error {
	return nil
}

func (eh *EstateHandler) GetInfo(c echo.Context) error {
	return nil
}

func (eh *EstateHandler) CreateEstate(e echo.Context) error {
	return nil
}

func (eh *EstateHandler) DeleteEstate(e echo.Context) error {
	return nil
}
