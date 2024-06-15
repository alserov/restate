package controller

import (
	"fmt"
	"github.com/alserov/restate/gateway/internal/clients"
	"github.com/alserov/restate/gateway/internal/models"
	"github.com/alserov/restate/gateway/internal/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type EstateHandler struct {
	estateClient clients.EstateClient
}

func (eh *EstateHandler) GetList(c echo.Context) error {
	param := models.GetEstateListParameter{
		Country: c.QueryParam("country"),
		City:    c.QueryParam("city"),
	}

	floor, err := strconv.Atoi(c.QueryParam("floor"))
	if err == nil {
		param.Floor = int32(floor)
	}

	square, err := strconv.Atoi(c.QueryParam("square"))
	if err == nil {
		param.Square = float32(square)
	}

	minPrice, err := strconv.Atoi(c.QueryParam("minPrice"))
	if err == nil {
		param.MinPrice = float32(minPrice)
	}

	maxPrice, err := strconv.Atoi(c.QueryParam("maxPrice"))
	if err == nil {
		param.MaxPrice = float32(maxPrice)
	}

	list, err := eh.estateClient.GetList(c.Request().Context(), param)
	if err != nil {
		return fmt.Errorf("failed to get estate list: %w", err)
	}

	_ = c.JSON(http.StatusOK, list)

	return nil
}

func (eh *EstateHandler) GetInfo(c echo.Context) error {
	estateID := c.Param("id")

	info, err := eh.estateClient.GetInfo(c.Request().Context(), estateID)
	if err != nil {
		return fmt.Errorf("failed to get estate info: %w", err)
	}

	_ = c.JSON(http.StatusOK, info)

	return nil
}

func (eh *EstateHandler) CreateEstate(c echo.Context) error {
	var estate models.Estate
	if err := c.Bind(&estate); err != nil {
		return utils.NewError(err.Error(), utils.InvalidData)
	}

	err := eh.estateClient.CreateEstate(c.Request().Context(), estate)
	if err != nil {
		return fmt.Errorf("failed to create estate: %w", err)
	}

	_ = c.JSON(http.StatusCreated, nil)

	return nil
}

func (eh *EstateHandler) DeleteEstate(c echo.Context) error {
	estateID := c.Param("id")

	err := eh.estateClient.DeleteEstate(c.Request().Context(), estateID)
	if err != nil {
		return fmt.Errorf("failed to delete estate: %w", err)
	}

	return nil
}
