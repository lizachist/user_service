package repository

import (
	"database/sql"
	"errors"
	"user_service/internal/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(database *sql.DB) domain.UserRepository {
	return &userRepository{db: database}
}

func (r *userRepository) Create(user *domain.User) error {
	query := `
        INSERT INTO users (username, email, password_hash, first_name, last_name, is_active)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(query,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.FirstName,
		user.LastName,
		user.IsActive,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	return err
}

func (r *userRepository) GetByID(id int) (*domain.User, error) {
	user := &domain.User{}
	query := `
        SELECT id, username, email, password_hash, first_name, last_name, is_active, created_at, updated_at
        FROM users WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Update(user *domain.User) error {
	query := `
        UPDATE users
        SET username = $1, email = $2, first_name = $3, last_name = $4, is_active = $5, updated_at = CURRENT_TIMESTAMP
        WHERE id = $6`

	result, err := r.db.Exec(query,
		user.Username,
		user.Email,
		user.FirstName,
		user.LastName,
		user.IsActive,
		user.ID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}
