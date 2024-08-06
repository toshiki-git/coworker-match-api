package databases

import (
	"database/sql"
	"encoding/json"

	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/interfaces/repositories"
)

type sqlHobbyRepository struct {
	db *sql.DB
}

func NewSQLHobbyRepository(db *sql.DB) repositories.HobbyRepository {
	return &sqlHobbyRepository{db: db}
}

func (hr *sqlHobbyRepository) GetAllHobby() ([]models.GetHobbyResponseInner, error) {
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
	rows, err := hr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allHobby []models.GetHobbyResponseInner

	for rows.Next() {
		var data models.GetHobbyResponseInner
		var hobbiesData []byte

		if err := rows.Scan(&data.CategoryId, &data.CategoryName, &hobbiesData); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(hobbiesData, &data.Hobbies); err != nil {
			return nil, err
		}

		allHobby = append(allHobby, data)
	}
	return allHobby, nil
}
