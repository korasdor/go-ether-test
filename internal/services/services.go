package services

import (
	"github.com/korasdor/go-ether-test/internal/models"
	"github.com/korasdor/go-ether-test/pkg/cache"
)

type Authorization interface {
	SingUp(signUpData models.SignUpData) error
	SingIn(signInData models.SignInData) (models.Tokens, error)
	RefreshTokens() (models.Tokens, error)
}

type Services struct {
	AuthorizationService Authorization
}

type Deps struct {
	Repos interface{}
	Cache cache.Cache
}

func NewServices(deps *Deps) *Services {
	return &Services{
		AuthorizationService: NewAuthorizationService(),
	}
}
