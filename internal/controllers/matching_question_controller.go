package controllers

import (
	"net/http"

	"github.com/coworker-match-api/internal/common"
	"github.com/coworker-match-api/internal/usecases"
)

type IMatchingQuestionController interface {
	CreateMatching(w http.ResponseWriter, r *http.Request)
	GetMatchingQuestion(w http.ResponseWriter, r *http.Request)
}

type matchingQuestionController struct {
	mqu usecases.IMatchingQuestionUsecase
}

func NewMatchingQuestionController(mqu usecases.IMatchingQuestionUsecase) IMatchingQuestionController {
	return &matchingQuestionController{mqu: mqu}
}

func (mqc *matchingQuestionController) CreateMatching(w http.ResponseWriter, r *http.Request) {
}

func (mqc *matchingQuestionController) GetMatchingQuestion(w http.ResponseWriter, r *http.Request) {
	response, err := mqc.mqu.GetMatchingQuestion()
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to get matching question")
		return
	}
	common.RespondWithJSON(w, http.StatusOK, response)
}
