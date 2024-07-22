package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/coworker-match-api/internal/models"
)

type CreateUserRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

type UserResponse struct {
	User models.User `json:"user"`
}

func (h *Handler) UserHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 {
		writeError(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	userID := pathParts[3]

	switch r.Method {
	case http.MethodGet:
		handleGetUser(w, h.DB, userID)
	case http.MethodPut:
		handlePutUser(w, r, h.DB, userID)
	default:
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	handlePostUser(w, r, h.DB)
}

func handleGetUser(w http.ResponseWriter, db *sql.DB, userID string) {
	const getUserSQL = `SELECT * FROM users WHERE user_id = $1`
	var user models.User
	if err := db.QueryRow(getUserSQL, userID).Scan(
		&user.UserID, &user.UserName, &user.Email, &user.AvatarURL, &user.Age,
		&user.Gender, &user.Birthplace, &user.JobType, &user.LINEAccount,
		&user.DiscordAccount, &user.Biography, &user.CreatedAt, &user.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			writeError(w, "User not found", http.StatusNotFound)
		} else {
			writeError(w, fmt.Sprintf("Error querying data: %v", err), http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, UserResponse{User: user})
}

func handlePostUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var req CreateUserRequest
	// コンテキストからユーザーIDを取得
	userID, ok := GetUserID(r.Context())
	if !ok {
		http.Error(w, "UserID not found", http.StatusUnauthorized)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	const insertUserSQL = `
		INSERT INTO users (user_id, user_name, email, avatar_url)
		VALUES ($1, $2, $3, $4)
		RETURNING user_id, user_name, email, avatar_url, age, gender,
		          birthplace, job_type, line_account, discord_account,
		          biography, created_at, updated_at
	`
	var user models.User
	if err := db.QueryRow(insertUserSQL, userID, req.Name, req.Email, req.AvatarURL).Scan(
		&user.UserID, &user.UserName, &user.Email, &user.AvatarURL, &user.Age,
		&user.Gender, &user.Birthplace, &user.JobType, &user.LINEAccount,
		&user.DiscordAccount, &user.Biography, &user.CreatedAt, &user.UpdatedAt,
	); err != nil {
		writeError(w, fmt.Sprintf("Error inserting data: %v", err), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, UserResponse{User: user})
}

func handlePutUser(w http.ResponseWriter, r *http.Request, db *sql.DB, userID string) {
	var userUpdates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&userUpdates); err != nil {
		writeError(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	setClause, args := buildUpdateClause(userUpdates)
	if len(setClause) == 0 {
		writeError(w, "No fields to update", http.StatusBadRequest)
		return
	}

	args = append(args, userID)
	const updateUserSQLTemplate = "UPDATE users SET %s, updated_at = NOW() WHERE user_id = $%d"
	updateUserSQL := fmt.Sprintf(updateUserSQLTemplate, strings.Join(setClause, ", "), len(args))

	if _, err := db.Exec(updateUserSQL, args...); err != nil {
		writeError(w, fmt.Sprintf("Error updating user: %v", err), http.StatusInternalServerError)
		return
	}

	handleGetUser(w, db, userID)
}

func buildUpdateClause(userUpdates map[string]interface{}) ([]string, []interface{}) {
	setClause := []string{}
	args := []interface{}{}
	argID := 1

	fields := map[string]string{
		"user_name":       "user_name",
		"email":           "email",
		"avatar_url":      "avatar_url",
		"age":             "age",
		"gender":          "gender",
		"birthplace":      "birthplace",
		"job_type":        "job_type",
		"line_account":    "line_account",
		"discord_account": "discord_account",
		"biography":       "biography",
	}

	for field, column := range fields {
		if value, ok := userUpdates[field]; ok {
			setClause = append(setClause, fmt.Sprintf("%s = $%d", column, argID))
			args = append(args, value)
			argID++
		}
	}

	return setClause, args
}

func (h *Handler) UserExistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// コンテキストからユーザーIDを取得
	userID, ok := GetUserID(r.Context())
	if !ok {
		http.Error(w, "UserID not found", http.StatusUnauthorized)
		return
	}

	const userExistSQL = `SELECT EXISTS(SELECT 1 FROM users WHERE user_id = $1)`
	var exists bool
	if err := h.DB.QueryRow(userExistSQL, userID).Scan(&exists); err != nil {
		respondWithJSON(w, map[string]bool{"exists": exists})
		return
	}

	respondWithJSON(w, map[string]bool{"exists": exists})
}
