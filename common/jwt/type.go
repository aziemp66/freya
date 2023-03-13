package jwt

import "github.com/golang-jwt/jwt/v4"

type AuthClaims struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type UserClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
