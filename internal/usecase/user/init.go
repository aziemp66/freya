package user

import (
	"context"
	"fmt"
	"time"

	errorCommon "github.com/aziemp66/freya-be/common/error"
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
		return errorCommon.NewInternalServerError("Failed to hash password")
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
		return "", errorCommon.NewUnauthorizedError("Password or email is wrong")
	}

	err = u.passwordManager.CheckPasswordHash(password, userData.Password)

	if err != nil {
		return "", errorCommon.NewInvariantError("Password or email is wrong")
	}

	token, err = u.jwtManager.GenerateAuthToken(userData.Email, fmt.Sprintf("%s %s", userData.FirstName, userData.LastName), string(userData.Role), 24*time.Hour)

	if err != nil {
		return "", errorCommon.NewInternalServerError("Failed to generate token")
	}

	return token, err
}

func (u *UserUsecaseImplementation) ForgotPassword(ctx context.Context, email string) (err error) {
	user, err := u.userRepository.FindByEmail(ctx, email)

	if err != nil {
		return err
	}

	token, err := u.jwtManager.GenerateAuthToken(user.ID.Hex(), fmt.Sprintf("%s %s", user.FirstName, user.LastName), string(user.Role), 24*time.Hour)

	if err != nil {
		return errorCommon.NewInternalServerError("Failed to generate token")
	}

	mailTemplate, err := mailCommon.RenderPasswordResetTemplate(token)

	if err != nil {
		return errorCommon.NewInternalServerError(err.Error())
	}

	message := mailCommon.NewMessage(u.mailDialer.Username, user.Email, "Reset Password", mailTemplate)

	err = u.mailDialer.DialAndSend(message)

	if err != nil {
		return errorCommon.NewInternalServerError("Failed to send mail")
	}

	return nil
}

func (u *UserUsecaseImplementation) ResetPassword(ctx context.Context, token, newPassword string) (err error) {
	newPassword, err = u.passwordManager.HashPassword(newPassword)

	if err != nil {
		return errorCommon.NewInternalServerError("Failed to hash password")
	}

	claims, err := u.jwtManager.VerifyAuthToken(token)

	if err != nil {
		return errorCommon.NewInvariantError("Token is invalid")
	}

	err = u.userRepository.UpdatePassword(ctx, claims.ID, newPassword)

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
		return errorCommon.NewInvariantError("Old password is wrong")
	}

	hashedPassword, err := u.passwordManager.HashPassword(newPassword)

	if err != nil {
		return errorCommon.NewInternalServerError("Failed to hash password")
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
		return errorCommon.NewInvariantError("Invalid id format")
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
	user, err := u.userRepository.FindByEmail(ctx, email)

	if err != nil {
		return err
	}

	token, err := u.jwtManager.GenerateAuthToken(user.ID.Hex(), fmt.Sprintf("%s %s", user.FirstName, user.LastName), string(user.Role), 24*time.Hour)

	if err != nil {
		return errorCommon.NewInternalServerError("Failed to generate token")
	}

	templates, err := mailCommon.RenderEmailVerificationTemplate(token)

	if err != nil {
		return errorCommon.NewInternalServerError(err.Error())
	}

	msg := mailCommon.NewMessage(u.mailDialer.Username, user.Email, "Freya - Email Activation", templates)

	u.mailDialer.DialAndSend(msg)

	return
}

func (u *UserUsecaseImplementation) Activate(ctx context.Context, token string) (err error) {
	claims, err := u.jwtManager.VerifyAuthToken(token)

	if err != nil {
		return errorCommon.NewUnauthorizedError("Token is invalid")
	}

	user, err := u.userRepository.FindByID(ctx, claims.ID)

	if err != nil {
		return err
	}

	if user.IsEmailVerified {
		return errorCommon.NewInvariantError("Email is already verified")
	}

	err = u.userRepository.UpdateVerifiedEmail(ctx, user.ID.Hex())

	if err != nil {
		return err
	}

	return
}
