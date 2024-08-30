package usecases

import (
	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/repositories"
)

type IUserUsecase interface {
	CreateUser(userId string, req *models.CreateUserReq) (*models.CreateUserRes, error)
	GetUserById(userId string) (*models.GetUserRes, error)
	UpdateUser(userId string, req *models.UpdateUserReq) (*models.UpdateUserRes, error)
	IsUserExist(userId string) (bool, error)
}

type userUsecase struct {
	ur repositories.IUserRepo
}

func NewUserUsecase(ur repositories.IUserRepo) IUserUsecase {
	return &userUsecase{ur: ur}
}

func (u *userUsecase) CreateUser(userId string, req *models.CreateUserReq) (*models.CreateUserRes, error) {
	user, err := u.ur.CreateUser(userId, req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) GetUserById(userId string) (*models.GetUserRes, error) {
	user, err := u.ur.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) UpdateUser(userId string, req *models.UpdateUserReq) (*models.UpdateUserRes, error) {
	user, err := u.ur.UpdateUser(userId, req)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (u *userUsecase) IsUserExist(userId string) (bool, error) {
	isUserExist, err := u.ur.IsUserExist(userId)
	if err != nil {
		return false, err
	}
	return isUserExist, nil
}
