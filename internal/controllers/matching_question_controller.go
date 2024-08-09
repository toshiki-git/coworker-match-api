package controllers

import (
	"encoding/json"
	"net/http"

	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/common"
	"github.com/coworker-match-api/internal/usecases"
)

type MatchingQuestionController struct {
	mqu usecases.IMatchingQuestionUsecase
}

func NewMatchingQuestionController(mqu usecases.IMatchingQuestionUsecase) *MatchingQuestionController {
	return &MatchingQuestionController{mqu: mqu}
}

func (mqc *MatchingQuestionController) CreateMatching(w http.ResponseWriter, r *http.Request) {
	var req models.CreateQuestionReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := common.ExtractUserIdFromContext(r)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response, err := mqc.mqu.CreateMatching(userId, req)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.RespondWithJSON(w, http.StatusCreated, response)
}

func (mqc *MatchingQuestionController) GetMatchingQuestion(w http.ResponseWriter, r *http.Request) {
	response, err := mqc.mqu.GetMatchingQuestion()
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.RespondWithJSON(w, http.StatusOK, response)
}
