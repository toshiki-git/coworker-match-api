package databases

import (
	"database/sql"

	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/interfaces/repositories"
)

type sqlUserHobbyRepo struct {
	db *sql.DB
}

func NewSQLUserHobbyRepo(db *sql.DB) repositories.UserHobbyRepo {
	return &sqlUserHobbyRepo{db: db}
}

func (r *sqlUserHobbyRepo) CreateUserHobby(req models.CreateUserHobbyRequest, userId string) (models.CreateUserHobbyResponse, error) {
	// トランザクションを開始
	tx, err := r.db.Begin()
	if err != nil {
		return models.CreateUserHobbyResponse{}, err
	}
	defer tx.Rollback()

	// user_hobbies テーブルから古いエントリを削除
	_, err = tx.Exec(`DELETE FROM user_hobbies WHERE user_id = $1`, userId)
	if err != nil {
		return models.CreateUserHobbyResponse{}, err
	}

	// user_hobbies テーブルに新しいエントリを挿入
	for _, hobbyID := range req.HobbyIds {
		_, err := tx.Exec(`
			INSERT INTO user_hobbies (user_id, hobby_id)
			VALUES ($1, $2)
		`, userId, hobbyID)
		if err != nil {
			return models.CreateUserHobbyResponse{}, err
		}
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return models.CreateUserHobbyResponse{}, err
	}

	var result models.CreateUserHobbyResponse = models.CreateUserHobbyResponse(req)
	return result, nil
}

func (r *sqlUserHobbyRepo) GetAllUserHobby(userId string) ([]models.Hobby, error) {
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

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return []models.Hobby{}, err
	}
	defer rows.Close()

	var response []models.Hobby
	for rows.Next() {
		var hobby models.Hobby
		if err := rows.Scan(&hobby.HobbyId, &hobby.HobbyName); err != nil {
			return []models.Hobby{}, err
		}
		response = append(response, hobby)
	}

	if err := rows.Err(); err != nil {
		return []models.Hobby{}, err
	}

	return response, nil
}

func (r *sqlUserHobbyRepo) UpdateUserHobby(req models.UpdateUserHobbyRequest, userId string) (models.UpdateUserHobbyResponse, error) {
	// トランザクションを開始
	tx, err := r.db.Begin()
	if err != nil {
		return models.UpdateUserHobbyResponse{}, err
	}
	defer tx.Rollback()

	// user_hobbies テーブルから古いエントリを削除
	_, err = tx.Exec(`DELETE FROM user_hobbies WHERE user_id = $1`, userId)
	if err != nil {
		return models.UpdateUserHobbyResponse{}, err
	}

	// user_hobbies テーブルに新しいエントリを挿入
	for _, hobbyId := range req.HobbyIds {
		_, err := tx.Exec(`
			INSERT INTO user_hobbies (user_id, hobby_id)
			VALUES ($1, $2)
		`, userId, hobbyId)
		if err != nil {
			return models.UpdateUserHobbyResponse{}, err
		}
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return models.UpdateUserHobbyResponse{}, err
	}

	return models.UpdateUserHobbyResponse(req), nil
}
