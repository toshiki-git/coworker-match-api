package controllers

import (
	"net/http"

	"github.com/coworker-match-api/internal/usecases"
)

type MessageController struct {
	mu usecases.IMessageUsecase
}

func NewMessageController(mu usecases.IMessageUsecase) *MessageController {
	return &MessageController{mu: mu}
}

func (mc *MessageController) GetMessages(w http.ResponseWriter, r *http.Request) {
}

func (mc *MessageController) CreateMessage(w http.ResponseWriter, r *http.Request) {
}

func (mc *MessageController) UpdateMessage(w http.ResponseWriter, r *http.Request) {
}
