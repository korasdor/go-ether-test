package services

import (
	"context"

	"github.com/korasdor/go-ether-test/internal/models"
	"github.com/korasdor/go-ether-test/internal/repository"
)

type UsersService struct {
	repo *repository.Repositories
}

func NewUsersService(repo *repository.Repositories) *UsersService {
	return &UsersService{
		repo: repo,
	}
}

func (s *UsersService) GetUser(ctx context.Context, userId string) (models.UserData, error) {
	return models.UserData{}, nil
}

func (s *UsersService) DeleteUser(ctx context.Context, userId string) error {
	return nil
}
