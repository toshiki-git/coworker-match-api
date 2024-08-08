package controllers

import (
	"net/http"

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
}

func (mqc *MatchingQuestionController) GetMatchingQuestion(w http.ResponseWriter, r *http.Request) {
	response, err := mqc.mqu.GetMatchingQuestion()
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to get matching question")
		return
	}
	common.RespondWithJSON(w, http.StatusOK, response)
}
