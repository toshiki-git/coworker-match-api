package repositories

import (
	models "github.com/coworker-match-api/gen/go"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetUserById(userId string) (models.User, error)
	UpdateUser(userId string, updates map[string]interface{}) error
	IsUserExist(userId string) (bool, error)
}
