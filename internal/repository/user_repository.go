package repository

import (
	"database/sql"

	"go-rest-api-template/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	_, err := r.db.Exec(`INSERT INTO users (name, email) VALUES ($1, $2)`, user.Name, user.Email)
	return err
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
	row := r.db.QueryRow(`SELECT id, name, email FROM users WHERE id=$1`, id)
	u := &models.User{}
	if err := row.Scan(&u.ID, &u.Name, &u.Email); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	row := r.db.QueryRow(`SELECT id, name, email FROM users WHERE email=$1`, email)
	u := &models.User{}
	if err := row.Scan(&u.ID, &u.Name, &u.Email); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) GetByName(name string) (*models.User, error) {
	row := r.db.QueryRow(`SELECT id, name, email FROM users WHERE name=$1`, name)
	u := &models.User{}
	if err := row.Scan(&u.ID, &u.Name, &u.Email); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) List() ([]*models.User, error) {
	rows, err := r.db.Query(`SELECT id, name, email FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*models.User{}
	for rows.Next() {
		u := &models.User{}
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}
