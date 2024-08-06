package controllers

import (
	"encoding/json"
	"net/http"

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
		http.Error(w, "Failed to get all hobbies", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(allHobby); err != nil {
		http.Error(w, "Failed to get all hobbies", http.StatusInternalServerError)
	}
}
