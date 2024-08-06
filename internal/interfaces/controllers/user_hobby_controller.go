package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/common"
	"github.com/coworker-match-api/internal/usecases"
	"github.com/gorilla/mux"
)

type UserHobbyController struct {
	userHobbyUsecase usecases.UserHobbyUsecase
}

func NewUserHobbyController(uh usecases.UserHobbyUsecase) *UserHobbyController {
	return &UserHobbyController{userHobbyUsecase: uh}
}

func (c *UserHobbyController) CreateUserHobby(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserHobbyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	key := common.UserIdKey
	userId, ok := r.Context().Value(key).(string)
	if !ok {
		http.Error(w, "Failed to get userId from context", http.StatusInternalServerError)
		return
	}

	response, err := c.userHobbyUsecase.CreateUserHobby(req, userId)
	if err != nil {
		http.Error(w, "Failed to get all hobbies", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to get all hobbies", http.StatusInternalServerError)
	}
}

func (c *UserHobbyController) GetAllUserHobby(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	allHobby, err := c.userHobbyUsecase.GetAllUserHobby(userId)
	if err != nil {
		http.Error(w, "Failed to get all hobbies", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(allHobby); err != nil {
		http.Error(w, "Failed to get all hobbies", http.StatusInternalServerError)
	}
}

func (c *UserHobbyController) UpdateUserHobby(w http.ResponseWriter, r *http.Request) {
	var req models.UpdateUserHobbyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	key := common.UserIdKey
	userId, ok := r.Context().Value(key).(string)
	if !ok {
		http.Error(w, "Failed to get userId from context", http.StatusInternalServerError)
		return
	}

	allHobby, err := c.userHobbyUsecase.UpdateUserHobby(req, userId)
	if err != nil {
		http.Error(w, "Failed to get all hobbies", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(allHobby); err != nil {
		http.Error(w, "Failed to get all hobbies", http.StatusInternalServerError)
	}
}
