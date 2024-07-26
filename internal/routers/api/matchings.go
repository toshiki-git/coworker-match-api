package api

import (
	"database/sql"
	"net/http"

	models "github.com/coworker-match-api/gen/go"
)

func (h *Handler) MatchingHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetMatching(w, r, h.DB)
	default:
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetMatching(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// コンテキストからユーザーIDを取得するロジックを追加
	userID, ok := GetUserID(r.Context())
	if !ok {
		http.Error(w, "UserID not found", http.StatusUnauthorized)
		return
	}

	// マッチしたユーザーを取得するクエリ
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

	rows, err := db.Query(query, userID)
	if err != nil {
		http.Error(w, "Failed to get matches", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var response []models.GetMatchingResponseInner
	for rows.Next() {
		var data models.GetMatchingResponseInner
		if err := rows.Scan(&data.MatchingId, &data.AvatarUrl, &data.MatchUserName, &data.LastMessage, &data.UnreadMessageCount); err != nil {
			http.Error(w, "Failed to scan matches", http.StatusInternalServerError)
			return
		}
		response = append(response, data)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error reading rows", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, response)
}
