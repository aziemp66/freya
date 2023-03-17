package user

import (
	"context"
	"fmt"
	"time"

	httpCommon "github.com/aziemp66/freya-be/common/http"
	"github.com/aziemp66/freya-be/common/jwt"
	mailCommon "github.com/aziemp66/freya-be/common/mail"
	"github.com/aziemp66/freya-be/common/password"
	UserDomain "github.com/aziemp66/freya-be/internal/domain/user"
	UserRepository "github.com/aziemp66/freya-be/internal/repository/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/gomail.v2"
)

type UserUsecaseImplementation struct {
	userRepository  UserRepository.Repository
	passwordManager *password.PasswordHashManager
	jwtManager      *jwt.JWTManager
	mailDialer      *gomail.Dialer
}

func NewUserUsecaseImplementation(userRepository UserRepository.Repository, passwordManager *password.PasswordHashManager, jwtManager *jwt.JWTManager, mailDialer *gomail.Dialer) *UserUsecaseImplementation {
	return &UserUsecaseImplementation{userRepository, passwordManager, jwtManager, mailDialer}
}

func (u *UserUsecaseImplementation) Register(ctx context.Context, email, password, firstName, lastName string, birthDay time.Time) (err error) {
	hashedPassword, err := u.passwordManager.HashPassword(password)

	if err != nil {
		return err
	}

	err = u.userRepository.Insert(ctx, UserDomain.User{
		Email:           email,
		Password:        hashedPassword,
		FirstName:       firstName,
		LastName:        lastName,
		BirthDay:        birthDay,
		IsEmailVerified: false,
	})

	if err != nil {
		return err
	}

	err = u.SendMailActivation(ctx, email)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecaseImplementation) Login(ctx context.Context, email, password string) (token string, err error) {
	userData, err := u.userRepository.FindByEmail(ctx, email)

	if err != nil {
		return "", err
	}

	err = u.passwordManager.CheckPasswordHash(password, userData.Password)

	if err != nil {
		return "", err
	}

	token, err = u.jwtManager.GenerateAuthToken(userData.Email, fmt.Sprintf("%s %s", userData.FirstName, userData.LastName), string(userData.Role), 24*time.Hour)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserUsecaseImplementation) ForgotPassword(ctx context.Context, email string) (err error) {
	userData, err := u.userRepository.FindByEmail(ctx, email)

	if err != nil {
		return err
	}

	token, err := u.jwtManager.GenerateUserToken(userData.ID.Hex(), userData.Password, 24*time.Hour)

	if err != nil {
		return err
	}

	mailPasswordReset := mailCommon.PasswordReset{
		Email: userData.Email,
		Token: token,
	}

	mailTemplate, err := mailCommon.RenderPasswordResetTemplate(mailPasswordReset)

	if err != nil {
		return err
	}

	message := mailCommon.NewMessage(u.mailDialer.Username, userData.Email, "Reset Password", mailTemplate)

	err = u.mailDialer.DialAndSend(message)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecaseImplementation) ResetPassword(ctx context.Context, id, token, newPassword string) (err error) {
	newPassword, err = u.passwordManager.HashPassword(newPassword)

	if err != nil {
		return err
	}

	user, err := u.userRepository.FindByID(ctx, id)

	if err != nil {
		return err
	}

	err = u.jwtManager.VerifyUserToken(token, user.Password)

	if err != nil {
		return err
	}

	err = u.userRepository.UpdatePassword(ctx, id, newPassword)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecaseImplementation) UpdatePassword(ctx context.Context, id, oldPassword, newPassword string) (err error) {
	userData, err := u.userRepository.FindByID(ctx, id)

	if err != nil {
		return err
	}

	err = u.passwordManager.CheckPasswordHash(oldPassword, userData.Password)

	if err != nil {
		return err
	}

	hashedPassword, err := u.passwordManager.HashPassword(newPassword)

	if err != nil {
		return err
	}

	err = u.userRepository.UpdatePassword(ctx, id, hashedPassword)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecaseImplementation) GetById(ctx context.Context, id string) (user httpCommon.User, err error) {
	userData, err := u.userRepository.FindByID(ctx, id)

	if err != nil {
		return httpCommon.User{}, err
	}

	user = httpCommon.User{
		Id:              userData.ID.Hex(),
		FirstName:       userData.FirstName,
		LastName:        userData.LastName,
		Email:           userData.Email,
		BirthDay:        userData.BirthDay,
		Role:            string(userData.Role),
		IsEmailVerified: userData.IsEmailVerified,
		CreatedAt:       userData.CreatedAt,
		UpdatedAt:       userData.UpdatedAt,
	}

	return
}

func (u *UserUsecaseImplementation) Update(ctx context.Context, id string, user httpCommon.UpdateUser) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	err = u.userRepository.Update(ctx, UserDomain.User{
		ID:        objId,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		BirthDay:  user.BirthDay,
	})

	if err != nil {
		return err
	}

	return nil
}

// send mail activation
func (u *UserUsecaseImplementation) SendMailActivation(ctx context.Context, email string) (err error) {
	if err != nil {
		panic(err)
	}
	return
}

func (u *UserUsecaseImplementation) Activate(ctx context.Context, id string) (err error) {
	return
}
