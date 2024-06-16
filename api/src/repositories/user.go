package repositories

import (
	"api/src/models"
	"database/sql"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (repo userRepository) Create(user models.User) (uint64, error) {
	stm, err := repo.db.Prepare(`
		INSERT INTO USER(
			name, nick, email, password
		) VALUES ($1, $2, $3, $4);
	`)

	if err != nil {
		return 0, err
	}

	defer stm.Close()

	result, err := stm.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}
