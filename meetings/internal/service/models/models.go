package models

import "time"

type Meeting struct {
	ID           string    `json:"id"`
	Timestamp    time.Time `json:"timestamp"`
	EstateID     string    `json:"estateID" db:"estate_id"`
	VisitorPhone string    `json:"visitorPhone" db:"visitor_phone"`
}

const (
	MinMeetingTimestamp = 9
	MaxMeetingTimestamp = 20
)
