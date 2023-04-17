package services

import (
	"context"

	"github.com/korasdor/go-ether-test/internal/models"
	"github.com/korasdor/go-ether-test/internal/repository"
	"github.com/korasdor/go-ether-test/pkg/blockchain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersService struct {
	repo              *repository.Repositories
	blockchainManager blockchain.BlockchainManager
}

func NewUsersService(repo *repository.Repositories, blockchainManager blockchain.BlockchainManager) *UsersService {
	return &UsersService{
		repo:              repo,
		blockchainManager: blockchainManager,
	}
}

func (s *UsersService) GetUser(ctx context.Context, userId primitive.ObjectID) (models.UserData, error) {
	return s.repo.UsersRepo.GetById(ctx, userId)
}

func (s *UsersService) UpdateUser(ctx context.Context, user models.UserData) (models.UserData, error) {
	return s.repo.UsersRepo.UpdateUser(ctx, user)
}

func (s *UsersService) DeleteUser(ctx context.Context, userId primitive.ObjectID) error {
	return nil
}
