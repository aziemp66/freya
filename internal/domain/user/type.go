package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	User struct {
		ID              primitive.ObjectID `bson:"_id,omitempty"`
		FirstName       string             `bson:"first_name,omitempty"`
		LastName        string             `bson:"last_name,omitempty"`
		Email           string             `bson:"email,omitempty"`
		Password        string             `bson:"password,omitempty"`
		BirthDay        time.Time          `bson:"birthday,omitempty"`
		IsEmailVerified bool               `bson:"is_email_verified"`
		Role            role               `bson:"role,omitempty"`
		CreatedAt       time.Time          `bson:"created_at,omitempty"`
		UpdatedAt       time.Time          `bson:"updated_at,omitempty"`
	}

	role string
)
