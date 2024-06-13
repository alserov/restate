package controller

import (
	"github.com/alserov/restate/gateway/internal/log"
	"github.com/alserov/restate/gateway/internal/middleware"
	"github.com/labstack/echo/v4"
)

type Controller interface {
	SetupRoutes()
}

func NewController(app *echo.Echo, lg log.Logger) *controller {
	return &controller{
		app: app,
		lg:  lg,
	}
}

type controller struct {
	app *echo.Echo
	lg  log.Logger

	EstateHandler   *EstateHandler
	MeetingsHandler *MeetingsHandler
}

func (c *controller) SetupRoutes() {
	v1 := c.app.Group("/v1", middleware.WithLogger(c.lg), middleware.WithErrorHandler)

	estate := v1.Group("/estate")
	meetings := v1.Group("/meetings")

	// GET
	estate.GET("/list", c.EstateHandler.GetList)
	estate.GET("/info", c.EstateHandler.GetInfo)

	meetings.GET("/meetings", c.MeetingsHandler.GetMeetings)
	meetings.GET("/available", c.MeetingsHandler.GetAvailableTime)

	// POST
	estate.POST("/new", c.EstateHandler.CreateEstate)

	meetings.POST("/arrange", c.MeetingsHandler.ArrangeMeeting)

	// DELETE
	estate.DELETE("/delete", c.EstateHandler.DeleteEstate)

	// PUT

	meetings.PUT("/cancel", c.MeetingsHandler.CancelMeeting)
}
