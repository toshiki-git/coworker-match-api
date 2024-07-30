package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	models "github.com/coworker-match-api/gen/go"
)

func (h *Handler) UserHobbyHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 {
		writeError(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	userID := pathParts[3]

	switch r.Method {
	case http.MethodGet:
		handleGetUserHobbies(w, h.DB, userID)
	case http.MethodPut:
		handlePutUserHobbies(w, r, h.DB, userID)
	default:
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) CreateUserHobbyHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := GetUserID(r.Context())
	if !ok {
		writeError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPost:
		handlePostUserHobby(w, r, h.DB, userID)
	case http.MethodPut:
		handlePutUserHobbies(w, r, h.DB, userID)
	default:
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handlePutUserHobbies(w http.ResponseWriter, r *http.Request, db *sql.DB, userID string) {
	var req models.CreateUserHobbyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	// トランザクションを開始
	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "Error starting transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	// user_hobbies テーブルから古いエントリを削除
	_, err = tx.Exec(`DELETE FROM user_hobbies WHERE user_id = $1`, userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting data: %v", err), http.StatusInternalServerError)
		return
	}

	// user_hobbies テーブルに新しいエントリを挿入
	for _, hobbyID := range req.HobbyIds {
		_, err := tx.Exec(`
			INSERT INTO user_hobbies (user_id, hobby_id)
			VALUES ($1, $2)
		`, userID, hobbyID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error inserting data: %v", err), http.StatusInternalServerError)
			return
		}
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		http.Error(w, "Error committing transaction", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, req)

}

func handlePostUserHobby(w http.ResponseWriter, r *http.Request, db *sql.DB, userID string) {
	var req models.CreateUserHobbyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	// トランザクションを開始
	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "Error starting transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	// user_hobbies テーブルに新しいエントリを挿入
	for _, hobbyID := range req.HobbyIds {
		_, err := tx.Exec(`
			INSERT INTO user_hobbies (user_id, hobby_id)
			VALUES ($1, $2)
		`, userID, hobbyID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error inserting data: %v", err), http.StatusInternalServerError)
			return
		}
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		http.Error(w, "Error committing transaction", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, req)
}

func handleGetUserHobbies(w http.ResponseWriter, db *sql.DB, userID string) {
	query := `
	SELECT
		h.hobby_id,
		h.hobby_name
	FROM
		users u
	LEFT JOIN
		user_hobbies uh ON u.user_id = uh.user_id
	LEFT JOIN
		hobbies h ON uh.hobby_id = h.hobby_id
	WHERE u.user_id = $1;
	`

	rows, err := db.Query(query, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var response []models.Hobby
	for rows.Next() {
		var hobby models.Hobby
		if err := rows.Scan(&hobby.HobbyId, &hobby.HobbyName); err != nil {
			respondWithJSON(w, []models.Hobby{})
			return
		}
		response = append(response, hobby)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, response)
}
