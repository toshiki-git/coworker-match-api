package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type CreateMessageHandlerRequest struct {
	MatchingID     string `json:"matching_id"`
	QuestionCardID string `json:"question_card_id"`
}

type MessagePair struct {
	Me struct {
		MessageID   string    `json:"message_id"`
		MessageText string    `json:"message_text"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	} `json:"me"`
	You struct {
		MessageID   string    `json:"message_id"`
		MessageText string    `json:"message_text"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	} `json:"you"`
}

type Message struct {
	QuestionID   string      `json:"question_id"`
	QuestionText string      `json:"question_text"`
	CreatedAt    time.Time   `json:"created_at"`
	MessagePair  MessagePair `json:"message_pair"`
}

type MatchUser struct {
	UserID    string   `json:"user_id"`
	UserName  string   `json:"user_name"`
	AvatarURL string   `json:"avatar_url"`
	Hobbies   []string `json:"hobbies"`
}

type MainData struct {
	MatchUser MatchUser `json:"match_user"`
	Messages  []Message `json:"messages"`
}

type UpdateMessageHandlerRequest struct {
	Answer string `json:"answer"`
}

func (h *Handler) MessageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetMessages(w, r, h.DB)
	case http.MethodPost:
		handlePostMessage(w, r, h.DB)
	case http.MethodPut:
		handlePutMessage(w, r, h.DB)
	default:
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetMessages(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	userID, ok := GetUserID(r.Context())
	if !ok {
		http.Error(w, "UserID not found", http.StatusUnauthorized)
		return
	}

	// クエリパラメータからmatching_idを取得
	matchingID := r.URL.Query().Get("matching_id")
	if matchingID == "" {
		writeError(w, "MatchingID is required", http.StatusBadRequest)
		return
	}

	// 相手のユーザー情報とメッセージを取得するクエリ
	query := `
		SELECT 
			other_user.user_id, 
			other_user.user_name, 
			other_user.avatar_url, 
			array_agg(h.hobby_name) AS hobbies, 
			msg.message_id, 
			msg.message_text, 
			msg.created_at, 
			msg.updated_at, 
			q.question_card_id AS question_id, 
			q.question_card_text AS question_text, 
			q.created_at AS question_created_at,
			msg.user_id
		FROM 
			matches AS m
		JOIN 
			users AS other_user ON 
				(m.sender_user_id = other_user.user_id AND m.receiver_user_id = $1)
				OR 
				(m.receiver_user_id = other_user.user_id AND m.sender_user_id = $1)
		JOIN 
			messages AS msg ON m.matching_id = msg.matching_id
		JOIN 
			question_cards AS q ON msg.question_card_id = q.question_card_id
		LEFT JOIN 
			user_hobbies uh ON uh.user_id = other_user.user_id
		LEFT JOIN 
			hobbies h ON h.hobby_id = uh.hobby_id
		WHERE 
			m.matching_id = $2
		GROUP BY 
			other_user.user_id, msg.message_id, q.question_card_id, msg.user_id
		ORDER BY
			msg.created_at ASC
	`

	rows, err := db.Query(query, userID, matchingID)
	if err != nil {
		writeError(w, fmt.Sprintf("Error querying messages: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var mainData MainData
	messageMap := make(map[string]*Message) // question_idをキーにする

	for rows.Next() {
		var matchUser MatchUser
		var hobbies []byte
		var messageID, questionID, questionText, messageUserID string
		var messageText sql.NullString
		var createdAt, updatedAt, questionCreatedAt time.Time

		if err := rows.Scan(
			&matchUser.UserID, &matchUser.UserName, &matchUser.AvatarURL, &hobbies,
			&messageID, &messageText, &createdAt, &updatedAt,
			&questionID, &questionText, &questionCreatedAt, &messageUserID); err != nil {
			writeError(w, fmt.Sprintf("Error scanning row: %v", err), http.StatusInternalServerError)
			return
		}

		matchUser.Hobbies = []string{}
		json.Unmarshal(hobbies, &matchUser.Hobbies)
		mainData.MatchUser = matchUser

		if _, exists := messageMap[questionID]; !exists {
			messageMap[questionID] = &Message{
				QuestionID:   questionID,
				QuestionText: questionText,
				CreatedAt:    questionCreatedAt,
			}
		}

		if messageUserID == userID {
			messageMap[questionID].MessagePair.Me.MessageID = messageID
			messageMap[questionID].MessagePair.Me.MessageText = messageText.String
			messageMap[questionID].MessagePair.Me.CreatedAt = createdAt
			messageMap[questionID].MessagePair.Me.UpdatedAt = updatedAt
		} else {
			messageMap[questionID].MessagePair.You.MessageID = messageID
			messageMap[questionID].MessagePair.You.MessageText = messageText.String
			messageMap[questionID].MessagePair.You.CreatedAt = createdAt
			messageMap[questionID].MessagePair.You.UpdatedAt = updatedAt
		}
	}

	for _, msg := range messageMap {
		mainData.Messages = append(mainData.Messages, *msg)
	}

	respondWithJSON(w, mainData)
}

func handlePostMessage(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var req CreateMessageHandlerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	// コンテキストからユーザーIDを取得するロジックを追加
	userID, ok := GetUserID(r.Context())
	if !ok {
		http.Error(w, "UserID not found", http.StatusUnauthorized)
		return
	}

	var otherUserID string
	query := `
		SELECT 
			CASE 
				WHEN sender_user_id = $1 THEN receiver_user_id
				ELSE sender_user_id
			END AS other_user_id
		FROM 
			matches
		WHERE 
			matching_id = $2`

	if err := db.QueryRow(query, userID, req.MatchingID).Scan(&otherUserID); err != nil {
		writeError(w, fmt.Sprintf("Error querying other user ID: %v", err), http.StatusInternalServerError)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		writeError(w, fmt.Sprintf("Failed to begin transaction: %v", err), http.StatusInternalServerError)
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var messageID string

	// 自分のIDでメッセージを作成
	err = tx.QueryRow(`
		INSERT INTO messages (matching_id, question_card_id, user_id)
		VALUES ($1, $2, $3)
		RETURNING message_id`, req.MatchingID, req.QuestionCardID, userID).Scan(&messageID)
	if err != nil {
		writeError(w, fmt.Sprintf("Failed to insert message for user: %v", err), http.StatusInternalServerError)
		return
	}

	// 相手のIDでメッセージを作成
	_, err = tx.Exec(`
		INSERT INTO messages (matching_id, question_card_id, user_id)
		VALUES ($1, $2, $3)`, req.MatchingID, req.QuestionCardID, otherUserID)
	if err != nil {
		writeError(w, fmt.Sprintf("Failed to insert message for other user: %v", err), http.StatusInternalServerError)
		return
	}

	// 成功レスポンスを返す
	response := map[string]string{
		"message_id": messageID,
	}

	respondWithJSON(w, response)
}
func handlePutMessage(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var req UpdateMessageHandlerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	// URLからメッセージIDを取得
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 {
		writeError(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	messageID := pathParts[3]

	// メッセージを更新するクエリ
	query := `
		UPDATE messages 
		SET message_text = $1, updated_at = NOW() 
		WHERE message_id = $2
	`

	result, err := db.Exec(query, req.Answer, messageID)
	if err != nil {
		writeError(w, fmt.Sprintf("Failed to update message: %v", err), http.StatusInternalServerError)
		return
	}

	// 更新された行数を確認
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		writeError(w, fmt.Sprintf("Error checking rows affected: %v", err), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		writeError(w, "Message not found or not authorized to update", http.StatusNotFound)
		return
	}

	// 成功レスポンスを返す
	response := map[string]string{
		"message_id": messageID, // メッセージIDをレスポンスに含める
	}

	respondWithJSON(w, response)
}
