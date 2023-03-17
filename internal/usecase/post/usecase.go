package post

import (
	"context"

	httpCommon "github.com/aziemp66/freya-be/common/http"
)

type Usecase interface {
	InsertPost(ctx context.Context, authorId, title, content string) (err error)
	GetPostById(ctx context.Context, id string) (post httpCommon.Post, err error)
	GetAllPost(ctx context.Context) (posts []httpCommon.Post, err error)
	DeletePost(ctx context.Context, id string) (err error)
	InsertComment(ctx context.Context, authorId, postId, content string) (err error)
	GetAllCommentByPostId(ctx context.Context, postId string) (comments []httpCommon.Comment, err error)
	DeleteComment(ctx context.Context, id string) (err error)
}
