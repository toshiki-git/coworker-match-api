package repositories

import (
	"database/sql"

	models "github.com/coworker-match-api/gen/go"
)

type IQuestionCardRepo interface {
	GetQuestionCards(matchingId string) (*models.GetQuestionCardRes, error)
}

type questionCardRepo struct {
	db *sql.DB
}

func NewQuestionCardRepo(db *sql.DB) IQuestionCardRepo {
	return &questionCardRepo{db: db}
}

func (qcr *questionCardRepo) GetQuestionCards(matchingId string) (*models.GetQuestionCardRes, error) {
	query := `
		SELECT
			qc.question_card_id,
			qc.question_card_text,
			CASE WHEN m.question_card_id IS NOT NULL THEN true ELSE false END AS is_used
		FROM
			question_cards AS qc
		LEFT JOIN
			(
				SELECT DISTINCT question_card_id
				FROM messages
				WHERE matching_id = $1
			) AS m
		ON qc.question_card_id = m.question_card_id;
	`

	rows, err := qcr.db.Query(query, matchingId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := models.GetQuestionCardRes{
		QuestionCards: []models.GetQuestionCardResQuestionCardsInner{},
	}

	for rows.Next() {
		var data models.GetQuestionCardResQuestionCardsInner
		if err := rows.Scan(&data.QuestionCardId, &data.QuestionCardText, &data.IsUsed); err != nil {
			return nil, err
		}
		response.QuestionCards = append(response.QuestionCards, data)
	}

	return &response, nil

}
