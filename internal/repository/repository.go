package repository

import (
	"context"

	"github.com/korasdor/go-ether-test/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users interface {
	Create(ctx context.Context, userData models.UserData) error
	GetByCredentials(ctx context.Context, sidnInData models.SignInData) (models.UserData, error)
	GetById(ctx context.Context, sidnInData models.SignInData) (models.UserData, error)
}

type Repositories struct {
	UsersRepo Users
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		UsersRepo: NewUsersRepo(db),
	}
}
