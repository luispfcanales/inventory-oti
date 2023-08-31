package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var STORE_SESSION *session.Store
var KEY_SESSION_USERNAME string = "username"

func CreateStoreSession() {
	STORE_SESSION = session.New()

}

func GetInstanceSession(c *fiber.Ctx) *session.Session {
	sess, err := STORE_SESSION.Get(c)
	if err != nil {
		panic(err)
	}
	return sess
}
