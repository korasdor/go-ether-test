package repository

import (
	"context"
	"errors"

	"github.com/korasdor/go-ether-test/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepo struct {
	db *mongo.Collection
}

func NewUsersRepo(db *mongo.Database) *UsersRepo {
	return &UsersRepo{
		db: db.Collection(usersCollection),
	}
}

func (r *UsersRepo) Create(ctx context.Context, userData models.UserData) error {
	_, err := r.db.InsertOne(ctx, userData)
	if mongo.IsDuplicateKeyError(err) {
		return models.ErrUserAlreadyExists
	}

	return err
}

func (r *UsersRepo) GetByCredentials(ctx context.Context, sidnInData models.SignInData) (models.UserData, error) {
	var user models.UserData

	if err := r.db.FindOne(ctx, sidnInData).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user, models.ErrUserNotFound
		}

		return user, err
	}

	return user, nil
}

func (r *UsersRepo) GetById(ctx context.Context, userId primitive.ObjectID) (models.UserData, error) {
	var user models.UserData

	if err := r.db.FindOne(ctx, bson.M{"_id": userId}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user, models.ErrUserNotFound
		}

		return user, err
	}

	return user, nil
}

func (r *UsersRepo) UpdateUser(ctx context.Context, user models.UserData) (models.UserData, error) {
	res, err := r.db.UpdateOne(ctx,
		bson.M{"_id": user.ID},
		bson.M{"$set": user},
	)
	if err != nil {
		return user, err
	}

	if res.MatchedCount == 0 {
		return user, models.ErrUserNotFound
	}

	return user, nil
}
