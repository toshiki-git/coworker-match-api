package controllers

import (
	"net/http"

	"github.com/coworker-match-api/internal/usecases"
)

type IMessageController interface {
	GetMessages(w http.ResponseWriter, r *http.Request)
	CreateMessage(w http.ResponseWriter, r *http.Request)
	UpdateMessage(w http.ResponseWriter, r *http.Request)
}

type messageController struct {
	mu usecases.IMessageUsecase
}

func NewMessageController(mu usecases.IMessageUsecase) IMessageController {
	return &messageController{mu: mu}
}

func (mc *messageController) GetMessages(w http.ResponseWriter, r *http.Request) {
}

func (mc *messageController) CreateMessage(w http.ResponseWriter, r *http.Request) {
}

func (mc *messageController) UpdateMessage(w http.ResponseWriter, r *http.Request) {
}
