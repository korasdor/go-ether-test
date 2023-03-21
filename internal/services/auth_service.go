package services

import "github.com/korasdor/go-ether-test/internal/models"

type AuthorizationService struct {
}

func NewAuthorizationService() *AuthorizationService {
	return &AuthorizationService{}
}

func (s *AuthorizationService) SingUp(signUpData models.SignUpData) error {
	return nil
}
func (s *AuthorizationService) SingIn(signInData models.SignInData) (models.Tokens, error) {
	return models.Tokens{}, nil
}
func (s *AuthorizationService) RefreshTokens() (models.Tokens, error) {
	return models.Tokens{}, nil
}
