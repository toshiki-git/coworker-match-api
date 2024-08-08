package repositories

import (
	"database/sql"

	models "github.com/coworker-match-api/gen/go"
)

var choice1Images = []string{
	"/yes1_image.png",
	"/yes2_image.png",
	"/yes3_image.png",
	"/yes4_image.png",
	"/yes5_image.png",
}

var choice2Images = []string{
	"/no1_image.png",
	"/no2_image.png",
	"/no3_image.png",
	"/no4_image.png",
	"/no5_image.png",
}

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
	query := `
			SELECT
				question_id,
				question_text
			FROM
				matching_questions
			LIMIT
				5;
			`
	rows, err := mqr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var response models.GetQuestionRes

	for rows.Next() {
		var data models.Question
		if err := rows.Scan(&data.QuestionId, &data.QuestionText); err != nil {
			return nil, err
		}

		/* data.Choice1 = *models.NewChoice("", "")
		data.Choice2 = *models.NewChoice("", "") */

		data.Choice1.SetChoiceText("YES")
		data.Choice1.SetChoiceImageUrl(choice1Images[len(response.Questions)%len(choice1Images)])
		data.Choice2.SetChoiceText("NO")
		data.Choice2.SetChoiceImageUrl(choice2Images[len(response.Questions)%len(choice2Images)])

		response.Questions = append(response.Questions, data)
	}
	return &response, nil
}
