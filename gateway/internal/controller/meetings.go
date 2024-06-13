package controller

import "github.com/labstack/echo/v4"

type MeetingsHandler struct {
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
