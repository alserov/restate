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

// GetMeetings godoc
// @Summary      GetMeetings
// @Tags         meetings
// @Accept       json
// @Produce      json
// @Param        estateID   query    string  false  "estate id"
// @Param        phoneNumber   query    string  false  "phone number id"
// @Success      200  {array}   models.Meeting
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /meetings/list [get]
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
	} else {
		return utils.NewError("neither phone number, nor estate id provided", utils.InvalidData)
	}

	_ = c.JSON(http.StatusOK, mtngs)

	return nil
}

// GetAvailableTime godoc
// @Summary      GetAvailableTime
// @Tags         meetings
// @Accept       json
// @Produce      json
// @Param        estateID   query    string  true  "estate id"
// @Success      200  {array}   time.Time
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /meetings/available [get]
func (mh *MeetingsHandler) GetAvailableTime(c echo.Context) error {
	estateID := c.QueryParam("estateID")

	tStamps, err := mh.meetingsClient.GetAvailableTime(wrappers.Ctx(c), estateID)
	if err != nil {
		return fmt.Errorf("failed to get available timestamps: %w", err)
	}

	_ = c.JSON(http.StatusOK, tStamps)

	return nil
}

// ArrangeMeeting godoc
// @Summary      ArrangeMeeting
// @Tags         meetings
// @Accept       json
// @Produce      json
// @Param        input   body    models.Meeting  true  "meeting info"
// @Success      201  {array}   int
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /meetings/arrange [post]
func (mh *MeetingsHandler) ArrangeMeeting(c echo.Context) error {
	var mtng models.Meeting
	if err := c.Bind(&mtng); err != nil {
		return utils.NewError(err.Error(), utils.InvalidData)
	}

	err := mh.meetingsClient.ArrangeMeeting(wrappers.Ctx(c), mtng)
	if err != nil {
		return fmt.Errorf("failed to arrange meeting: %w", err)
	}

	_ = c.JSON(http.StatusCreated, nil)

	return nil
}

// CancelMeeting godoc
// @Summary      CancelMeeting
// @Tags         meetings
// @Accept       json
// @Produce      json
// @Param        input   body    models.CancelMeetingParameter  true  "meeting info"
// @Success      201  {array}   int
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /meetings/cancel [put]
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
