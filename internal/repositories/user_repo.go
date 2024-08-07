package repositories

import (
	"database/sql"

	models "github.com/coworker-match-api/gen/go"
)

type IUserRepo interface {
	CreateUser(userId string, req *models.CreateUserRequest) (*models.CreateUserResponse, error)
	GetUserById(userId string) (*models.GetUserResponse, error)
	UpdateUser(userId string, req *models.UpdateUserRequest) (*models.UpdateUserResponse, error)
	IsUserExist(userId string) (bool, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) IUserRepo {
	return &userRepo{db: db}
}

func (ur *userRepo) CreateUser(userId string, req *models.CreateUserRequest) (*models.CreateUserResponse, error) {
	const query = `
				INSERT INTO 
					users (user_id, user_name, email, avatar_url)
				VALUES 
					($1, $2, $3, $4)
				RETURNING
					user_id, user_name, email, avatar_url
				`
	var user models.User
	err := ur.db.QueryRow(query, userId, req.UserName, req.Email, req.AvatarUrl).Scan(
		&user.UserId, &user.UserName, &user.Email, &user.AvatarUrl,
	)
	if err != nil {
		return nil, err
	}

	return &models.CreateUserResponse{User: user}, nil
}

func (ur *userRepo) GetUserById(userId string) (*models.GetUserResponse, error) {
	query := `
			SELECT
				user_id, user_name, email, avatar_url
			FROM 
				users
			WHERE
				user_id = $1
			`

	var user models.User
	ur.db.QueryRow(query, userId).Scan(
		&user.UserId, &user.UserName, &user.Email, &user.AvatarUrl,
	)

	return &models.GetUserResponse{User: user}, nil
}

func (r *userRepo) UpdateUser(userId string, req *models.UpdateUserRequest) (*models.UpdateUserResponse, error) {
	return nil, nil
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
