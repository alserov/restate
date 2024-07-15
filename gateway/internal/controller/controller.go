package controller

import (
	estate "github.com/alserov/restate/estate/pkg/grpc"
	"github.com/alserov/restate/gateway/internal/clients"
	"github.com/alserov/restate/gateway/internal/log"
	"github.com/alserov/restate/gateway/internal/metrics"
	"github.com/alserov/restate/gateway/internal/middleware"
	"github.com/alserov/restate/gateway/internal/middleware/wrappers"
	meetings "github.com/alserov/restate/meetings/pkg/grpc"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Controller interface {
	SetupRoutes()
}

type Clients struct {
	Estate   estate.EstateServiceClient
	Meetings meetings.MeetingsServiceClient
}

func NewController(app *echo.Echo, metr metrics.Metrics, lg log.Logger, cls *Clients) *controller {
	return &controller{
		app:  app,
		lg:   lg,
		metr: metr,

		EstateHandler:   &EstateHandler{estateClient: clients.NewEstateClient(cls.Estate), logger: lg},
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
	v1 := c.app.Group("/v1",
		// wrap request context
		wrappers.WithIdempotencyKey,
		wrappers.WithLogger(c.lg),

		middleware.WithRequestObserver(c.metr),
		middleware.WithErrorHandler,
		middleware.WithRateLimiter(10_000),
	)

	estate := v1.Group("/estate")
	meetings := v1.Group("/meetings")

	// GET
	c.app.GET("/swagger/*", echoSwagger.EchoWrapHandler())

	estate.GET("/list", c.EstateHandler.GetList)
	estate.GET("/info/:id", c.EstateHandler.GetInfo)

	meetings.GET("/list", c.MeetingsHandler.GetMeetings)
	meetings.GET("/available", c.MeetingsHandler.GetAvailableTime)

	// POST
	estate.POST("/new", c.EstateHandler.CreateEstate)

	meetings.POST("/arrange", c.MeetingsHandler.ArrangeMeeting)

	// DELETE
	estate.DELETE("/delete/:id", c.EstateHandler.DeleteEstate)

	// PUT
	meetings.PUT("/cancel", c.MeetingsHandler.CancelMeeting)
}
