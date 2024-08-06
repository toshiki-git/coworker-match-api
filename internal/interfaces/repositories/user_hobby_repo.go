package repositories

import (
	models "github.com/coworker-match-api/gen/go"
)

type UserHobbyRepo interface {
	CreateUserHobby(models.CreateUserHobbyRequest, string) (models.CreateUserHobbyResponse, error)
	GetAllUserHobby(string) ([]models.Hobby, error)
	UpdateUserHobby(models.UpdateUserHobbyRequest, string) (models.UpdateUserHobbyResponse, error)
}
