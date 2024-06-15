package controller

import (
	estate "github.com/alserov/restate/estate/pkg/grpc"
	"github.com/alserov/restate/gateway/internal/clients"
	"github.com/alserov/restate/gateway/internal/log"
	"github.com/alserov/restate/gateway/internal/metrics"
	"github.com/alserov/restate/gateway/internal/middleware"
	meetings "github.com/alserov/restate/meetings/pkg/grpc"
	"github.com/labstack/echo/v4"
)

type Controller interface {
	SetupRoutes()
}

type Clients struct {
	Estate   estate.EstateServiceClient
	Meetings meetings.MeetingsServiceClient
}

func NewController(app *echo.Echo, lg log.Logger, metr metrics.Metrics, cls *Clients) *controller {
	return &controller{
		app:  app,
		lg:   lg,
		metr: metr,

		EstateHandler:   &EstateHandler{estateClient: clients.NewEstateClient(cls.Estate)},
		MeetingsHandler: &MeetingsHandler{meetingsClient: clients.NewMeetingsClient(cls.Meetings)},
	}
}

type controller struct {
	app  *echo.Echo
	lg   log.Logger
	metr metrics.Metrics

	EstateHandler   *EstateHandler
	MeetingsHandler *MeetingsHandler
}

func (c *controller) SetupRoutes() {
	v1 := c.app.Group("/v1", middleware.WithLogger(c.lg), middleware.WithRequestObserver(c.metr), middleware.WithErrorHandler)

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
