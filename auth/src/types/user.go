package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id               primitive.ObjectID `json:"-" bson:"_id"`
	Username         string             `json:"username" bson:"username"`
	Email            string             `json:"email" bson:"email"`
	Password         string             `json:"password" bson:"-"`
	HashedPassword   string             `json:"-" bson:"hashedPassword"`
	Salt             string             `json:"-" bson:"salt"`
	Verified         bool               `bson:"verified"`
	VerificationCode uint64             `bson:"verificationCode"`
	CreatedAt        time.Time          `bson:"createdAt"`
	UpdatedAt        time.Time          `bson:"updatedAt"`
}
