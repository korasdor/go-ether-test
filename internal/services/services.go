package services

import (
	"context"

	"github.com/korasdor/go-ether-test/internal/models"
	"github.com/korasdor/go-ether-test/internal/repository"
	"github.com/korasdor/go-ether-test/pkg/cache"
	"github.com/korasdor/go-ether-test/pkg/hash"
)

type Authorization interface {
	SingUp(ctx context.Context, signUpData models.SignUpData) error
	SingIn(ctx context.Context, signInData models.SignInData) (models.Tokens, error)
	RefreshTokens() (models.Tokens, error)
}

type Services struct {
	AuthorizationService Authorization
}

type Deps struct {
	Repos  *repository.Repositories
	Cache  cache.Cache
	Hasher hash.PasswordHasher
}

func NewServices(deps *Deps) *Services {
	return &Services{
		AuthorizationService: NewAuthorizationService(deps.Repos, deps.Hasher),
	}
}
