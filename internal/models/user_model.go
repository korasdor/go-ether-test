package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserData struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name"`
	Email        string             `json:"email" bson:"email"`
	Phone        string             `json:"phone" bson:"phone"`
	Password     string             `json:"-" bson:"password"`
	RegisteredAt time.Time          `json:"registeredAt" bson:"registeredAt"`
	LastVisitAt  time.Time          `json:"lastVisitAt" bson:"lastVisitAt"`
	Wallets      []Wallet           `json:"wallets" bson:"wallets"`
}

type Wallet struct {
	ID      string `json:"id" bson:"_id"`
	Address string `json:"address" bson:"address"`
}
