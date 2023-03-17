package post

import (
	"context"

	postDomain "github.com/aziemp66/freya-be/internal/domain/post"
)

type Repository interface {
	InsertPost(ctx context.Context, post postDomain.Post) (err error)
	FindPostByID(ctx context.Context, id string) (post postDomain.Post, err error)
	FindAllPosts(ctx context.Context) (posts []postDomain.Post, err error)
	DeletePost(ctx context.Context, id string) (err error)
	InsertComment(ctx context.Context, comment postDomain.Comment) (err error)
	FindAllCommentsByPostID(ctx context.Context, id string) (comments []postDomain.Comment, err error)
	FindCommentByID(ctx context.Context, id string) (comment postDomain.Comment, err error)
	DeleteComment(ctx context.Context, id string) (err error)
}
