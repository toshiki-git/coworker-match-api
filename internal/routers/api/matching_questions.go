package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

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

func (h *Handler) MatchingQuestionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetMatchingQuestions(w, h.DB)
	case http.MethodPost:
		handlePostMatchingQuestions(w, r, h.DB)
	default:
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetMatchingQuestions(w http.ResponseWriter, db *sql.DB) {
	query := `
	SELECT
		question_id,
		question_text
	FROM
		matching_questions
	LIMIT
		5;
	`

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var response []models.Question

	for rows.Next() {
		var data models.Question
		if err := rows.Scan(&data.QuestionId, &data.QuestionText); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data.Choice1 = *models.NewChoice("", "")
		data.Choice2 = *models.NewChoice("", "")

		data.Choice1.SetChoiceText("YES")
		data.Choice1.SetChoiceImageUrl(choice1Images[len(response)%len(choice1Images)])
		data.Choice2.SetChoiceText("NO")
		data.Choice2.SetChoiceImageUrl(choice2Images[len(response)%len(choice2Images)])

		response = append(response, data)
	}
	respondWithJSON(w, response)
}

func handlePostMatchingQuestions(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var req models.CreateQuestionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	userID, ok := GetUserID(r.Context())
	if !ok {
		http.Error(w, "UserID not found", http.StatusUnauthorized)
		return
	}

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
			LIMIT 1
			`

	var response models.CreateQuestionResponse
	if err := db.QueryRow(query, userID).Scan(&response.ReceiverUserId); err != nil {
		writeError(w, "マッチできるユーザがいません", http.StatusInternalServerError)
		return
	}

	err := db.QueryRow("INSERT INTO matchings (sender_user_id, receiver_user_id) VALUES ($1, $2) RETURNING matching_id, created_at",
		userID, response.ReceiverUserId).Scan(&response.MatchingId, &response.MatchingDate)
	if err != nil {
		http.Error(w, "Failed to create match", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, response)
}
