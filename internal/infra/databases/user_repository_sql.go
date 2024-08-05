package databases

import (
	"database/sql"
	"fmt"
	"strings"

	models "github.com/coworker-match-api/gen/go"
	"github.com/coworker-match-api/internal/interfaces/repositories"
)

type sqlUserRepository struct {
	db *sql.DB
}

func NewSQLUserRepository(db *sql.DB) repositories.UserRepository {
	return &sqlUserRepository{db: db}
}

func (ur *sqlUserRepository) CreateUser(user models.User) (models.User, error) {
	const query = `
				INSERT INTO 
					users (user_id, user_name, email, avatar_url)
				VALUES 
					($1, $2, $3, $4)
				RETURNING
					user_id, user_name, email, avatar_url
				`

	row := ur.db.QueryRow(query, user.UserName, user.Email, user.AvatarUrl)
	err := row.Scan(&user.UserId, &user.UserName, &user.Email, &user.AvatarUrl)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur *sqlUserRepository) GetUserByID(userId string) (models.User, error) {
	query := `
			SELECT
				user_id, user_name, email, avatar_url
			FROM 
				users
			WHERE
				user_id = $1
			`
	row := ur.db.QueryRow(query, userId)
	var user models.User
	err := row.Scan(&user.UserId, &user.UserName, &user.Email, &user.AvatarUrl)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *sqlUserRepository) UpdateUser(userID string, updates map[string]interface{}) error {
	setClause, args := buildUpdateClause(updates)
	if len(setClause) == 0 {
		return fmt.Errorf("no fields to update")
	}

	args = append(args, userID)
	const updateUserSQLTemplate = "UPDATE users SET %s, updated_at = NOW() WHERE user_id = $%d"
	updateUserSQL := fmt.Sprintf(updateUserSQLTemplate, strings.Join(setClause, ", "), len(args))

	_, err := r.db.Exec(updateUserSQL, args...)
	return err
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
