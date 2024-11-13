package infra

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	UserID int `json:"userId"`
	jwt.RegisteredClaims
}
