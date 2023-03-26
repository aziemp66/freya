package user

import (
	"context"

	userDomain "github.com/aziemp66/freya-be/internal/domain/user"
)

type Repository interface {
	Insert(ctx context.Context, user userDomain.User) (err error)
	FindByID(ctx context.Context, id string) (user userDomain.User, err error)
	FindByEmail(ctx context.Context, email string) (user userDomain.User, err error)
	FindAllPsychologists(ctx context.Context) (users []userDomain.User, err error)
	Update(ctx context.Context, user userDomain.User) (err error)
	UpdateVerifiedEmail(ctx context.Context, id string) (err error)
	UpdatePassword(ctx context.Context, id, password string) (err error)
}
