package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"go_web_api/internal/domain/user"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(ctx context.Context, u *user.User) error {
		query := `
			INSERT INTO users (name, email, password, created_at)
			VALUES ($1, $2, $3, $4)
			RETURNING id
	`

	err := r.DB.QueryRowContext(ctx, query,
		u.Name,
		u.Email,
		u.Password, // Agora a senha já vem pronta (criptografada)
		time.Now(),
	).Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
    }
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	query := `
		SELECT id, name, email, password, created_at
		FROM users
		WHERE email = $1
	`

	row := r.DB.QueryRowContext(ctx, query, email)

	var u user.User
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id int64) (*user.User, error) {
	query := `
		SELECT id, name, email, password, created_at
		FROM users
		WHERE id = $1
	`

	row := r.DB.QueryRowContext(ctx, query, id)

	var u user.User
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

var _ user.Repository = (*UserRepository)(nil)
