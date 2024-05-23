package models

import "time"

type Meeting struct {
	ID           string    `json:"id"`
	Timestamp    time.Time `json:"timestamp"`
	EstateID     string    `json:"estateID"`
	VisitorPhone string    `json:"visitorPhone"`
}
