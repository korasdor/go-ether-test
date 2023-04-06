package services

import (
	"context"
	"errors"
	"time"

	"github.com/korasdor/go-ether-test/internal/models"
	"github.com/korasdor/go-ether-test/internal/repository"
	"github.com/korasdor/go-ether-test/pkg/auth"
	"github.com/korasdor/go-ether-test/pkg/hash"
)

type AuthorizationService struct {
	repo            *repository.Repositories
	hasher          hash.PasswordHasher
	tokenManager    auth.TokenManager
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewAuthorizationService(
	repo *repository.Repositories,
	hasher hash.PasswordHasher,
	tokenManager auth.TokenManager,
	accessTokenTTL time.Duration,
	refeshTokenTTL time.Duration,
) *AuthorizationService {

	return &AuthorizationService{
		repo:            repo,
		hasher:          hasher,
		tokenManager:    tokenManager,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refeshTokenTTL,
	}
}

func (s *AuthorizationService) SignUp(ctx context.Context, signUpData models.SignUpData) error {
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

func (s *AuthorizationService) SignIn(ctx context.Context, signInData models.SignInData, tokenBinding *auth.TokenBinding) (models.Tokens, error) {
	var token models.Tokens

	passwordHash, err := s.hasher.Hash(signInData.Password)
	if err != nil {
		return token, err
	}

	signInData.Password = passwordHash
	user, err := s.repo.UsersRepo.GetByCredentials(ctx, signInData)
	if err != nil {
		return token, err
	}

	token, err = s.generateTokens(user.ID.Hex(), tokenBinding)
	if err != nil {
		return token, err
	}

	return token, nil
}

func (s *AuthorizationService) RefreshTokens(refreshToken string, tokenBinding *auth.TokenBinding) (models.Tokens, error) {
	var token models.Tokens

	tokenData, err := s.tokenManager.ParseJWT(refreshToken)
	if err != nil {
		return token, err
	}

	if tokenBinding.IPAddr != tokenData.TokenBinding.IPAddr || tokenBinding.UserAgent != tokenData.TokenBinding.UserAgent {
		return token, errors.New("token binding not the same, the difference is in the IP address and user agent")
	}

	token, err = s.generateTokens(tokenData.UserId, tokenData.TokenBinding)
	if err != nil {
		return token, err
	}

	return token, nil
}

func (s *AuthorizationService) generateTokens(userId string, tokenBinding *auth.TokenBinding) (models.Tokens, error) {
	var token models.Tokens

	accessToken, err := s.tokenManager.NewJWT(userId, tokenBinding, s.accessTokenTTL)
	if err != nil {
		return token, err
	}

	refreshToken, err := s.tokenManager.NewJWT(userId, tokenBinding, s.refreshTokenTTL)
	if err != nil {
		return token, err
	}

	token = models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return token, nil
}
