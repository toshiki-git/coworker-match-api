package api

import (
	"database/sql"
	"net/http"
)

func (h *Handler) MatchesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetMatches(w, r, h.DB)
	default:
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetMatches(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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
				END as image_url,
				CASE
					WHEN m.sender_user_id = $1 THEN ru.user_name
					ELSE su.user_name
				END as user_name,
				COALESCE(lm.message_text, 'メッセージはありません') as message,
				COALESCE(uc.unread_count, 0) as unread_count
			FROM
				matches m
				LEFT JOIN users su ON m.sender_user_id = su.user_id
				LEFT JOIN users ru ON m.receiver_user_id = ru.user_id
				LEFT JOIN (
					SELECT
						matching_id,
						message_text,
						created_at
					FROM
						messages
					WHERE
						(matching_id, created_at) IN (
							SELECT
								matching_id,
								MAX(created_at)
							FROM
								messages
							GROUP BY
								matching_id
						)
				) lm ON m.matching_id = lm.matching_id
				LEFT JOIN (
					SELECT
						matching_id,
						COUNT(*) as unread_count
					FROM
						messages
					WHERE
						is_read = false
						AND user_id != $1
					GROUP BY
						matching_id
				) uc ON m.matching_id = uc.matching_id
			WHERE
				m.sender_user_id = $1
				OR m.receiver_user_id = $1`

	rows, err := db.Query(query, userID)
	if err != nil {
		http.Error(w, "Failed to get matches", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		var matchID, imageURL, userName, message string
		var unreadCount int
		if err := rows.Scan(&matchID, &imageURL, &userName, &message, &unreadCount); err != nil {
			http.Error(w, "Failed to scan matches", http.StatusInternalServerError)
			return
		}
		results = append(results, map[string]interface{}{
			"matching_id":  matchID,
			"avatar_url":   imageURL,
			"name":         userName,
			"last_message": message,
			"unread_count": unreadCount,
		})
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error reading rows", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, results)
}
