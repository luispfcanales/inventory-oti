package ports

import (
	"github.com/luispfcanales/inventory-oti/models"
)

type AuthService interface {
	AuthUser(email, pwd string) (models.User, error)
	ValidateTokenCookie(token string) bool
}

// api service consumer
type ApiService interface {
	GetDataByDni(dni string) (models.PersonServiceDNI, error)
}
