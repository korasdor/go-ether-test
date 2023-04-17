package services

import (
	"context"
	"time"

	"github.com/korasdor/go-ether-test/internal/models"
	"github.com/korasdor/go-ether-test/internal/repository"
	"github.com/korasdor/go-ether-test/pkg/auth"
	"github.com/korasdor/go-ether-test/pkg/blockchain"
	"github.com/korasdor/go-ether-test/pkg/cache"
	"github.com/korasdor/go-ether-test/pkg/hash"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Authorization interface {
	SignUp(ctx context.Context, signUpData models.SignUpData) error
	SignIn(ctx context.Context, signInData models.SignInData, tokenBinding *auth.TokenBinding) (models.Tokens, error)
	RefreshTokens(refreshToken string, tokenBinding *auth.TokenBinding) (models.Tokens, error)
}

type Users interface {
	GetUser(ctx context.Context, userId primitive.ObjectID) (models.UserData, error)
	UpdateUser(ctx context.Context, user models.UserData) (models.UserData, error)
	DeleteUser(ctx context.Context, userId primitive.ObjectID) error
}

type Services struct {
	AuthorizationService Authorization
	UsersService         Users
}

type Deps struct {
	Repos             *repository.Repositories
	Cache             cache.Cache
	Hasher            hash.PasswordHasher
	TokenManager      auth.TokenManager
	BlockchainManager blockchain.BlockchainManager
	AccessTokenTTL    time.Duration
	RefreshTokenTTL   time.Duration
}

func NewServices(deps *Deps) *Services {
	return &Services{
		AuthorizationService: NewAuthorizationService(deps.Repos, deps.Hasher, deps.TokenManager, deps.AccessTokenTTL, deps.RefreshTokenTTL),
		UsersService:         NewUsersService(deps.Repos, deps.BlockchainManager),
	}
}
