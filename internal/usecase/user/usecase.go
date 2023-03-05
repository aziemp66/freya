package user

import (
	"context"

	httpCommon "github.com/aziemp66/freya-be/common/http"
)

type Usecase interface {
	Register(ctx context.Context, user httpCommon.AddUser) (err error)
	Login(ctx context.Context, email, password string) (token string, err error)
	ForgotPassword(ctx context.Context, email string) (err error)
	ResetPassword(ctx context.Context, id, oldPassword, newPassword string) (err error)
	GetById(ctx context.Context, id string) (user httpCommon.User, err error)
	Update(ctx context.Context, user httpCommon.UpdateUser) (err error)
	sendMailActivation(ctx context.Context, email string) (err error)
	Activate(ctx context.Context, id string) (err error)
}
