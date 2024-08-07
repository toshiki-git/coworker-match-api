package controllers

import (
	"encoding/json"
	"net/http"

	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/common"
	"github.com/coworker-match-api/internal/usecases"
	"github.com/gorilla/mux"
)

type IUserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	IsUserExist(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	uu usecases.IUserUsecase
}

func NewUserController(uu usecases.IUserUsecase) IUserController {
	return &userController{uu: uu}
}

func (uc *userController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	key := common.UserIdKey
	userId, ok := r.Context().Value(key).(string)
	if !ok {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to get userId from context")
		return
	}

	createdUser, err := uc.uu.CreateUser(userId, &req)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}
	common.RespondWithJSON(w, http.StatusOK, createdUser)
}

func (uc *userController) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	if userId == "" {
		common.RespondWithError(w, http.StatusBadRequest, "Missing userId")
		return
	}

	user, err := uc.uu.GetUserById(userId)
	if err != nil {
		common.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	common.RespondWithJSON(w, http.StatusOK, user)
}

func (uc *userController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	if userId == "" {
		common.RespondWithError(w, http.StatusBadRequest, "Missing userId")
		return
	}

	var req models.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := uc.uu.UpdateUser(userId, &req)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to update user")
		return
	}
	common.RespondWithJSON(w, http.StatusOK, user)
}

func (uc *userController) IsUserExist(w http.ResponseWriter, r *http.Request) {
	key := common.UserIdKey
	userId, ok := r.Context().Value(key).(string)
	if !ok {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to get userId from context")
		return
	}

	isExist, err := uc.uu.IsUserExist(userId)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to check user existence")
		return
	}

	common.RespondWithJSON(w, http.StatusOK, map[string]bool{"exists": isExist})
}
