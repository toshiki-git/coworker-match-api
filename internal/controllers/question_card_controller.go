package controllers

import (
	"net/http"

	"github.com/coworker-match-api/internal/usecases"
)

type IQuestionCardController interface {
	GetQuestionCards(w http.ResponseWriter, r *http.Request)
}

type questionCardController struct {
	qcu usecases.IQuestionCardUsecase
}

func NewQuestionCardController(qcu usecases.IQuestionCardUsecase) IQuestionCardController {
	return &questionCardController{qcu: qcu}
}

func (qcc *questionCardController) GetQuestionCards(w http.ResponseWriter, r *http.Request) {

}
