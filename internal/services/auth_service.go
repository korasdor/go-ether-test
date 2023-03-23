package services

import (
	"context"
	"fmt"
	"time"

	"github.com/korasdor/go-ether-test/internal/models"
	"github.com/korasdor/go-ether-test/internal/repository"
	"github.com/korasdor/go-ether-test/pkg/hash"
)

type AuthorizationService struct {
	repo   *repository.Repositories
	hasher hash.PasswordHasher
}

func NewAuthorizationService(repo *repository.Repositories, hasher hash.PasswordHasher) *AuthorizationService {
	return &AuthorizationService{
		repo:   repo,
		hasher: hasher,
	}
}

func (s *AuthorizationService) SingUp(ctx context.Context, signUpData models.SignUpData) error {
	passwordHash, err := s.hasher.Hash(signUpData.Password)
	if err != nil {
		return err
	}

	userData := models.UserData{
		Name:         signUpData.Name,
		Password:     passwordHash,
		Phone:        signUpData.Phone,
		Email:        signUpData.Email,
		RegisteredAt: time.Now(),
		LastVisitAt:  time.Now(),
	}

	return s.repo.UsersRepo.Create(ctx, userData)
}

func (s *AuthorizationService) SingIn(ctx context.Context, signInData models.SignInData) (models.Tokens, error) {
	token := models.Tokens{}

	passwordHash, err := s.hasher.Hash(signInData.Password)
	if err != nil {
		return token, err
	}

	signInData.Password = passwordHash
	user, err := s.repo.UsersRepo.GetByCredentials(ctx, signInData)
	if err != nil {
		return token, err
	}

	fmt.Println(user)

	return token, nil
}

func (s *AuthorizationService) RefreshTokens() (models.Tokens, error) {
	return models.Tokens{}, nil
}
