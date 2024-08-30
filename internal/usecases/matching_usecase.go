package usecases

import (
	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/repositories"
)

type IMatchingUsecase interface {
	GetMatchings(userId string) (*models.GetMatchingRes, error)
	GetMatchingUser(userId, matchingId string) (*models.GetMatchingUserRes, error)
}

type matchingUsecase struct {
	mr repositories.IMatchingRepo
}

func NewMatchingUsecase(mr repositories.IMatchingRepo) IMatchingUsecase {
	return &matchingUsecase{mr: mr}
}

func (m *matchingUsecase) GetMatchings(userId string) (*models.GetMatchingRes, error) {
	matchings, err := m.mr.GetMatchings(userId)
	if err != nil {
		return nil, err
	}
	return matchings, nil
}

func (m *matchingUsecase) GetMatchingUser(userId, matchingId string) (*models.GetMatchingUserRes, error) {
	matchingUser, err := m.mr.GetMatchingUser(userId, matchingId)
	if err != nil {
		return nil, err
	}
	return matchingUser, nil
}
