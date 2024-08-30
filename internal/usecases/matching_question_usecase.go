package usecases

import (
	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/repositories"
)

type IMatchingQuestionUsecase interface {
	CreateMatching(userId string, req models.CreateQuestionReq) (*models.CreateQuestionRes, error)
	GetMatchingQuestion() (*models.GetQuestionRes, error)
}

type matchingQuestionUsecase struct {
	mqr repositories.IMatchingQuestionRepo
}

func NewMatchingQuestionUsecase(mqr repositories.IMatchingQuestionRepo) IMatchingQuestionUsecase {
	return &matchingQuestionUsecase{mqr: mqr}
}

func (mqu *matchingQuestionUsecase) CreateMatching(userId string, req models.CreateQuestionReq) (*models.CreateQuestionRes, error) {
	receiverUserId, err := mqu.mqr.GetMatchingCandidate(userId)
	if err != nil {
		return nil, err
	}

	response, err := mqu.mqr.InsertMatching(userId, receiverUserId)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (mqu *matchingQuestionUsecase) GetMatchingQuestion() (*models.GetQuestionRes, error) {
	return mqu.mqr.GetMatchingQuestion()
}
