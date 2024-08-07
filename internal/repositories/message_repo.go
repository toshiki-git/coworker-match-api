package repositories

import (
	"database/sql"

	models "github.com/coworker-match-api/gen/go"
)

type IMessageRepo interface {
	GetMessages(userId, matchingId string) (*models.GetMessageRes, error)
	CreateMessage(userId, matchingId string, req *models.CreateMessageReq) (*models.CreateMessageRes, error)
	UpdateMessage(messageId string, req models.UpdateMessageReq) (*models.UpdateMessageRes, error)
}

type messageRepo struct {
	db *sql.DB
}

func NewMessageRepo(db *sql.DB) IMessageRepo {
	return &messageRepo{db: db}
}

func (mr *messageRepo) GetMessages(userId, matchingId string) (*models.GetMessageRes, error) {
	return nil, nil
}

func (mr *messageRepo) CreateMessage(userId, matchingId string, req *models.CreateMessageReq) (*models.CreateMessageRes, error) {
	return nil, nil
}

func (mr *messageRepo) UpdateMessage(messageId string, req models.UpdateMessageReq) (*models.UpdateMessageRes, error) {
	return nil, nil
}
