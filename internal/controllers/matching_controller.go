package controllers

import (
	"net/http"

	"github.com/coworker-match-api/internal/common"
	"github.com/coworker-match-api/internal/usecases"
)

type MatchingController struct {
	mu usecases.IMatchingUsecase
}

func NewMatchingController(mu usecases.IMatchingUsecase) *MatchingController {
	return &MatchingController{mu: mu}
}

func (mc *MatchingController) GetMatchings(w http.ResponseWriter, r *http.Request) {
	key := common.UserIdKey
	userId, ok := r.Context().Value(key).(string)
	if !ok {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to get userId from context")
		return
	}

	response, err := mc.mu.GetMatchings(userId)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.RespondWithJSON(w, http.StatusOK, response)
}

func (mc *MatchingController) GetMatchingUser(w http.ResponseWriter, r *http.Request) {
	matchingId, err := common.ExtractPathParam(r, "matchingId")
	if err != nil {
		common.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	key := common.UserIdKey
	userId, ok := r.Context().Value(key).(string)
	if !ok {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to get userId from context")
		return
	}

	response, err := mc.mu.GetMatchingUser(userId, matchingId)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.RespondWithJSON(w, http.StatusOK, response)

}
