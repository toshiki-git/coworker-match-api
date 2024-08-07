package repositories

import (
	"database/sql"

	models "github.com/coworker-match-api/gen/go"
)

type IUserHobbyRepo interface {
	CreateUserHobby(*models.CreateUserHobbyReq, string) (*models.CreateUserHobbyRes, error)
	GetAllUserHobby(string) ([]*models.Hobby, error)
	UpdateUserHobby(*models.UpdateUserHobbyReq, string) (*models.UpdateUserHobbyRes, error)
}

type userHobbyRepo struct {
	db *sql.DB
}

func NewUserHobbyRepo(db *sql.DB) IUserHobbyRepo {
	return &userHobbyRepo{db: db}
}

func (uhr *userHobbyRepo) CreateUserHobby(req *models.CreateUserHobbyReq, userId string) (*models.CreateUserHobbyRes, error) {
	// トランザクションを開始
	tx, err := uhr.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// user_hobbies テーブルから古いエントリを削除
	_, err = tx.Exec(`DELETE FROM user_hobbies WHERE user_id = $1`, userId)
	if err != nil {
		return nil, err
	}

	// user_hobbies テーブルに新しいエントリを挿入
	for _, hobbyID := range req.HobbyIds {
		_, err := tx.Exec(`
			INSERT INTO user_hobbies (user_id, hobby_id)
			VALUES ($1, $2)
		`, userId, hobbyID)
		if err != nil {
			return nil, err
		}
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	var result models.CreateUserHobbyRes = models.CreateUserHobbyRes(*req)
	return &result, nil
}

func (uhr *userHobbyRepo) GetAllUserHobby(userId string) ([]*models.Hobby, error) {
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

	rows, err := uhr.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var response []*models.Hobby
	for rows.Next() {
		var hobby models.Hobby
		if err := rows.Scan(&hobby.HobbyId, &hobby.HobbyName); err != nil {
			return nil, err
		}
		response = append(response, &hobby)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return response, nil
}

func (uhr *userHobbyRepo) UpdateUserHobby(req *models.UpdateUserHobbyReq, userId string) (*models.UpdateUserHobbyRes, error) {
	// トランザクションを開始
	tx, err := uhr.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// user_hobbies テーブルから古いエントリを削除
	_, err = tx.Exec(`DELETE FROM user_hobbies WHERE user_id = $1`, userId)
	if err != nil {
		return nil, err
	}

	// user_hobbies テーブルに新しいエントリを挿入
	for _, hobbyId := range req.HobbyIds {
		_, err := tx.Exec(`
			INSERT INTO user_hobbies (user_id, hobby_id)
			VALUES ($1, $2)
		`, userId, hobbyId)
		if err != nil {
			return nil, err
		}
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	var result models.UpdateUserHobbyRes = models.UpdateUserHobbyRes(*req)
	return &result, nil
}
