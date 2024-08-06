package usecases

import (
	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/repositories"
)

type IHobbyUsecase interface {
	//TODO: CreateHobby() ()
	GetAllHobby() ([]models.GetHobbyResponseInner, error)
}

type hobbyUsecase struct {
	hr repositories.IHobbyRepo
}

func NewHobbyUsecase(hr repositories.IHobbyRepo) IHobbyUsecase {
	return &hobbyUsecase{hr: hr}
}

func (h *hobbyUsecase) GetAllHobby() ([]models.GetHobbyResponseInner, error) {
	return h.hr.GetAllHobby()
}
