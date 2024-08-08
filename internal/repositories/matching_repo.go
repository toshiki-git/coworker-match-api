package repositories

import (
	"database/sql"

	models "github.com/coworker-match-api/gen/go"
)

type IMatchingRepo interface {
	GetMatchings(userId string) (*models.GetMatchingRes, error)
	GetMatchingUser(userId, matchingId string) (*models.GetMatchingUserRes, error)
}

type matchingRepo struct {
	db *sql.DB
}

func NewMatchingRepo(db *sql.DB) IMatchingRepo {
	return &matchingRepo{db: db}
}

func (mr *matchingRepo) GetMatchings(userId string) (*models.GetMatchingRes, error) {
	query := `
			SELECT
				m.matching_id,
				CASE
					WHEN m.sender_user_id = $1 THEN ru.avatar_url
					ELSE su.avatar_url
				END AS avatar_url,
				CASE
					WHEN m.sender_user_id = $1 THEN ru.user_name
					ELSE su.user_name
				END AS match_user_name,
				COALESCE(qc.question_card_text, 'メッセージはありません') AS last_message,
				0 AS unread_message_count
			FROM
				matchings m
				JOIN users su ON m.sender_user_id = su.user_id
				JOIN users ru ON m.receiver_user_id = ru.user_id
				LEFT JOIN (
					SELECT DISTINCT ON (matching_id)
						msg.matching_id,
						msg.question_card_id,
						msg.created_at
					FROM 
						messages msg
					JOIN (
						SELECT
							matching_id,
							MAX(created_at) AS max_created_at
						FROM
							messages
						GROUP BY
							matching_id
					) max_msgs ON msg.matching_id = max_msgs.matching_id AND msg.created_at = max_msgs.max_created_at
				) lm ON lm.matching_id = m.matching_id
				LEFT JOIN question_cards qc ON lm.question_card_id = qc.question_card_id
			WHERE 
				su.user_id = $1 OR ru.user_id = $1
			`
	rows, err := mr.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := &models.GetMatchingRes{
		Matchings: []models.GetMatchingResMatchingsInner{},
	}

	for rows.Next() {
		var data models.GetMatchingResMatchingsInner
		if err := rows.Scan(&data.MatchingId, &data.AvatarUrl, &data.MatchUserName, &data.LastMessage, &data.UnreadMessageCount); err != nil {
			return nil, err
		}
		response.Matchings = append(response.Matchings, data)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return response, nil
}

func (mr *matchingRepo) GetMatchingUser(userId, matchingId string) (*models.GetMatchingUserRes, error) {
	query := `
			SELECT
				CASE 
					WHEN su.user_id = $1 THEN ru.user_id
					ELSE su.user_id
				END AS user_id,
				CASE
					WHEN su.user_id = $1 THEN ru.user_name
					ELSE su.user_name
				END AS user_name,
				CASE
					WHEN su.user_id = $1 THEN ru.avatar_url
					ELSE su.avatar_url
				END AS avatar_url,
				CASE
					WHEN su.user_id = $1 THEN ru.email
					ELSE su.email
				END AS email
			FROM
				matchings m
				JOIN users su ON m.sender_user_id = su.user_id
				JOIN users ru ON m.receiver_user_id = ru.user_id
			WHERE
				m.matching_id = $2
		`
	var response models.GetMatchingUserRes

	if err := mr.db.QueryRow(query, userId, matchingId).Scan(&response.UserId, &response.UserName, &response.AvatarUrl, &response.Email); err != nil {
		return nil, err
	}

	return &response, nil
}
