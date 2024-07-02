package models

import "time"

type (
	Meetings            []Meeting
	AvailableTimestamps []time.Time

	Meeting struct {
		Id           string    `json:"id" swaggerignore:"true"`
		Timestamp    time.Time `json:"timestamp"`
		EstateID     string    `json:"estateID"`
		VisitorPhone string    `json:"visitorPhone"`
	}

	CancelMeetingParameter struct {
		ID           string `json:"id"`
		VisitorPhone string `json:"visitorPhone"`
	}
)
