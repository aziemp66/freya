package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	User struct {
		ID              primitive.ObjectID `bson:"_id"`
		FirstName       string             `bson:"first_name"`
		LastName        string             `bson:"last_name"`
		Email           string             `bson:"email"`
		Password        string             `bson:"password"`
		BirthDay        time.Time          `bson:"birthday"`
		IsEmailVerified bool               `bson:"is_email_verified"`
		Role            role               `bson:"role"`
		CreatedAt       time.Time          `bson:"created_at"`
		UpdatedAt       time.Time          `bson:"updated_at"`
	}

	role string
)
