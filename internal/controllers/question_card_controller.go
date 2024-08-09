package controllers

import (
	"net/http"

	"github.com/coworker-match-api/internal/common"
	"github.com/coworker-match-api/internal/usecases"
)

type QuestionCardController struct {
	qcu usecases.IQuestionCardUsecase
}

func NewQuestionCardController(qcu usecases.IQuestionCardUsecase) *QuestionCardController {
	return &QuestionCardController{qcu: qcu}
}

func (qcc *QuestionCardController) GetQuestionCards(w http.ResponseWriter, r *http.Request) {
	matchingId, err := common.ExtractPathParam(r, "matchingId")
	if err != nil {
		common.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response, err := qcc.qcu.GetQuestionCards(matchingId)

	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.RespondWithJSON(w, http.StatusOK, response)

}
