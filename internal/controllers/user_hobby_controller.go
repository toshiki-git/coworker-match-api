package controllers

import (
	"encoding/json"
	"net/http"

	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/common"
	"github.com/coworker-match-api/internal/usecases"
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
		common.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := common.ExtractUserIdFromContext(r)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response, err := uhc.uhu.CreateUserHobby(&req, userId)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.RespondWithJSON(w, http.StatusOK, response)
}

func (uhc *UserHobbyController) GetAllUserHobby(w http.ResponseWriter, r *http.Request) {
	userId, err := common.ExtractPathParam(r, "userId")
	if err != nil {
		common.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

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
		common.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := common.ExtractUserIdFromContext(r)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response, err := uhc.uhu.UpdateUserHobby(&req, userId)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.RespondWithJSON(w, http.StatusOK, response)
}

func (uhc *UserHobbyController) GetUserCategoryPercentages(w http.ResponseWriter, r *http.Request) {
	userId, err := common.ExtractPathParam(r, "userId")
	if err != nil {
		common.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response, err := uhc.uhu.GetUserCategoryPercentages(userId)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.RespondWithJSON(w, http.StatusOK, response)
}
