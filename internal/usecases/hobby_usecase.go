package usecases

import (
	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/repositories"
)

type IHobbyUsecase interface {
	//TODO: CreateHobby()
	GetAllHobby() (*models.GetHobbyRes, error)
}

type hobbyUsecase struct {
	hr repositories.IHobbyRepo
}

func NewHobbyUsecase(hr repositories.IHobbyRepo) IHobbyUsecase {
	return &hobbyUsecase{hr: hr}
}

func (h *hobbyUsecase) GetAllHobby() (*models.GetHobbyRes, error) {
	allHobby, err := h.hr.GetAllHobby()
	if err != nil {
		return nil, err
	}
	return allHobby, nil
}
