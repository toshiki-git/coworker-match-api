package repositories

import (
	"database/sql"

	models "github.com/coworker-match-api/gen/go"
)

type IMatchingQuestionRepo interface {
	CreateMatching(userId string, req models.CreateQuestionReq) (*models.CreateQuestionRes, error)
	GetMatchingQuestion() (*models.GetQuestionRes, error)
}

type matchingQuestionRepo struct {
	db *sql.DB
}

func NewMatchingQuestionRepo(db *sql.DB) IMatchingQuestionRepo {
	return &matchingQuestionRepo{db: db}
}

func (mqr *matchingQuestionRepo) CreateMatching(userId string, req models.CreateQuestionReq) (*models.CreateQuestionRes, error) {
	return nil, nil
}

func (mqr *matchingQuestionRepo) GetMatchingQuestion() (*models.GetQuestionRes, error) {
	return nil, nil
}
