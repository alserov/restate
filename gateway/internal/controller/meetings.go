package controller

import (
	"fmt"
	"github.com/alserov/restate/gateway/internal/clients"
	"github.com/alserov/restate/gateway/internal/middleware/wrappers"
	"github.com/alserov/restate/gateway/internal/models"
	"github.com/alserov/restate/gateway/internal/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MeetingsHandler struct {
	meetingsClient clients.MeetingsClient
}

func (mh *MeetingsHandler) GetMeetings(c echo.Context) error {
	var (
		err   error
		mtngs models.Meetings
	)

	if estateID := c.QueryParam("estateID"); estateID != "" {
		mtngs, err = mh.meetingsClient.GetMeetingsByEstateID(wrappers.Ctx(c), estateID)
		if err != nil {
			return fmt.Errorf("failed to get meetings: %w", err)
		}
	} else if phoneNumber := c.QueryParam("phoneNumber"); phoneNumber != "" {
		mtngs, err = mh.meetingsClient.GetMeetingsByPhoneNumber(wrappers.Ctx(c), phoneNumber)
		if err != nil {
			return fmt.Errorf("failed to get meetings: %w", err)
		}
	}

	_ = c.JSON(http.StatusOK, mtngs)

	return nil
}

func (mh *MeetingsHandler) GetAvailableTime(c echo.Context) error {
	estateID := c.QueryParam("estateID")

	tStamps, err := mh.meetingsClient.GetAvailableTime(wrappers.Ctx(c), estateID)
	if err != nil {
		return fmt.Errorf("failed to get available timestamps: %w", err)
	}

	_ = c.JSON(http.StatusOK, tStamps)

	return nil
}

func (mh *MeetingsHandler) ArrangeMeeting(c echo.Context) error {
	var mtng models.Meeting
	if err := c.Bind(&mtng); err != nil {
		return utils.NewError(err.Error(), utils.InvalidData)
	}

	err := mh.meetingsClient.ArrangeMeeting(wrappers.Ctx(c), mtng)
	if err != nil {
		return fmt.Errorf("failed to arrange meeting: %w", err)
	}

	return nil
}

func (mh *MeetingsHandler) CancelMeeting(c echo.Context) error {
	var par models.CancelMeetingParameter
	if err := c.Bind(&par); err != nil {
		return utils.NewError(err.Error(), utils.InvalidData)
	}

	err := mh.meetingsClient.CancelMeeting(wrappers.Ctx(c), par)
	if err != nil {
		return fmt.Errorf("failed to arrange meeting: %w", err)
	}

	return nil
}
