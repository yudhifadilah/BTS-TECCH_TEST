package models

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UserID uint `json:"userId"`
	jwt.RegisteredClaims
}
