package repositories

import (
	"database/sql"

	models "github.com/coworker-match-api/gen/go"
)

type IMatchingRepo interface {
	GetMatchings(userId string) (*models.GetMatchingRes, error)
	GetMatchingUserId(userId, matchingId string) (string, error)
	//GetMatchingUser(userId, matchingId string) (*models.GetMatchingUserRes, error)
}

type matchingRepo struct {
	db *sql.DB
}

func NewMatchingRepo(db *sql.DB) IMatchingRepo {
	return &matchingRepo{db: db}
}

func (mr *matchingRepo) GetMatchingUserId(userId, matchingId string) (string, error) {
	var otherUserId string
	query := `
		SELECT 
			CASE 
				WHEN sender_user_id = $1 THEN receiver_user_id
				ELSE sender_user_id
			END AS other_user_id
		FROM 
			matchings
		WHERE 
			matching_id = $2
	`
	err := mr.db.QueryRow(query, userId, matchingId).Scan(&otherUserId)
	if err != nil {
		return "", err
	}

	return otherUserId, nil
}

func (mr *matchingRepo) GetMatchings(userId string) (*models.GetMatchingRes, error) {
	query := `
		WITH matched_users AS (
			SELECT
				m.matching_id,
				CASE
					WHEN m.sender_user_id = $1 THEN m.receiver_user_id
					ELSE m.sender_user_id
				END AS match_user_id,
				m.created_at AS matching_created_at
			FROM
				matchings m
			WHERE
				m.sender_user_id = $1 OR m.receiver_user_id = $1
		),

		latest_messages AS (
			SELECT DISTINCT ON (matching_id)
				matching_id,
				question_card_id,
				updated_at
			FROM
				messages
			ORDER BY
				matching_id, updated_at DESC
		)

		SELECT
			mu.matching_id,
			u.avatar_url,
			u.user_name AS match_user_name,
			COALESCE(qc.question_card_text, 'メッセージはありません') AS last_message,
			0 AS unread_message_count
		FROM
			matched_users mu
		JOIN users u ON u.user_id = mu.match_user_id
		LEFT JOIN latest_messages lm ON lm.matching_id = mu.matching_id
		LEFT JOIN question_cards qc ON lm.question_card_id = qc.question_card_id
		ORDER BY
			COALESCE(lm.updated_at, mu.matching_created_at) DESC;
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

/* func (mr *matchingRepo) GetMatchingUser(userId, matchingId string) (*models.GetMatchingUserRes, error) {
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
*/
