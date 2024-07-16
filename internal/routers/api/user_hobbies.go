package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type CreateUserHobbyRequest struct {
	UserID   string   `json:"user_id"`
	HobbyIDs []string `json:"hobby_ids"`
}

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
	default:
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) CreateUserHobbyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	handlePostUserHobby(w, r, h.DB)
}

func handlePostUserHobby(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var req CreateUserHobbyRequest

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
	for _, hobbyID := range req.HobbyIDs {
		_, err := tx.Exec(`
			INSERT INTO user_hobbies (user_id, hobby_id)
			VALUES ($1, $2)
		`, req.UserID, hobbyID)
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

	var hobbies []Hobby
	for rows.Next() {
		var hobby Hobby
		if err := rows.Scan(&hobby.HobbyID, &hobby.HobbyName); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		hobbies = append(hobbies, hobby)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, hobbies)
}
