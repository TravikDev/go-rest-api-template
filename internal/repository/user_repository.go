package repository

import (
	"database/sql"

	"go-rest-api-template/internal/models"
)

type UserRepositoryInterface interface {
	Create(user *models.User) error
	GetByID(id int) (*models.User, error)
	List() ([]*models.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	_, err := r.db.Exec(`INSERT INTO users (username) VALUES ($1)`, user.Username)
	return err
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
	row := r.db.QueryRow(`SELECT id, username FROM users WHERE id=$1`, id)
	u := &models.User{}
	if err := row.Scan(&u.ID, &u.Username); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	row := r.db.QueryRow(`SELECT id, username FROM users WHERE username=$1`, username)
	u := &models.User{}
	if err := row.Scan(&u.ID, &u.Username); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) List() ([]*models.User, error) {
	rows, err := r.db.Query(`SELECT id, username FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*models.User{}
	for rows.Next() {
		u := &models.User{}
		if err := rows.Scan(&u.ID, &u.Username); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}
