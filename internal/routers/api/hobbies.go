package api

import (
	"database/sql"
	"net/http"
)

type Hobby struct {
	HobbyID   string `json:"hobby_id"`
	HobbyName string `json:"hobby_name"`
}

type Category struct {
	CategoryID string  `json:"category_id"`
	Hobbies    []Hobby `json:"hobbies"`
}

type Response map[string]Category

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
		h.hobby_id,
		h.hobby_name
	FROM
		hobbies h
	LEFT JOIN
		categories c ON c.category_id = h.category_id;
	`

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	response := make(Response)
	for rows.Next() {
		var categoryID, categoryName, hobbyID, hobbyName sql.NullString
		if err := rows.Scan(&categoryID, &categoryName, &hobbyID, &hobbyName); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !categoryID.Valid || !categoryName.Valid {
			continue
		}

		categoryKey := convertToKey(categoryName.String)
		if category, exists := response[categoryKey]; exists {
			if hobbyID.Valid && hobbyName.Valid {
				category.Hobbies = append(category.Hobbies, Hobby{
					HobbyID:   hobbyID.String,
					HobbyName: hobbyName.String,
				})
			}
			response[categoryKey] = category
		} else {
			newCategory := Category{
				CategoryID: categoryID.String,
				Hobbies:    []Hobby{},
			}
			if hobbyID.Valid && hobbyName.Valid {
				newCategory.Hobbies = append(newCategory.Hobbies, Hobby{
					HobbyID:   hobbyID.String,
					HobbyName: hobbyName.String,
				})
			}
			response[categoryKey] = newCategory
		}
	}

	respondWithJSON(w, response)
}

func convertToKey(categoryName string) string {
	switch categoryName {
	case "インドア":
		return "indoor"
	case "ゲーム":
		return "games"
	case "技術":
		return "technicalHobbies"
	case "スポーツ":
		return "sports"
	case "アウトドア":
		return "outdoor"
	case "音楽":
		return "music"
	default:
		return "unknown"
	}
}

func handlePostHobby(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//TODO: 趣味を提案できるAPIを実装する
}
