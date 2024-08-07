package controllers

import (
	"net/http"

	"github.com/coworker-match-api/internal/common"
	"github.com/coworker-match-api/internal/usecases"
)

type IHobbyController interface {
	//TODO: CreateHobby(w http.ResponseWriter, r *http.Request)
	GetAllHobby(w http.ResponseWriter, r *http.Request)
}

type hobbyController struct {
	hu usecases.IHobbyUsecase
}

func NewHobbyController(hu usecases.IHobbyUsecase) IHobbyController {
	return &hobbyController{hu: hu}
}

func (hc *hobbyController) GetAllHobby(w http.ResponseWriter, r *http.Request) {
	allHobby, err := hc.hu.GetAllHobby()
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to get all hobbies")
		return
	}

	common.RespondWithJSON(w, http.StatusOK, allHobby)
}
