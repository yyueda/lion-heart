package repository

import "github.com/yyueda/lion-heart/backend/donation-service/internal/models"

type DonationRepo interface {
	Create(d *models.Donation) error
	ListByCampaign(campaignID string) ([]*models.Donation, error)
	ListByUser(userID string) ([]*models.Donation, error)
}
