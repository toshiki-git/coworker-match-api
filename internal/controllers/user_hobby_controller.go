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
	uhu usecases.IUserHobbyUsecase
}

func NewUserHobbyController(uhu usecases.IUserHobbyUsecase) *UserHobbyController {
	return &UserHobbyController{uhu: uhu}
}

func (uhc *UserHobbyController) CreateUserHobby(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserHobbyReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding request body: %v", err))
		return
	}

	key := common.UserIdKey
	userId, ok := r.Context().Value(key).(string)
	if !ok {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to get userId from context")
		return
	}

	response, err := uhc.uhu.CreateUserHobby(&req, userId)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to create user hobby")
		return
	}

	common.RespondWithJSON(w, http.StatusOK, response)
}

func (uhc *UserHobbyController) GetAllUserHobby(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	allHobby, err := uhc.uhu.GetAllUserHobby(userId)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.RespondWithJSON(w, http.StatusOK, allHobby)
}

func (uhc *UserHobbyController) UpdateUserHobby(w http.ResponseWriter, r *http.Request) {
	var req models.UpdateUserHobbyReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding request body: %v", err))
		return
	}

	key := common.UserIdKey
	userId, ok := r.Context().Value(key).(string)
	if !ok {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to get userId from context")
		return
	}

	response, err := uhc.uhu.UpdateUserHobby(&req, userId)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to update user hobby")
		return
	}

	common.RespondWithJSON(w, http.StatusOK, response)
}
