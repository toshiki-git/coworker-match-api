package usecases

import (
	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/interfaces/repositories"
)

type IUserHobbyUsecase interface {
	CreateUserHobby(models.CreateUserHobbyRequest, string) (models.CreateUserHobbyResponse, error)
	GetAllUserHobby(string) ([]models.Hobby, error)
	UpdateUserHobby(models.UpdateUserHobbyRequest, string) (models.UpdateUserHobbyResponse, error)
}

type userHobbyUsecase struct {
	uhr repositories.UserHobbyRepo
}

func NewUserHobbyUsecase(uhr repositories.UserHobbyRepo) IUserHobbyUsecase {
	return &userHobbyUsecase{uhr: uhr}
}

func (u *userHobbyUsecase) GetAllUserHobby(userId string) ([]models.Hobby, error) {
	return u.uhr.GetAllUserHobby(userId)
}

func (u *userHobbyUsecase) CreateUserHobby(req models.CreateUserHobbyRequest, userId string) (models.CreateUserHobbyResponse, error) {
	return u.uhr.CreateUserHobby(req, userId)
}

func (u *userHobbyUsecase) UpdateUserHobby(req models.UpdateUserHobbyRequest, userId string) (models.UpdateUserHobbyResponse, error) {
	return u.uhr.UpdateUserHobby(req, userId)
}
