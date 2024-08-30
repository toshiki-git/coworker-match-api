package usecases

import (
	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/repositories"
)

type IMessageUsecase interface {
	GetAndMarkMessages(userId, matchingId string) (*models.GetMessageRes, error)
	CreateMessage(userId, matchingId string, req *models.CreateMessageReq) (*models.CreateMessageRes, error)
	UpdateMessage(messageId string, req models.UpdateMessageReq) (*models.UpdateMessageRes, error)
}

type messageUsecase struct {
	mr repositories.IMessageRepo
}

func NewMessageUsecase(mr repositories.IMessageRepo) IMessageUsecase {
	return &messageUsecase{mr: mr}
}

func (mu *messageUsecase) GetAndMarkMessages(userId, matchingId string) (*models.GetMessageRes, error) {
	// メッセージを取得
	messages, err := mu.mr.GetMessages(userId, matchingId)
	if err != nil {
		return nil, err
	}

	// メッセージのis_readを更新
	if err := mu.mr.UpdateMessagesReadStatus(matchingId, userId); err != nil {
		return nil, err
	}

	return messages, nil
}

func (mu *messageUsecase) CreateMessage(userId, matchingId string, req *models.CreateMessageReq) (*models.CreateMessageRes, error) {
	otherUserId, err := mu.mr.GetOtherUserId(userId, matchingId)
	if err != nil {
		return nil, err
	}

	messageId, err := mu.mr.CreateMessage(userId, otherUserId, matchingId, req)
	if err != nil {
		return nil, err
	}
	return messageId, nil
}

func (mu *messageUsecase) UpdateMessage(messageId string, req models.UpdateMessageReq) (*models.UpdateMessageRes, error) {
	return mu.mr.UpdateMessage(messageId, req)
}
