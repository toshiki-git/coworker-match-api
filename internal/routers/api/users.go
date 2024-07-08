package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/coworker-match-api/internal/models"
)

func (h *Handler) UsersHandler(w http.ResponseWriter, r *http.Request) {
	getUserSQL := `SELECT * FROM users`
	rows, err := h.DB.Query(getUserSQL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error querying data: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.UserID, &user.UserName, &user.Email, &user.AvatarURL, &user.Age, &user.Gender, &user.Birthplace, &user.JobType, &user.LINEAccount, &user.DiscordAccount, &user.Biography, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error scanning data: %v", err), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Error with rows: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
	}
}
