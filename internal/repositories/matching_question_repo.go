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
	GetMatchingCandidate(userId string) (string, error)
	InsertMatching(userId string, receiverUserId string) (*models.CreateQuestionRes, error)
	GetMatchingQuestion() (*models.GetQuestionRes, error)
}

type matchingQuestionRepo struct {
	db *sql.DB
}

func NewMatchingQuestionRepo(db *sql.DB) IMatchingQuestionRepo {
	return &matchingQuestionRepo{db: db}
}

func (mqr *matchingQuestionRepo) GetMatchingCandidate(userId string) (string, error) {
	query := `
			WITH excluded_users AS (
    			SELECT receiver_user_id AS user_id 
    			FROM matchings 
    			WHERE sender_user_id = $1
    			UNION
    			SELECT sender_user_id 
    			FROM matchings 
    			WHERE receiver_user_id = $1
			)
			SELECT 
				u.user_id
			FROM 
				users u
			WHERE 
				user_id != $1
				AND u.user_id NOT IN (SELECT user_id FROM excluded_users)
			ORDER BY
				RANDOM()
			LIMIT 1;
			`
	var receiverUserId string

	if err := mqr.db.QueryRow(query, userId).Scan(&receiverUserId); err != nil {
		return "", err
	}

	return receiverUserId, nil
}

// マッチング情報をデータベースに挿入する関数
func (mqr *matchingQuestionRepo) InsertMatching(userId string, receiverUserId string) (*models.CreateQuestionRes, error) {
	var response models.CreateQuestionRes

	err := mqr.db.QueryRow("INSERT INTO matchings (sender_user_id, receiver_user_id) VALUES ($1, $2) RETURNING matching_id, created_at",
		userId, receiverUserId).Scan(&response.MatchingId, &response.MatchingDate)
	if err != nil {
		return nil, err
	}

	response.SetSenderUserId(userId)
	response.ReceiverUserId = receiverUserId

	return &response, nil
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
