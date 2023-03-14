package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

func NewDB(connectionString string, databaseName string) *mongo.Database {
	clientOptions := options.Client()
	clientOptions.ApplyURI(connectionString)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Connect(ctx)

	if err != nil {
		panic(err)
	}

	return client.Database(databaseName)
}
