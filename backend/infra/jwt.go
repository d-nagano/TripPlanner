package infra

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	UserID string `json:"userId"`
	jwt.RegisteredClaims
}
