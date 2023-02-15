package user

import (
	"context"

	httpCommon "github.com/aziemp66/freya-be/common/http"
	"github.com/aziemp66/freya-be/common/jwt"
	"github.com/aziemp66/freya-be/common/password"
	UserDomain "github.com/aziemp66/freya-be/internal/domain/user"
	UserRepository "github.com/aziemp66/freya-be/internal/repository/user"
	"gopkg.in/gomail.v2"
)

type UserUsecaseImplementation struct {
	userRepository  UserRepository.Repository
	passwordManager password.PasswordHashManager
	jwtManager      jwt.JWTManager
	mailDialer      *gomail.Dialer
}

func NewUserUsecaseImplementation(userRepository UserRepository.Repository, passwordManager password.PasswordHashManager, jwtManager jwt.JWTManager, mailDialer *gomail.Dialer) *UserUsecaseImplementation {
	return &UserUsecaseImplementation{userRepository, passwordManager, jwtManager, mailDialer}
}

func (u *UserUsecaseImplementation) Register(ctx context.Context, user httpCommon.AddUser) (err error) {
	user.Password, err = u.passwordManager.HashPassword(user.Password)

	if err != nil {
		return err
	}

	err = u.userRepository.Insert(ctx, UserDomain.User{
		Email:           user.Email,
		Password:        user.Password,
		IsEmailVerified: false,
	})

	if err != nil {
		return err
	}

	err = u.sendMailActivation(ctx, user.Email)

	if err != nil {
		return err
	}

	return nil
}

// send mail activation
func (u *UserUsecaseImplementation) sendMailActivation(ctx context.Context, email string) (err error) {
	if err != nil {
		panic(err)
	}
	return
}
