package model

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	UserID uint   `json:"user_id"`
	Phone  string `json:"phone"`
	Role   string `json:"role"`

	jwt.RegisteredClaims
}
