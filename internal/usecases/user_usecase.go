package usecases

import (
	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/interfaces/repositories"
)

type UserUsecase interface {
	CreateUser(user models.User) (models.User, error)
	GetUserByID(userId string) (models.User, error)
	UpdateUser(userId string, updates map[string]interface{}) (models.User, error)
	IsUserExist(userId string) (bool, error)
}

type userUsecase struct {
	userRepo repositories.UserRepository
}

func NewUserUsecase(repo repositories.UserRepository) UserUsecase {
	return &userUsecase{userRepo: repo}
}

func (u *userUsecase) CreateUser(user models.User) (models.User, error) {
	return u.userRepo.CreateUser(user)
}

func (u *userUsecase) GetUserByID(userId string) (models.User, error) {
	return u.userRepo.GetUserByID(userId)
}

func (u *userUsecase) UpdateUser(userId string, updates map[string]interface{}) (models.User, error) {
	err := u.userRepo.UpdateUser(userId, updates)
	if err != nil {
		return models.User{}, err
	}
	return u.userRepo.GetUserByID(userId)
}

func (u *userUsecase) IsUserExist(userId string) (bool, error) {
	return u.userRepo.IsUserExist(userId)
}
