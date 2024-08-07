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
	return u.ur.CreateUser(userId, req)
}

func (u *userUsecase) GetUserById(userId string) (*models.GetUserRes, error) {
	return u.ur.GetUserById(userId)
}

func (u *userUsecase) UpdateUser(userId string, req *models.UpdateUserReq) (*models.UpdateUserRes, error) {
	return u.ur.UpdateUser(userId, req)

}

func (u *userUsecase) IsUserExist(userId string) (bool, error) {
	return u.ur.IsUserExist(userId)
}
