package repositories

import (
	"database/sql"
	"encoding/json"

	models "github.com/coworker-match-api/gen/go"
)

type IHobbyRepo interface {
	//TODO: CreateHobby()
	GetAllHobby() ([]*models.GetHobbyResponseInner, error)
}

type hobbyRepo struct {
	db *sql.DB
}

func NewHobbyRepo(db *sql.DB) IHobbyRepo {
	return &hobbyRepo{db: db}
}

func (hr *hobbyRepo) GetAllHobby() ([]*models.GetHobbyResponseInner, error) {
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

	var allHobby []*models.GetHobbyResponseInner

	for rows.Next() {
		var data models.GetHobbyResponseInner
		var hobbiesData []byte

		if err := rows.Scan(&data.CategoryId, &data.CategoryName, &hobbiesData); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(hobbiesData, &data.Hobbies); err != nil {
			return nil, err
		}

		allHobby = append(allHobby, &data)
	}
	return allHobby, nil
}
