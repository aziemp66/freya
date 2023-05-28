package jwt

import "github.com/golang-jwt/jwt/v4"

type AuthClaims struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}
