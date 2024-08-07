package repositories

import (
	"database/sql"

	models "github.com/coworker-match-api/gen/go"
)

type IMatchingRepo interface {
	GetMatchings(userId, matchingId string) (*models.GetMatchingRes, error)
	GetMatchingUser(userId, matchingId string) (*models.GetMatchingUserRes, error)
}

type matchingRepo struct {
	db *sql.DB
}

func NewMatchingRepo(db *sql.DB) IMatchingRepo {
	return &matchingRepo{db: db}
}

func (mr *matchingRepo) GetMatchings(userId, matchingId string) (*models.GetMatchingRes, error) {
	return nil, nil
}

func (mr *matchingRepo) GetMatchingUser(userId, matchingId string) (*models.GetMatchingUserRes, error) {
	return nil, nil
}
