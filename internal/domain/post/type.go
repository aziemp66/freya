package post

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Post struct {
		ID        primitive.ObjectID `bson:"_id,omitempty"`
		Title     string             `bson:"title,omitempty"`
		Content   string             `bson:"content,omitempty"`
		AuthorId  primitive.ObjectID `bson:"author_id,omitempty"`
		CreatedAt time.Time          `bson:"created_at,omitempty"`
		UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	}
	Comment struct {
		ID        primitive.ObjectID `bson:"_id,omitempty"`
		Content   string             `bson:"content,omitempty"`
		AuthorId  primitive.ObjectID `bson:"author_id,omitempty"`
		PostId    primitive.ObjectID `bson:"post_id,omitempty"`
		CreatedAt time.Time          `bson:"created_at,omitempty"`
		UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	}
)
