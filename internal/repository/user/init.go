package user

import (
	"context"
	"time"

	errorCommon "github.com/aziemp66/freya-be/common/error"
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
		return errorCommon.NewInvariantError("Failed to insert user")
	}

	return nil
}

func (r *UserRepositoryImplementation) FindByID(ctx context.Context, id string) (user userDomain.User, err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return user, errorCommon.NewInvariantError("Invalid user id format")
	}

	err = r.db.Collection("users").FindOne(ctx, bson.M{"_id": objId}).Decode(&user)

	if err != nil {
		return user, errorCommon.NewInvariantError("User not found")
	}

	return user, nil
}

func (r *UserRepositoryImplementation) FindByEmail(ctx context.Context, email string) (user userDomain.User, err error) {
	err = r.db.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err != nil {
		return user, errorCommon.NewInvariantError("User not found")
	}

	return user, nil
}

func (r *UserRepositoryImplementation) FindAllPsychologists(ctx context.Context) (users []userDomain.User, err error) {
	cursor, err := r.db.Collection("users").Find(ctx, bson.M{"role": "psychologist"})

	if err != nil {
		return users, errorCommon.NewInternalServerError("Failed to fetch psychologists")
	}

	defer cursor.Close(ctx)

	err = cursor.All(ctx, &users)

	if err != nil {
		return users, errorCommon.NewInternalServerError("Failed to fetch psychologists")
	}

	return users, nil
}

func (r *UserRepositoryImplementation) Update(ctx context.Context, user userDomain.User) (err error) {
	user.UpdatedAt = time.Now()

	_, err = r.db.Collection("users").UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})

	if err != nil {
		return errorCommon.NewInvariantError("Failed to update user")
	}

	return nil
}

func (r *UserRepositoryImplementation) UpdateVerifiedEmail(ctx context.Context, id string) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return errorCommon.NewInvariantError("Invalid user id format")
	}

	_, err = r.db.Collection("users").UpdateByID(ctx, objId, bson.M{"$set": bson.M{"is_email_verified": true, "updated_at": time.Now()}})

	if err != nil {
		return errorCommon.NewInternalServerError("Failed to update user")
	}

	return nil
}

func (r *UserRepositoryImplementation) UpdatePassword(ctx context.Context, id, password string) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return errorCommon.NewInvariantError("Invalid user id format")
	}

	_, err = r.db.Collection("users").UpdateByID(ctx, objId, bson.M{"$set": bson.M{"password": password, "updated_at": time.Now()}})

	if err != nil {
		return errorCommon.NewInvariantError("Failed to update user")
	}

	return nil
}
