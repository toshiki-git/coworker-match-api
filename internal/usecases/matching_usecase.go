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
	ur repositories.IUserRepo
}

func NewMatchingUsecase(mr repositories.IMatchingRepo, ur repositories.IUserRepo) IMatchingUsecase {
	return &matchingUsecase{mr: mr, ur: ur}
}

func (mu *matchingUsecase) GetMatchings(userId string) (*models.GetMatchingRes, error) {
	matchings, err := mu.mr.GetMatchings(userId)
	if err != nil {
		return nil, err
	}
	return matchings, nil
}

func (mu *matchingUsecase) GetMatchingUser(userId, matchingId string) (*models.GetMatchingUserRes, error) {
	// matchingIdに基づいて相手のユーザーIDを取得
	matchingUserId, err := mu.mr.GetMatchingUserId(userId, matchingId)
	if err != nil {
		return nil, err
	}

	// 取得したユーザーIDでユーザー情報を取得
	user, err := mu.ur.GetUserById(matchingUserId)
	if err != nil {
		return nil, err
	}

	// *models.GetMatchingUserRes 型にマッピング
	matchingUser := &models.GetMatchingUserRes{
		UserId:    user.UserId,
		UserName:  user.UserName,
		Email:     user.Email,
		AvatarUrl: user.AvatarUrl,
	}

	return matchingUser, nil
}
