package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	models "github.com/coworker-match-api/gen/go"
)

func (h *Handler) HobbyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetHobbies(w, h.DB)
	case http.MethodPost:
		handlePostHobby(w, r, h.DB)
	default:
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetHobbies(w http.ResponseWriter, db *sql.DB) {
	query := `
		SELECT 
			c.category_id, 
			c.category_name, 
			COALESCE(json_agg(json_build_object('hobbyId', h.hobby_id, 'hobbyName', h.hobby_name)) FILTER (WHERE h.hobby_id IS NOT NULL), '[]'::json) AS hobbies
		FROM 
			categories c
		LEFT JOIN 
			hobbies h 
		ON 
			c.category_id = h.category_id
		GROUP BY 
			c.category_id, c.category_name;
    	`

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var response []models.GetHobbyResponseInner

	for rows.Next() {
		var data models.GetHobbyResponseInner
		var hobbiesData []byte

		if err := rows.Scan(&data.CategoryId, &data.CategoryName, &hobbiesData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.Unmarshal(hobbiesData, &data.Hobbies); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response = append(response, data)
	}

	respondWithJSON(w, response)
}

func handlePostHobby(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//TODO: 趣味を提案できるAPIを実装する
}
