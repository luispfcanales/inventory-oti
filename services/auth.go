package services

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

// Keytoken is the secret key that signs the token
var Keytoken = []byte("luiskey")

type JWTCustomClaims struct {
	IDUSER   string
	Fullname string
	jwt.RegisteredClaims
}

type auth struct {
	repo ports.StorageUserService
}

// NewAuth return instance of auth service
func NewAuth(r ports.StorageUserService) *auth {
	return &auth{
		repo: r,
	}
}

func (a *auth) AuthUser(username, password string) (models.User, error) {
	var userNotFound error = errors.New("user not found")
	u, err := a.repo.GetUserWithCredentials(username, password)

	if err != nil {
		log.Println("Error Service Auth:", err)
		return models.User{}, userNotFound
	}

	if u.Key == "" {
		return models.User{}, userNotFound
	}

	t, err := a.GenerateToken(u.Key, fmt.Sprintf("%s %s", u.FirstName, u.LastName))
	if err != nil {
		return models.User{}, err
	}
	u.AccessToken = t
	return u, nil
}

func (a *auth) ValidateTokenCookie(token string) bool {
	err := a.parseToken(token)
	if err != nil {
		log.Println("Error Service Auth -> fn validate token:", err)
		return false
	}
	return true
}
func (a *auth) parseToken(tokenString string) error {

	t, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Keytoken, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return errors.New("unautorized")
		}
		return errors.New("internal server errror")
	}

	if claims, ok := t.Claims.(*JWTCustomClaims); ok && t.Valid {
		log.Println("ID USER AUTH -> ", claims.IDUSER)
		return nil
	}

	return fiber.ErrUnauthorized
}

func (a *auth) GenerateToken(id string, fullname string) (string, error) {
	claims := &JWTCustomClaims{
		id,
		fullname,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(Keytoken)
	if err != nil {
		return "", err
	}
	return t, nil
}
