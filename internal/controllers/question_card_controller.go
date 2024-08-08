package controllers

import (
	"net/http"

	"github.com/coworker-match-api/internal/usecases"
)

type QuestionCardController struct {
	qcu usecases.IQuestionCardUsecase
}

func NewQuestionCardController(qcu usecases.IQuestionCardUsecase) *QuestionCardController {
	return &QuestionCardController{qcu: qcu}
}

func (qcc *QuestionCardController) GetQuestionCards(w http.ResponseWriter, r *http.Request) {

}
