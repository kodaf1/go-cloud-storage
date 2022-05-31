package user

import (
	"context"
)

type Service interface {
	Create(ctx context.Context, dto *SignUpDTO) *User
	GetByEmailAndPassword(ctx context.Context, email, password string) (*User, error)
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{storage: storage}
}

func (s *service) Create(ctx context.Context, dto *SignUpDTO) *User {
	panic("IMPLEMENT ME")
}

func (s *service) GetByEmailAndPassword(ctx context.Context, email, password string) (*User, error) {
	panic("IMPLEMENT ME")
}
