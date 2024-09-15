package usecases

import (
	"errors"
	"sort"

	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/repositories"
)

type IUserHobbyUsecase interface {
	CreateUserHobby(*models.CreateUserHobbyReq, string) (*models.CreateUserHobbyRes, error)
	GetAllUserHobby(string) (*models.GetUserHobbyRes, error)
	UpdateUserHobby(*models.UpdateUserHobbyReq, string) (*models.UpdateUserHobbyRes, error)
	GetUserCategoryPercentages(userId string) (*models.GetUserCategoryPercentagesRes, error)
}

type userHobbyUsecase struct {
	uhr repositories.IUserHobbyRepo
}

func NewUserHobbyUsecase(uhr repositories.IUserHobbyRepo) IUserHobbyUsecase {
	return &userHobbyUsecase{uhr: uhr}
}

func (uhu *userHobbyUsecase) GetAllUserHobby(userId string) (*models.GetUserHobbyRes, error) {
	allUserHobby, err := uhu.uhr.GetAllUserHobby(userId)
	if err != nil {
		return nil, err
	}
	return allUserHobby, nil
}

func (uhu *userHobbyUsecase) CreateUserHobby(req *models.CreateUserHobbyReq, userId string) (*models.CreateUserHobbyRes, error) {
	hobbyIds, err := uhu.uhr.CreateUserHobby(req, userId)
	if err != nil {
		return nil, err
	}
	return hobbyIds, nil
}

func (uhu *userHobbyUsecase) UpdateUserHobby(req *models.UpdateUserHobbyReq, userId string) (*models.UpdateUserHobbyRes, error) {
	hobbyIds, err := uhu.uhr.UpdateUserHobby(req, userId)
	if err != nil {
		return nil, err
	}
	return hobbyIds, nil
}

func (uhu *userHobbyUsecase) GetUserCategoryPercentages(userId string) (*models.GetUserCategoryPercentagesRes, error) {
	userHobbiesByCategory, err := uhu.uhr.GetUserHobbiesByCategory(userId)
	if err != nil {
		return nil, err
	}

	totalHobbiesByCategory, err := uhu.uhr.GetTotalHobbiesByCategory()
	if err != nil {
		return nil, err
	}

	response := &models.GetUserCategoryPercentagesRes{
		UserId:     userId,
		Categories: []models.CategoryInterest{},
	}

	// categoryIdの順番を固定するために、まずキーをスライスに収集
	categoryIds := make([]string, 0, len(totalHobbiesByCategory))
	for categoryId := range totalHobbiesByCategory {
		categoryIds = append(categoryIds, categoryId)
	}

	// categoryIdをソート
	sort.Strings(categoryIds)

	// ソートされた順序で処理を実行
	for _, categoryId := range categoryIds {
		totalHobby, ok := totalHobbiesByCategory[categoryId]
		if !ok {
			return nil, errors.New("total hobby count not found")
		}

		// userHobbiesByCategory にそのカテゴリーが存在するか確認
		userHobby, exists := userHobbiesByCategory[categoryId]
		interestPercentage := 0 // ユーザーの趣味がない場合は0%
		if exists {
			interestPercentage = (userHobby.Count * 100) / totalHobby.Count
		}

		// カテゴリーをレスポンスに追加
		response.Categories = append(response.Categories, models.CategoryInterest{
			CategoryId:         categoryId,
			CategoryName:       totalHobby.Name, // 存在しない場合は空でいいなら、存在チェックを入れる
			InterestPercentage: int32(interestPercentage),
		})
	}

	return response, nil
}
