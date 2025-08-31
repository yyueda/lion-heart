package repository

import "github.com/yyueda/lion-heart/backend/campaign-service/internal/models"

type CampaignRepository interface {
	Create(c *models.Campaign) error
	GetByID(id string) (*models.Campaign, error)
	List() ([]*models.Campaign, error)
	Update(c *models.Campaign) error
	Delete(id string) error
}
