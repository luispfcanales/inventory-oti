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
	u, err := a.repo.SelectUserWithCredentials(username, password)

	if err != nil {
		log.Println("Error Service Auth:", err)
		return models.User{}, userNotFound
	}

	if u.Person.IDPerson == 0 {
		return models.User{}, userNotFound
	}

	t, err := a.GenerateToken(
		u.IDRole,
		fmt.Sprintf("%s %s", u.FirstName, u.LastName),
		u.Active,
	)
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

	t, err := jwt.ParseWithClaims(tokenString, &models.JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Keytoken, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return errors.New("unautorized")
		}
		return errors.New("internal server errror")
	}

	if claims, ok := t.Claims.(*models.JWTCustomClaims); ok && t.Valid {
		log.Println("DETAIL USER AUTH -> ", fmt.Sprintf(
			"[ ROLE: %s, FULLNAME: %s, ACTIVE: %v ]",
			claims.Role,
			claims.Fullname,
			claims.Active,
		))
		return nil
	}

	return fiber.ErrUnauthorized
}

func (a *auth) GenerateToken(role string, fullname string, active bool) (string, error) {
	claims := &models.JWTCustomClaims{
		role,
		fullname,
		active,
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
