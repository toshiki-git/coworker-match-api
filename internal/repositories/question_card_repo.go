package repositories

import (
	"database/sql"

	models "github.com/coworker-match-api/gen/go"
)

type IQuestionCardRepo interface {
	GetQuestionCards(matchingId string) (*models.GetQuestionCardRes, error)
}

type questionCardRepo struct {
	db *sql.DB
}

func NewQuestionCardRepo(db *sql.DB) IQuestionCardRepo {
	return &questionCardRepo{db: db}
}

func (qcr *questionCardRepo) GetQuestionCards(matchingId string) (*models.GetQuestionCardRes, error) {
	return nil, nil
}
