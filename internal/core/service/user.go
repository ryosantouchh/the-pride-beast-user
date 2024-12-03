package service

import (
	"github.com/ryosantouchh/the-pride-beast-user/internal/core/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) FindUserById() {
	// userById, err := s.repo.Get()
}
