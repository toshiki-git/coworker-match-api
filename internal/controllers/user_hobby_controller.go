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

type IUserHobbyController interface {
	CreateUserHobby(w http.ResponseWriter, r *http.Request)
	GetAllUserHobby(w http.ResponseWriter, r *http.Request)
	UpdateUserHobby(w http.ResponseWriter, r *http.Request)
}

type userHobbyController struct {
	uhu usecases.IUserHobbyUsecase
}

func NewUserHobbyController(uhu usecases.IUserHobbyUsecase) IUserHobbyController {
	return &userHobbyController{uhu: uhu}
}

func (uhc *userHobbyController) CreateUserHobby(w http.ResponseWriter, r *http.Request) {
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

	response, err := uhc.uhu.CreateUserHobby(&req, userId)
	if err != nil {
		http.Error(w, "Failed to create user hobby", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (uhc *userHobbyController) GetAllUserHobby(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	allHobby, err := uhc.uhu.GetAllUserHobby(userId)
	if err != nil {
		http.Error(w, "Failed to get all hobbies", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(allHobby); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (uhc *userHobbyController) UpdateUserHobby(w http.ResponseWriter, r *http.Request) {
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

	response, err := uhc.uhu.UpdateUserHobby(&req, userId)
	if err != nil {
		http.Error(w, "Failed to update user hobby", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
