package user

import (
	"context"
	"time"

	httpCommon "github.com/aziemp66/freya-be/common/http"
)

type Usecase interface {
	Register(ctx context.Context, email, password, firstName, lastName string, birthday time.Time) (err error)
	Login(ctx context.Context, email, password string) (token string, err error)
	ForgotPassword(ctx context.Context, email string) (err error)
	ResetPassword(ctx context.Context, id, token, newPassword string) (err error)
	UpdatePassword(ctx context.Context, id, oldPassword, newPassword string) (err error)
	GetById(ctx context.Context, id string) (user httpCommon.User, err error)
	Update(ctx context.Context, id string, user httpCommon.UpdateUser) (err error)
	SendMailActivation(ctx context.Context, email string) (err error)
	Activate(ctx context.Context, id string) (err error)
}
