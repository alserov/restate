package controller

import (
	"fmt"
	"github.com/alserov/restate/gateway/internal/clients"
	"github.com/alserov/restate/gateway/internal/log"
	"github.com/alserov/restate/gateway/internal/metrics"
	"github.com/alserov/restate/gateway/internal/middleware/wrappers"
	"github.com/alserov/restate/gateway/internal/models"
	"github.com/alserov/restate/gateway/internal/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type EstateHandler struct {
	estateClient clients.EstateClient

	logger log.Logger

	metr metrics.Metrics
}

func NewEstateHandler(cl clients.EstateClient, metr metrics.Metrics, logger log.Logger) *EstateHandler {
	return &EstateHandler{
		estateClient: cl,
		logger:       logger,
		metr:         metr,
	}
}

// GetList godoc
// @Summary      GetList
// @Tags         estate
// @Accept       json
// @Produce      json
// @Param        country   query      string  false  "country"
// @Param        city   query      string  false  "city"
// @Param        floor   path      int  true  "floor"
// @Param        square   path      int  true  "square"
// @Param        minPrice   path      int  true  "min price"
// @Param        maxPrice   path      int  true  "max price"
// @Success      200  {array}   models.EstateInfo
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /estate/list [get]
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

	list, err := eh.estateClient.GetList(wrappers.Ctx(c), param)
	if err != nil {
		return fmt.Errorf("failed to get estate list: %w", err)
	}

	_ = c.JSON(http.StatusOK, list)

	return nil
}

// GetInfo godoc
// @Summary      GetInfo
// @Tags         estate
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "estate id"
// @Success      200  {object}  models.Estate
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /estate/info/{id} [get]
func (eh *EstateHandler) GetInfo(c echo.Context) error {
	estateID := c.Param("id")

	info, err := eh.estateClient.GetInfo(wrappers.Ctx(c), estateID)
	if err != nil {
		return fmt.Errorf("failed to get estate info: %w", err)
	}

	_ = c.JSON(http.StatusOK, info)

	return nil
}

// CreateEstate godoc
// @Summary      CreateEstate
// @Tags         estate
// @Accept       json
// @Produce      json
// @Param        input   body      models.Estate  true  "estate"
// @Success      201  {object}  string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /estate/new [post]
func (eh *EstateHandler) CreateEstate(c echo.Context) error {
	var estate models.Estate
	if err := c.Bind(&estate); err != nil {
		return utils.NewError(err.Error(), utils.InvalidData)
	}

	err := eh.estateClient.CreateEstate(wrappers.Ctx(c), estate)
	if err != nil {
		return fmt.Errorf("failed to create estate: %w", err)
	}

	_ = c.JSON(http.StatusCreated, nil)

	return nil
}

// DeleteEstate godoc
// @Summary      DeleteEstate
// @Tags         estate
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "estate id"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /estate/delete [delete]
func (eh *EstateHandler) DeleteEstate(c echo.Context) error {
	estateID := c.Param("id")

	err := eh.estateClient.DeleteEstate(wrappers.Ctx(c), estateID)
	if err != nil {
		return fmt.Errorf("failed to delete estate: %w", err)
	}

	return nil
}
