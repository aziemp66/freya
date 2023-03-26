package test

import (
	"context"
	"os/user"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	userDomain "github.com/aziemp66/freya-be/internal/domain/user"
	// userRepository "github.com/aziemp66/freya-be/internal/repository/user"

	chatDomain "github.com/aziemp66/freya-be/internal/domain/chat"
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

	// db.Collection("chatrooms").InsertOne(ctx, chatDomain.Chatroom{
	// 	ID:             primitive.NewObjectID(),
	// 	UserID:         primitive.NewObjectID(),
	// 	PsychologistID: primitive.NewObjectID(),
	// 	Messages:       []chatDomain.Message{},
	// 	CreatedAt:      time.Now(),
	// 	UpdatedAt:      time.Now(),
	// })

	// objId, err := primitive.ObjectIDFromHex("641fcfe51847d9be89237f58")

	// if err != nil {
	// 	t.Error(err)
	// }

	// db.Collection(("chatrooms")).UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$push": bson.M{"messages": chatDomain.Message{
	// 	ID:        primitive.NewObjectID(),
	// 	SenderID:  primitive.NewObjectID(),
	// 	Content:   "Ngentot",
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// }}})

	var messageReturn chatDomain.Message

	objId, err := primitive.ObjectIDFromHex("641fd102eb63b8837d78ffa3")

	if err != nil {
		t.Error(err)
	}

	var chatroom chatDomain.Chatroom

	db.Collection("chatrooms").FindOne(
		context.TODO(),
		bson.D{{Key: "messages", Value: bson.D{{Key: "$elemMatch", Value: bson.D{{Key: "_id", Value: objId}}}}}}).Decode(&chatroom)

	for _, message := range chatroom.Messages {
		if message.ID == objId {
			messageReturn = message
		}
	}

	t.Log(messageReturn)
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

	// userRepository := userRepository.NewUserRepositoryImplementation(db)

	// objectid, err := primitive.ObjectIDFromHex("6402cfe8d51715ec946fe123")

	// if err != nil {
	// 	t.Error(err)
	// }

	// user := userDomain.User{
	// 	ID:       objectid,
	// 	LastName: "Melza",
	// }

	// ctx := context.Background()
	// err = userRepository.Update(ctx, user)

	// if err != nil {
	// 	t.Error(err)
	// }

	objId, err := primitive.ObjectIDFromHex("641fd102eb63b8837d78ffa3")

	if err != nil {
		t.Error(err)
	}

	_, err = db.Collection("chatrooms").UpdateOne(ctx, bson.D{{Key: "messages", Value: bson.D{{Key: "$elemMatch", Value: bson.D{{Key: "_id", Value: objId}}}}}}, bson.M{"$pull": bson.M{"messages": bson.M{"_id": objId}}})

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
