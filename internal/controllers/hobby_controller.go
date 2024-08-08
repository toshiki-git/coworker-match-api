package controllers

import (
	"net/http"

	"github.com/coworker-match-api/internal/common"
	"github.com/coworker-match-api/internal/usecases"
)

type HobbyController struct {
	hu usecases.IHobbyUsecase
}

func NewHobbyController(hu usecases.IHobbyUsecase) *HobbyController {
	return &HobbyController{hu: hu}
}

func (hc *HobbyController) GetAllHobby(w http.ResponseWriter, r *http.Request) {
	allHobby, err := hc.hu.GetAllHobby()
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to get all hobbies")
		return
	}

	common.RespondWithJSON(w, http.StatusOK, allHobby)
}
