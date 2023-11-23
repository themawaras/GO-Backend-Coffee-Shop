package repositories

import (
	"BE-Golang-Coffee-Shop/internal/models"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	*sqlx.DB
}

func InitializeUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) RepositoryGetAllUser() ([]models.UserModel, error) {
	result := []models.UserModel{}
	query := `SELECT user_id, user_fullname, user_email, user_address, user_picture FROM users`

	err := r.Select(&result, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *UserRepository) RepositoryGetUserID(id string) ([]models.UserModel, error) {
	result := []models.UserModel{}
	query := `SELECT user_id, user_fullname, user_email, user_address, user_picture FROM users WHERE user_id = $1`
	err := r.Select(&result, query, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *UserRepository) RepositoryCreateUser(body *models.UserModel) (sql.Result, error) {
	query := `INSERT INTO users (user_fullname, user_email, user_phone, user_address, user_password) VALUES (:user_fullname, :user_email, :user_phone, :user_address, :user_password)`

	result, err := r.NamedExec(query, body)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *UserRepository) RepositoryUpdateUser(body *models.UserModel, id string) (sql.Result, error) {
	query := `UPDATE users SET user_fullname = :user_fullname, user_phone = :user_phone, user_address = :user_address, user_password = :user_password, updated_at = now() WHERE user_id =` + id
	result, err := r.NamedExec(query, body)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *UserRepository) RepositoryDeleteUser(body *models.UserModel, id string) (sql.Result, error) {
	query := `DELETE FROM users WHERE user_id = ` + id
	result, err := r.NamedExec(query, body)
	if err != nil {
		return result, err
	}
	return result, nil
}
