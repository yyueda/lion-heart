package models

import "time"

type Donation struct {
	ID         string
	CampaignID string
	DonorID    string
	Amount     float64
	CreatedAt  time.Time
}
