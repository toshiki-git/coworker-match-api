package usecases

import (
	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/interfaces/repositories"
)

type HobbyUsecase interface {
	//CreateHobby() ()
	GetAllHobby() ([]models.GetHobbyResponseInner, error)
}

type hobbyUsecase struct {
	hobbyRepo repositories.HobbyRepository
}

func NewHobbyUsecase(repo repositories.HobbyRepository) HobbyUsecase {
	return &hobbyUsecase{hobbyRepo: repo}
}

func (h *hobbyUsecase) GetAllHobby() ([]models.GetHobbyResponseInner, error) {
	return h.hobbyRepo.GetAllHobby()
}
