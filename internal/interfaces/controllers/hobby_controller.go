package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/coworker-match-api/internal/usecases"
)

type HobbyController struct {
	HobbyUsecase usecases.HobbyUsecase
}

func NewHobbyController(u usecases.HobbyUsecase) *HobbyController {
	return &HobbyController{HobbyUsecase: u}
}

func (c *HobbyController) GetAllHobby(w http.ResponseWriter, r *http.Request) {
	allHobby, err := c.HobbyUsecase.GetAllHobby()
	if err != nil {
		http.Error(w, "Failed to get all hobbies", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(allHobby); err != nil {
		http.Error(w, "Failed to get all hobbies", http.StatusInternalServerError)
	}
}
