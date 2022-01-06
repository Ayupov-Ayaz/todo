package auth

import "github.com/dgrijalva/jwt-go"

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}
