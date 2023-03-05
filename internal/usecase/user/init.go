package user

import (
	"github.com/aziemp66/freya-be/common/jwt"
	"github.com/aziemp66/freya-be/common/password"
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
