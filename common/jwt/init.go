package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTManager struct {
	AccessTokenKey []byte
	AdminTokenKey  []byte
}

func NewJWTManager(accessTokenKey string) *JWTManager {
	return &JWTManager{AccessTokenKey: []byte(accessTokenKey)}
}

func (j JWTManager) GenerateAuthToken(id string, name string, role string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, AuthClaims{
		ID:   id,
		Name: name,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	})

	tokenString, err := token.SignedString(j.AccessTokenKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j JWTManager) VerifyAuthToken(tokenString string) (id, name, role string, err error) {
	claims := &AuthClaims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return j.AccessTokenKey, nil
	})

	if err != nil {
		return
	}

	if !tkn.Valid {
		err = errors.New("token invalid")
		return
	}

	return claims.ID, claims.Name, claims.Role, nil
}

func (j JWTManager) GenerateUserToken(id string, password string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	})

	stringAccessToken := string(j.AccessTokenKey)

	userAccessToken := []byte(stringAccessToken + password)

	tokenString, err := token.SignedString(userAccessToken)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j JWTManager) VerifyUserToken(tokenString string, password string) (err error) {
	claims := &UserClaims{}

	stringAccessToken := string(j.AccessTokenKey)

	userAccessToken := []byte(stringAccessToken + password)

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return userAccessToken, nil
	})

	if err != nil {
		return err
	}

	if !tkn.Valid {
		err = errors.New("token invalid")
		return err
	}

	return nil
}
