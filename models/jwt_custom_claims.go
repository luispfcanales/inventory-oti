package models

import "github.com/golang-jwt/jwt/v5"

type JWTCustomClaims struct {
	Role     int
	Fullname string
	Active   bool
	jwt.RegisteredClaims
}
