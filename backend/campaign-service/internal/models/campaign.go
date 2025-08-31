package models

import "time"

type Campaign struct {
	ID           int
	CreatorID    string
	Title        string
	Description  string
	GoalAmount   float64
	RaisedAmount float64
	Status       string
	CreatedAt    time.Time
}
