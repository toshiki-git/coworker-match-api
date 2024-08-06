package usecases

import (
	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/interfaces/repositories"
)

type UserHobbyUsecase interface {
	CreateUserHobby(models.CreateUserHobbyRequest, string) (models.CreateUserHobbyResponse, error)
	GetAllUserHobby(string) ([]models.Hobby, error)
	UpdateUserHobby(models.UpdateUserHobbyRequest, string) (models.UpdateUserHobbyResponse, error)
}

type userHobbyUsecase struct {
	userHobbyRepo repositories.UserHobbyRepo
}

func NewUserHobbyUsecase(repo repositories.UserHobbyRepo) UserHobbyUsecase {
	return &userHobbyUsecase{userHobbyRepo: repo}
}

func (u *userHobbyUsecase) GetAllUserHobby(userId string) ([]models.Hobby, error) {
	return u.userHobbyRepo.GetAllUserHobby(userId)
}

func (u *userHobbyUsecase) CreateUserHobby(req models.CreateUserHobbyRequest, userId string) (models.CreateUserHobbyResponse, error) {
	return u.userHobbyRepo.CreateUserHobby(req, userId)
}

func (u *userHobbyUsecase) UpdateUserHobby(req models.UpdateUserHobbyRequest, userId string) (models.UpdateUserHobbyResponse, error) {
	return u.userHobbyRepo.UpdateUserHobby(req, userId)
}
