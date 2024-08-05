package usecases

import (
	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/interfaces/repositories"
)

type UserUsecase interface {
	CreateUser(user models.User) (models.User, error)
	GetUserByID(userID string) (models.User, error)
	UpdateUser(userID string, updates map[string]interface{}) (models.User, error)
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

func (u *userUsecase) GetUserByID(userID string) (models.User, error) {
	return u.userRepo.GetUserByID(userID)
}

func (u *userUsecase) UpdateUser(userID string, updates map[string]interface{}) (models.User, error) {
	err := u.userRepo.UpdateUser(userID, updates)
	if err != nil {
		return models.User{}, err
	}
	return u.userRepo.GetUserByID(userID)
}
