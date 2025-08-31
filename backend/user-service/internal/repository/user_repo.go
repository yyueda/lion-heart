package repository

import (
	"github.com/yyueda/lion-heart/backend/user-service/internal/models"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id string) (*models.User, error)
}
