package post

import (
	"context"

	postDomain "github.com/aziemp66/freya-be/internal/domain/post"
)

type Repository interface {
	Insert(ctx context.Context, post postDomain.Post) (err error)
	FindByID(ctx context.Context, id string) (post postDomain.Post, err error)
	FindAll(ctx context.Context) (posts []postDomain.Post, err error)
	Delete(ctx context.Context, id string) (err error)
}
