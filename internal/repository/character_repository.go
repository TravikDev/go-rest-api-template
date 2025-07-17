package repository

import (
	"database/sql"

	"go-rest-api-template/internal/models"
)

// CharacterRepositoryInterface defines operations for character persistence.
type CharacterRepositoryInterface interface {
	Create(c *models.Character) error
	GetByUserID(userID int) (*models.Character, error)
}

// CharacterRepository implements CharacterRepositoryInterface using sql.DB.
type CharacterRepository struct {
	db *sql.DB
}

// NewCharacterRepository creates a CharacterRepository.
func NewCharacterRepository(db *sql.DB) *CharacterRepository {
	return &CharacterRepository{db: db}
}

// Create inserts a new character for a user.
func (r *CharacterRepository) Create(c *models.Character) error {
	return r.db.QueryRow(
		`INSERT INTO characters (user_id, nickname, level, experience, pos_x, pos_y, pos_z)
         VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`,
		c.UserID, c.Nickname, c.Level, c.Experience, c.PosX, c.PosY, c.PosZ,
	).Scan(&c.ID)
}

// GetByUserID returns a character associated with the given user.
func (r *CharacterRepository) GetByUserID(userID int) (*models.Character, error) {
	row := r.db.QueryRow(
		`SELECT id, user_id, nickname, level, experience, pos_x, pos_y, pos_z
         FROM characters WHERE user_id=$1`, userID,
	)
	ch := &models.Character{}
	if err := row.Scan(&ch.ID, &ch.UserID, &ch.Nickname, &ch.Level, &ch.Experience, &ch.PosX, &ch.PosY, &ch.PosZ); err != nil {
		return nil, err
	}
	return ch, nil
}
