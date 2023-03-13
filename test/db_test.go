package test

import (
	"context"
	"os/user"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	userDomain "github.com/aziemp66/freya-be/internal/domain/user"
	userRepository "github.com/aziemp66/freya-be/internal/repository/user"
)

var ctx = context.Background()

func generateDB() *mongo.Database {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017/?connectTimeoutMS=10000")
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		panic(err)
	}

	err = client.Connect(ctx)

	if err != nil {
		panic(err)
	}

	db := client.Database("test")

	if db == nil {
		panic("db is nil")
	}

	return db
}

func TestDBInsert(t *testing.T) {
	db := generateDB()

	db.Collection("users").InsertOne(ctx, userDomain.User{
		ID:              primitive.NewObjectID(),
		FirstName:       "Aizen",
		LastName:        "Melza",
		Email:           "Aizen@gmail.com",
		Password:        "blaablablba",
		Role:            "user",
		IsEmailVerified: false,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	})
}

func TestDBRead(t *testing.T) {
	db := generateDB()

	user := user.User{}
	objectid, err := primitive.ObjectIDFromHex("64029aee95f63a3a323176dd")

	if err != nil {
		t.Error(err)
	}

	db.Collection("users").FindOne(ctx, bson.M{"_id": objectid}).Decode(&user)

	t.Log(user)
}

func TestDBUpdate(t *testing.T) {
	db := generateDB()

	userRepository := userRepository.NewUserRepositoryImplementation(db)

	objectid, err := primitive.ObjectIDFromHex("6402cfe8d51715ec946fe123")

	if err != nil {
		t.Error(err)
	}

	user := userDomain.User{
		ID:       objectid,
		LastName: "Melza",
	}

	ctx := context.Background()
	err = userRepository.Update(ctx, user)

	if err != nil {
		t.Error(err)
	}
}

func TestDBReplace(t *testing.T) {
	db := generateDB()

	objectid, err := primitive.ObjectIDFromHex("64029fa1871bcac885cc8c58")

	if err != nil {
		t.Error(err)
	}

	user := userDomain.User{
		ID:       objectid,
		LastName: "Melza",
	}

	result, err := db.Collection("users").ReplaceOne(ctx, bson.M{"_id": objectid}, user)

	if err != nil {
		t.Error(err)
	}

	t.Log(result)
}

func TestDBDelete(t *testing.T) {
	db := generateDB()

	objectid, err := primitive.ObjectIDFromHex("64029aee95f63a3a323176dd")

	if err != nil {
		t.Error(err)
	}

	result, err := db.Collection("users").DeleteOne(ctx, bson.M{"_id": objectid})

	if err != nil {
		t.Error(err)
	}

	t.Log(result)
}
