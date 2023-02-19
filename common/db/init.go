package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

func NewDB(connectionString string, databaseName string) (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI(connectionString)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)

	if err != nil {
		return nil, err
	}

	return client.Database(databaseName), nil
}
