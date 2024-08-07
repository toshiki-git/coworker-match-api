package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	models "github.com/coworker-match-api/gen/go"
)

type IUserRepo interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserById(userId string) (*models.User, error)
	UpdateUser(userId string, updates map[string]interface{}) error
	IsUserExist(userId string) (bool, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) IUserRepo {
	return &userRepo{db: db}
}

func (ur *userRepo) CreateUser(user *models.User) (*models.User, error) {
	const query = `
				INSERT INTO 
					users (user_id, user_name, email, avatar_url)
				VALUES 
					($1, $2, $3, $4)
				RETURNING
					user_id, user_name, email, avatar_url
				`

	row := ur.db.QueryRow(query, user.UserId, user.UserName, user.Email, user.AvatarUrl)
	err := row.Scan(&user.UserId, &user.UserName, &user.Email, &user.AvatarUrl)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepo) GetUserById(userId string) (*models.User, error) {
	query := `
			SELECT
				user_id, user_name, email, avatar_url
			FROM 
				users
			WHERE
				user_id = $1
			`
	row := ur.db.QueryRow(query, userId)
	user := &models.User{}
	err := row.Scan(&user.UserId, &user.UserName, &user.Email, &user.AvatarUrl)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepo) UpdateUser(userId string, updates map[string]interface{}) error {
	setClause, args := buildUpdateClause(updates)
	if len(setClause) == 0 {
		return fmt.Errorf("no fields to update")
	}

	args = append(args, userId)
	const updateUserSQLTemplate = "UPDATE users SET %s, updated_at = NOW() WHERE user_id = $%d"
	updateUserSQL := fmt.Sprintf(updateUserSQLTemplate, strings.Join(setClause, ", "), len(args))

	_, err := r.db.Exec(updateUserSQL, args...)
	return err
}

func (r *userRepo) IsUserExist(userId string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE user_id = $1)`
	row := r.db.QueryRow(query, userId)
	var isExist bool
	err := row.Scan(&isExist)
	if err != nil {
		return false, err
	}
	return isExist, nil
}

func buildUpdateClause(updates map[string]interface{}) ([]string, []interface{}) {
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
		if value, ok := updates[field]; ok {
			setClause = append(setClause, fmt.Sprintf("%s = $%d", column, argID))
			args = append(args, value)
			argID++
		}
	}

	return setClause, args
}
