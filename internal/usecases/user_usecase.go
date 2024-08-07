package usecases

import (
	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/repositories"
)

type IUserUsecase interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserById(userId string) (*models.User, error)
	UpdateUser(userId string, updates map[string]interface{}) (*models.User, error)
	IsUserExist(userId string) (bool, error)
}

type userUsecase struct {
	ur repositories.IUserRepo
}

func NewUserUsecase(ur repositories.IUserRepo) IUserUsecase {
	return &userUsecase{ur: ur}
}

func (u *userUsecase) CreateUser(user *models.User) (*models.User, error) {
	return u.ur.CreateUser(user)
}

func (u *userUsecase) GetUserById(userId string) (*models.User, error) {
	return u.ur.GetUserById(userId)
}

func (u *userUsecase) UpdateUser(userId string, updates map[string]interface{}) (*models.User, error) {
	err := u.ur.UpdateUser(userId, updates)
	if err != nil {
		return nil, err
	}
	return u.ur.GetUserById(userId)
}

func (u *userUsecase) IsUserExist(userId string) (bool, error) {
	return u.ur.IsUserExist(userId)
}
