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
	return m.mr.GetMatchings(userId)
}

func (m *matchingUsecase) GetMatchingUser(userId, matchingId string) (*models.GetMatchingUserRes, error) {
	return m.mr.GetMatchingUser(userId, matchingId)
}
