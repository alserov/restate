package controller

import (
	"github.com/alserov/restate/gateway/internal/clients"
	"github.com/labstack/echo/v4"
)

type MeetingsHandler struct {
	meetingsClient clients.MeetingsClient
}

func (mh *MeetingsHandler) GetMeetings(c echo.Context) error {
	return nil
}

func (mh *MeetingsHandler) GetAvailableTime(c echo.Context) error {
	return nil
}

func (mh *MeetingsHandler) ArrangeMeeting(c echo.Context) error {
	return nil
}

func (mh *MeetingsHandler) CancelMeeting(c echo.Context) error {
	return nil
}