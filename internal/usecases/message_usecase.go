package usecases

import (
	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/repositories"
)

type IMessageUsecase interface {
	GetMessages(userId, matchingId string) (*models.GetMessageRes, error)
	CreateMessage(userId, matchingId string, req *models.CreateMessageReq) (*models.CreateMessageRes, error)
	UpdateMessage(messageId string, req models.UpdateMessageReq) (*models.UpdateMessageRes, error)
}

type messageUsecase struct {
	mr repositories.IMessageRepo
}

func NewMessageUsecase(mr repositories.IMessageRepo) IMessageUsecase {
	return &messageUsecase{mr: mr}
}

func (mu *messageUsecase) GetMessages(userId, matchingId string) (*models.GetMessageRes, error) {
	return mu.mr.GetMessages(userId, matchingId)
}

func (mu *messageUsecase) CreateMessage(userId, matchingId string, req *models.CreateMessageReq) (*models.CreateMessageRes, error) {
	return mu.mr.CreateMessage(userId, matchingId, req)
}

func (mu *messageUsecase) UpdateMessage(messageId string, req models.UpdateMessageReq) (*models.UpdateMessageRes, error) {
	return mu.mr.UpdateMessage(messageId, req)
}
