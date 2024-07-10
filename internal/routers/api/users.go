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
	// パスからuser_idを抽出
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 4 {
		writeError(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	userID := parts[3]

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
	switch r.Method {
	case http.MethodPost:
		handlePostUser(w, r, h.DB)
	default:
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetUser(w http.ResponseWriter, db *sql.DB, userID string) {
	getUserSQL := `SELECT * FROM users WHERE user_id = $1`
	row := db.QueryRow(getUserSQL, userID)

	var user models.User
	err := row.Scan(&user.UserID, &user.UserName, &user.Email, &user.AvatarURL, &user.Age, &user.Gender, &user.Birthplace, &user.JobType, &user.LINEAccount, &user.DiscordAccount, &user.Biography, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			writeError(w, "User not found", http.StatusNotFound)
		} else {
			writeError(w, fmt.Sprintf("Error querying data: %v", err), http.StatusInternalServerError)
		}
		return
	}

	response := UserResponse{User: user}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		writeError(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
	}
}

func handlePostUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	insertUserSQL := `INSERT INTO users (user_name, email, avatar_url) VALUES ($1, $2, $3) RETURNING user_id, user_name, email, avatar_url, age, gender, birthplace, job_type, line_account, discord_account, biography, created_at, updated_at`
	var user models.User
	err := db.QueryRow(insertUserSQL, req.Name, req.Email, req.AvatarURL).Scan(&user.UserID, &user.UserName, &user.Email, &user.AvatarURL, &user.Age, &user.Gender, &user.Birthplace, &user.JobType, &user.LINEAccount, &user.DiscordAccount, &user.Biography, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		writeError(w, fmt.Sprintf("Error inserting data: %v", err), http.StatusInternalServerError)
		return
	}

	// 成功時のレスポンス
	response := UserResponse{User: user}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		writeError(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
	}
}

func handlePutUser(w http.ResponseWriter, r *http.Request, db *sql.DB, userID string) {
	var user map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		writeError(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	// 動的にSQLクエリを構築する
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
		if value, ok := user[field]; ok {
			setClause = append(setClause, fmt.Sprintf("%s = $%d", column, argID))
			args = append(args, value)
			argID++
		}
	}

	// 変更がない場合は何もせず終了
	if len(setClause) == 0 {
		writeError(w, "No fields to update", http.StatusBadRequest)
		return
	}

	// 更新のためのSQLクエリを構築
	updateUserSQL := fmt.Sprintf("UPDATE users SET %s, updated_at = NOW() WHERE user_id = $%d",
		strings.Join(setClause, ", "), argID)
	args = append(args, userID)

	_, err := db.Exec(updateUserSQL, args...)
	if err != nil {
		writeError(w, fmt.Sprintf("Error updating user: %v", err), http.StatusInternalServerError)
		return
	}

	// 更新後のユーザー情報を取得して返す
	handleGetUser(w, db, userID)
}
