package post

import (
	"context"
	"time"

	postDomain "github.com/aziemp66/freya-be/internal/domain/post"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepositoryImplementation struct {
	db *mongo.Database
}

func NewPostRepositoryImplementation(db *mongo.Database) *PostRepositoryImplementation {
	return &PostRepositoryImplementation{db}
}

func (r *PostRepositoryImplementation) InsertPost(ctx context.Context, post postDomain.Post) (err error) {
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	_, err = r.db.Collection("posts").InsertOne(ctx, post)

	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepositoryImplementation) FindPostByID(ctx context.Context, id string) (post postDomain.Post, err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return postDomain.Post{}, err
	}

	err = r.db.Collection("posts").FindOne(ctx, postDomain.Post{ID: objId}).Decode(&post)

	if err != nil {
		return postDomain.Post{}, err
	}

	return post, nil
}

func (r *PostRepositoryImplementation) FindAllPosts(ctx context.Context) (posts []postDomain.Post, err error) {
	cursor, err := r.db.Collection("posts").Find(ctx, postDomain.Post{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	err = cursor.All(ctx, &posts)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostRepositoryImplementation) DeletePost(ctx context.Context, id string) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = r.db.Collection("posts").DeleteOne(ctx, postDomain.Post{ID: objId})

	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepositoryImplementation) InsertComment(ctx context.Context, comment postDomain.Comment) (err error) {
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()

	_, err = r.db.Collection("comments").InsertOne(ctx, comment)

	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepositoryImplementation) FindAllCommentsByPostID(ctx context.Context, id string) (comments []postDomain.Comment, err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	cursor, err := r.db.Collection("comments").Find(ctx, postDomain.Comment{PostId: objId})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var comment postDomain.Comment

		err = cursor.Decode(&comment)

		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func (r *PostRepositoryImplementation) FindCommentByID(ctx context.Context, id string) (comment postDomain.Comment, err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return postDomain.Comment{}, err
	}

	err = r.db.Collection("comments").FindOne(ctx, postDomain.Comment{ID: objId}).Decode(&comment)

	if err != nil {
		return postDomain.Comment{}, err
	}

	return comment, nil
}

func (r *PostRepositoryImplementation) DeleteComment(ctx context.Context, id string) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = r.db.Collection("comments").DeleteOne(ctx, postDomain.Comment{ID: objId})

	if err != nil {
		return err
	}

	return nil
}
