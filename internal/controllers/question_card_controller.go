package controllers

import (
	"net/http"

	"github.com/coworker-match-api/internal/common"
	"github.com/coworker-match-api/internal/usecases"
	"github.com/gorilla/mux"
)

type QuestionCardController struct {
	qcu usecases.IQuestionCardUsecase
}

func NewQuestionCardController(qcu usecases.IQuestionCardUsecase) *QuestionCardController {
	return &QuestionCardController{qcu: qcu}
}

func (qcc *QuestionCardController) GetQuestionCards(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	matchingId := vars["matchingId"]
	if matchingId == "" {
		common.RespondWithError(w, http.StatusBadRequest, "Missing matchingId")
		return
	}

	response, err := qcc.qcu.GetQuestionCards(matchingId)

	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.RespondWithJSON(w, http.StatusOK, response)

}
