package controllers

import (
	"encoding/json"
	"net/http"

	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/common"
	"github.com/coworker-match-api/internal/usecases"
)

type MessageController struct {
	mu usecases.IMessageUsecase
}

func NewMessageController(mu usecases.IMessageUsecase) *MessageController {
	return &MessageController{mu: mu}
}

func (mc *MessageController) GetAndMarkMessages(w http.ResponseWriter, r *http.Request) {
	userId, err := common.ExtractUserIdFromContext(r)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	matchingId, err := common.ExtractPathParam(r, "matchingId")
	if err != nil {
		common.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response, err := mc.mu.GetAndMarkMessages(userId, matchingId)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.RespondWithJSON(w, http.StatusOK, response)
}

func (mc *MessageController) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var req models.CreateMessageReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := common.ExtractUserIdFromContext(r)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	matchingId, err := common.ExtractPathParam(r, "matchingId")
	if err != nil {
		common.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response, err := mc.mu.CreateMessage(userId, matchingId, &req)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.RespondWithJSON(w, http.StatusOK, response)
}

func (mc *MessageController) UpdateMessage(w http.ResponseWriter, r *http.Request) {
	var req models.UpdateMessageReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	messageId, err := common.ExtractPathParam(r, "messageId")
	if err != nil {
		common.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response, err := mc.mu.UpdateMessage(messageId, req)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.RespondWithJSON(w, http.StatusOK, response)
}
