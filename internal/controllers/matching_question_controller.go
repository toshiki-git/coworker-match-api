package controllers

import (
	"net/http"

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

}
