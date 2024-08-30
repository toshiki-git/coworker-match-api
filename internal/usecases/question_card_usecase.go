package usecases

import (
	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/repositories"
)

type IQuestionCardUsecase interface {
	GetQuestionCards(matchingId string) (*models.GetQuestionCardRes, error)
}

type questionCardUsecase struct {
	qcr repositories.IQuestionCardRepo
}

func NewQuestionCardUsecase(qcr repositories.IQuestionCardRepo) IQuestionCardUsecase {
	return &questionCardUsecase{qcr: qcr}
}

func (qcu *questionCardUsecase) GetQuestionCards(matchingId string) (*models.GetQuestionCardRes, error) {
	questionCards, err := qcu.qcr.GetQuestionCards(matchingId)
	if err != nil {
		return nil, err
	}
	return questionCards, nil
}
