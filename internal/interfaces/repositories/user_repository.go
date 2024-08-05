package repositories

import (
	models "github.com/coworker-match-api/gen/go"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetUserByID(userId string) (models.User, error)
	UpdateUser(userID string, updates map[string]interface{}) error
}
