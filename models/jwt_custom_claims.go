package models

import "github.com/golang-jwt/jwt/v5"

type JWTCustomClaims struct {
	Id       string
	Role     string
	Fullname string
	Active   bool
	jwt.RegisteredClaims
}
