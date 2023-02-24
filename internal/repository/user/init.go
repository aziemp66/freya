package user

import (
	"context"

	userDomain "github.com/aziemp66/freya-be/internal/domain/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImplementation struct {
	db *mongo.Database
}

func NewUserRepositoryImplementation(db *mongo.Database) *UserRepositoryImplementation {
	return &UserRepositoryImplementation{db}
}

func (r *UserRepositoryImplementation) Insert(ctx context.Context, user userDomain.User) (err error) {
	_, err = r.db.Collection("users").InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryImplementation) FindByEmail(ctx context.Context, email string) (user userDomain.User, err error) {
	err = r.db.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepositoryImplementation) Update(ctx context.Context, user userDomain.User) (err error) {
	_, err = r.db.Collection("users").ReplaceOne(ctx, bson.M{"_id": user.ID}, user)

	if err != nil {
		return err
	}

	return nil
}
