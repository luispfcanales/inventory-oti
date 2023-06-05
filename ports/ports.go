package ports

import "github.com/luispfcanales/inventory-oti/models"

type StorageService interface {
	GetUserWithCredentials(email, pwd string) (models.User, error)
}
type AuthService interface {
	AuthUser(email, pwd string) (models.User, error)
	ValidateTokenCookie(token string) bool
}
