package services

import (
	"context"

	model "saaster.tech/crud/internal/models"
	repository "saaster.tech/crud/internal/repositories"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetUsers(ctx context.Context) ([]model.User, error) {
	return s.Repo.GetAllUsers(ctx)
}

func (s *UserService) AddUser(ctx context.Context, user *model.User) error {
	return s.Repo.CreateUser(ctx, user)
}
