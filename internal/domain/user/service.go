package user

import (
	"context"
	"strings"
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

	return s.repo.Create(ctx, u)
}
