package user

import (
	"context"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) CreateUser(ctx context.Context, u *User) error {
	if !strings.Contains(u.Email, "@") {
		return ErrInvalidEmail
	}

	if len(u.Password) < 6 {
		return ErrInvalidPassword
	}

	existing, _ := s.repo.FindByEmail(ctx, u.Email)
	if existing != nil {
		return ErrEmailAlreadyUsed
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return s.repo.Create(ctx, u)
}

func (s *Service) GetUserByID(ctx context.Context, id int64) (*User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *Service) UpdateUser(ctx context.Context, u *User) error {
	return s.repo.Update(ctx, u)
}

func (s *Service) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *Service) Login(ctx context.Context, email, password string) (*User, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}
