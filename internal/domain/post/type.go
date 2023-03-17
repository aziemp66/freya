package post

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Post struct {
		ID       primitive.ObjectID `bson:"_id,omitempty"`
		Title    string             `bson:"title,omitempty"`
		Content  string             `bson:"content,omitempty"`
		AuthorId primitive.ObjectID `bson:"author_id,omitempty"`
	}
	Comment struct {
		ID       primitive.ObjectID `bson:"_id,omitempty"`
		Content  string             `bson:"content,omitempty"`
		AuthorId primitive.ObjectID `bson:"author_id,omitempty"`
		PostId   primitive.ObjectID `bson:"post_id,omitempty"`
	}
)
