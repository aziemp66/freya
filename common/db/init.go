package db

import (
	"context"

	"cloud.google.com/go/firestore"
)

func NewClient(ctx context.Context, projectId string) *firestore.Client {
	client, err := firestore.NewClient(ctx, projectId)

	if err != nil {
		panic(err)
	}

	return client
}
