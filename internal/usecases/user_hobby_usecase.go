package usecases

import (
	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/repositories"
)

type IUserHobbyUsecase interface {
	CreateUserHobby(*models.CreateUserHobbyReq, string) (*models.CreateUserHobbyRes, error)
	GetAllUserHobby(string) (*models.GetUserHobbyRes, error)
	UpdateUserHobby(*models.UpdateUserHobbyReq, string) (*models.UpdateUserHobbyRes, error)
}

type userHobbyUsecase struct {
	uhr repositories.IUserHobbyRepo
}

func NewUserHobbyUsecase(uhr repositories.IUserHobbyRepo) IUserHobbyUsecase {
	return &userHobbyUsecase{uhr: uhr}
}

func (u *userHobbyUsecase) GetAllUserHobby(userId string) (*models.GetUserHobbyRes, error) {
	allUserHobby, err := u.uhr.GetAllUserHobby(userId)
	if err != nil {
		return nil, err
	}
	return allUserHobby, nil
}

func (u *userHobbyUsecase) CreateUserHobby(req *models.CreateUserHobbyReq, userId string) (*models.CreateUserHobbyRes, error) {
	hobbyIds, err := u.uhr.CreateUserHobby(req, userId)
	if err != nil {
		return nil, err
	}
	return hobbyIds, nil
}

func (u *userHobbyUsecase) UpdateUserHobby(req *models.UpdateUserHobbyReq, userId string) (*models.UpdateUserHobbyRes, error) {
	hobbyIds, err := u.uhr.UpdateUserHobby(req, userId)
	if err != nil {
		return nil, err
	}
	return hobbyIds, nil
}
