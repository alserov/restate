package models

import "time"

type (
	Meetings            []Meeting
	AvailableTimestamps []time.Time

	Meeting struct {
		Id           string    `json:"id" swaggerignore:"true"`
		DateTime     string    `json:"date" swagger-example:"2024-12-01 13:30"`
		Timestamp    time.Time `json:"-" swaggerignore:"true"`
		EstateID     string    `json:"estateID"`
		VisitorPhone string    `json:"visitorPhone"`
	}

	CancelMeetingParameter struct {
		ID           string `json:"id"`
		VisitorPhone string `json:"visitorPhone"`
	}
)
