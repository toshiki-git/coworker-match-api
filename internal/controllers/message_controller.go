package controllers

import (
	"encoding/json"
	"net/http"

	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/common"
	"github.com/coworker-match-api/internal/usecases"
	"github.com/gorilla/mux"
)

type MessageController struct {
	mu usecases.IMessageUsecase
}

func NewMessageController(mu usecases.IMessageUsecase) *MessageController {
	return &MessageController{mu: mu}
}

func (mc *MessageController) GetMessages(w http.ResponseWriter, r *http.Request) {
	key := common.UserIdKey
	userId, ok := r.Context().Value(key).(string)
	if !ok {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to get userId from context")
		return
	}

	vars := mux.Vars(r)
	matchingId := vars["matchingId"]
	if matchingId == "" {
		common.RespondWithError(w, http.StatusBadRequest, "Missing matchingId")
		return
	}

	response, err := mc.mu.GetMessages(userId, matchingId)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to get messages")
		return
	}

	common.RespondWithJSON(w, http.StatusOK, response)
}

func (mc *MessageController) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var req models.CreateMessageReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	key := common.UserIdKey
	userId, ok := r.Context().Value(key).(string)
	if !ok {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to get userId from context")
		return
	}

	vars := mux.Vars(r)
	matchingId := vars["matchingId"]
	if matchingId == "" {
		common.RespondWithError(w, http.StatusBadRequest, "Missing matchingId")
		return
	}

	response, err := mc.mu.CreateMessage(userId, matchingId, &req)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to create message")
		return
	}

	common.RespondWithJSON(w, http.StatusOK, response)
}

func (mc *MessageController) UpdateMessage(w http.ResponseWriter, r *http.Request) {
	var req models.UpdateMessageReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	vars := mux.Vars(r)
	messageId := vars["messageId"]
	if messageId == "" {
		common.RespondWithError(w, http.StatusBadRequest, "Missing messageId")
		return
	}

	response, err := mc.mu.UpdateMessage(messageId, req)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to update message")
		return
	}

	common.RespondWithJSON(w, http.StatusOK, response)
}
