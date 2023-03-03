package user

import (
	"github.com/aziemp66/freya-be/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	User struct {
		ID              primitive.ObjectID `bson:"_id"`
		FirstName       string             `bson:"first_name"`
		LastName        string             `bson:"last_name"`
		Email           string             `bson:"email"`
		Password        string             `bson:"password"`
		IsEmailVerified bool               `bson:"is_email_verified"`
		Role            role               `bson:"role"`

		Timestamp
	}

	role      string
	Timestamp = domain.Timestamp
)
