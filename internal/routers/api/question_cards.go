package api

import (
	"database/sql"
	"fmt"
	"net/http"

	models "github.com/coworker-match-api/gen/go"
)

func (h *Handler) QuestionCardHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetQuestionCards(w, r, h.DB)
	default:
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetQuestionCards(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	matchingID := r.URL.Query().Get("matchingId")
	if matchingID == "" {
		writeError(w, "MatchingID is required", http.StatusBadRequest)
		return
	}

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

	rows, err := db.Query(query, matchingID)
	if err != nil {
		writeError(w, fmt.Sprintf("Error querying question cards: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var response []models.GetQuestionCardResponseInner

	for rows.Next() {
		var data models.GetQuestionCardResponseInner
		if err := rows.Scan(&data.QuestionCardId, &data.QuestionCardText, &data.IsUsed); err != nil {
			writeError(w, fmt.Sprintf("Error scanning row: %v", err), http.StatusInternalServerError)
			return
		}
		response = append(response, data)
	}

	respondWithJSON(w, response)
}
