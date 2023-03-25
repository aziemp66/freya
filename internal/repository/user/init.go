package user

import (
	"context"
	"time"

	userDomain "github.com/aziemp66/freya-be/internal/domain/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImplementation struct {
	db *mongo.Database
}

func NewUserRepositoryImplementation(db *mongo.Database) *UserRepositoryImplementation {
	return &UserRepositoryImplementation{db}
}

func (r *UserRepositoryImplementation) Insert(ctx context.Context, user userDomain.User) (err error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err = r.db.Collection("users").InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryImplementation) FindByID(ctx context.Context, id string) (user userDomain.User, err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return user, err
	}

	err = r.db.Collection("users").FindOne(ctx, bson.M{"_id": objId}).Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepositoryImplementation) FindByEmail(ctx context.Context, email string) (user userDomain.User, err error) {
	err = r.db.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepositoryImplementation) Update(ctx context.Context, user userDomain.User) (err error) {
	user.UpdatedAt = time.Now()

	_, err = r.db.Collection("users").UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryImplementation) UpdateVerifiedEmail(ctx context.Context, id string) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = r.db.Collection("users").UpdateByID(ctx, objId, bson.M{"$set": bson.M{"is_email_verified": true, "updated_at": time.Now()}})

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryImplementation) UpdatePassword(ctx context.Context, id, password string) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = r.db.Collection("users").UpdateByID(ctx, objId, bson.M{"$set": bson.M{"password": password, "updated_at": time.Now()}})

	if err != nil {
		return err
	}

	return nil
}
