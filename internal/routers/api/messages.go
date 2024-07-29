package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	models "github.com/coworker-match-api/gen/go"
)

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

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 {
		writeError(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	matchingID := pathParts[3]

	// 相手のユーザー情報とメッセージを取得するクエリ
	query := `
		SELECT 
			qc.question_card_id,
			qc.question_card_text,
			me_msg.message_id AS me_message_id,
			me_msg.message_text AS me_message_text,
			you_msg.message_id AS you_message_id,
			you_msg.message_text AS you_message_text
		FROM 
			matchings m
		JOIN (
			SELECT
				message_text,
				matching_id,
				question_card_id,
				message_id
			FROM
				messages
			WHERE
				user_id = $1
			) me_msg ON m.matching_id = me_msg.matching_id
		JOIN (
			SELECT
				message_text,
				matching_id,
				question_card_id,
				message_id
			FROM
				messages
			WHERE
				user_id != $1
			) you_msg ON m.matching_id = you_msg.matching_id
		JOIN question_cards qc ON me_msg.question_card_id = qc.question_card_id AND you_msg.question_card_id = qc.question_card_id
		WHERE 
			m.matching_id = $2
	`
	response := models.NewGetMessageResponse([]models.GetMessageResponseMessagesInner{})
	rows, err := db.Query(query, userID, matchingID)
	if err != nil {
		writeError(w, fmt.Sprintf("Error querying messages: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var data models.GetMessageResponseMessagesInner
		me := *models.NewMessage("", "")
		you := *models.NewMessage("", "")
		data.MessagePair = *models.NewGetMessageResponseMessagesInnerMessagePair(me, you)

		if err := rows.Scan(&data.QuestionCardId, &data.QuestionCardText, &data.MessagePair.Me.MessageId, &data.MessagePair.Me.MessageText, &data.MessagePair.You.MessageId, &data.MessagePair.You.MessageText); err != nil {
			writeError(w, fmt.Sprintf("Error scanning row: %v", err), http.StatusInternalServerError)
			return
		}
		response.Messages = append(response.Messages, data)
	}

	respondWithJSON(w, response)

	// メッセージをis_readに更新するクエリ
	updateQuery := `
		UPDATE messages 
		SET is_read = TRUE 
		WHERE matching_id = $1 
		AND user_id != $2 
		AND is_read = FALSE
	`

	if _, err := db.Exec(updateQuery, matchingID, userID); err != nil {
		writeError(w, fmt.Sprintf("Error updating messages: %v", err), http.StatusInternalServerError)
		return
	}
}

func handlePostMessage(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var req models.CreateMessageRequest

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 {
		writeError(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	matchingID := pathParts[3]

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
			matchings
		WHERE 
			matching_id = $2`

	if err := db.QueryRow(query, userID, matchingID).Scan(&otherUserID); err != nil {
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

	var response models.CreateMessageResponse

	// 自分のIDでメッセージを作成
	err = tx.QueryRow(`
		INSERT INTO messages (matching_id, question_card_id, user_id)
		VALUES ($1, $2, $3)
		RETURNING message_id`, matchingID, req.QuestionCardId, userID).Scan(&response.MessageId)
	if err != nil {
		writeError(w, fmt.Sprintf("Failed to insert message for user: %v", err), http.StatusInternalServerError)
		return
	}

	// 相手のIDでメッセージを作成
	_, err = tx.Exec(`
		INSERT INTO messages (matching_id, question_card_id, user_id)
		VALUES ($1, $2, $3)`, matchingID, req.QuestionCardId, otherUserID)
	if err != nil {
		writeError(w, fmt.Sprintf("Failed to insert message for other user: %v", err), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, response)
}

func handlePutMessage(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var req models.UpdateMessageRequest
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

	result, err := db.Exec(query, req.MessageText, messageID)
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

	response := &models.UpdateMessageResponse{
		MessageText: req.MessageText,
	}
	respondWithJSON(w, response)
}
