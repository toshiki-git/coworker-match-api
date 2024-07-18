package api

import (
	"database/sql"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/coworker-match-api/internal/models"
)

type MatchingQuestion struct {
	QuestionID   string `json:"question_id"`
	QuestionText string `json:"question_text"`
	Choice1      Choice `json:"choice1"`
	Choice2      Choice `json:"choice2"`
}

type Choice struct {
	ChoiceText     string `json:"choice_text"`
	ChoiceImageURL string `json:"choice_image_url"`
}

type AnswerMatchingQuestionsRequest struct {
	Answers []Answer `json:"answers"`
}

type Answer struct {
	QuestionID string `json:"question_id"`
	Answer     string `json:"answer"`
}

type Match struct {
	MatchingID     string    `json:"matching_id"`
	SenderUserID   string    `json:"sender_user_id"`
	ReceiverUserID string    `json:"receiver_user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

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

	response := []MatchingQuestion{}

	for rows.Next() {
		var questionID, questionText string
		if err := rows.Scan(&questionID, &questionText); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response = append(response, MatchingQuestion{
			QuestionID:   questionID,
			QuestionText: questionText,
			Choice1: Choice{
				ChoiceText:     "YES",
				ChoiceImageURL: choice1Images[len(response)%len(choice1Images)],
			},
			Choice2: Choice{
				ChoiceText:     "NO",
				ChoiceImageURL: choice2Images[len(response)%len(choice2Images)],
			}})
	}
	respondWithJSON(w, response)
}

func handlePostMatchingQuestions(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var req AnswerMatchingQuestionsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	userID, ok := GetUserID(r.Context())
	if !ok {
		http.Error(w, "UserID not found", http.StatusUnauthorized)
		return
	}

	query := `SELECT * FROM users
			  WHERE user_id != $1
			  AND user_id NOT IN (
				SELECT receiver_user_id FROM matches WHERE sender_user_id = $1
				UNION
				SELECT sender_user_id FROM matches WHERE receiver_user_id = $1
			  );`

	rows, err := db.Query(query, userID)
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UserID, &user.UserName, &user.Email, &user.AvatarURL, &user.Age,
			&user.Gender, &user.Birthplace, &user.JobType, &user.LINEAccount,
			&user.DiscordAccount, &user.Biography, &user.CreatedAt, &user.UpdatedAt); err != nil { // 必要に応じて他のフィールドもスキャン
			http.Error(w, "Failed to scan user", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		http.Error(w, "No other users found", http.StatusNotFound)
		return
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rnd.Intn(len(users))
	receiverUser := users[randomIndex]

	var matchID string
	err = db.QueryRow("INSERT INTO matches (sender_user_id, receiver_user_id, created_at) VALUES ($1, $2, $3) RETURNING matching_id",
		userID, receiverUser.UserID, time.Now()).Scan(&matchID)
	if err != nil {
		http.Error(w, "Failed to create match", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"matching_id":      matchID,
		"sender_user_id":   userID,
		"receiver_user_id": receiverUser.UserID,
		"match_date":       time.Now(),
	}

	respondWithJSON(w, response)
}
