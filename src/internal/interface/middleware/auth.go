package middleware

import (
	"artics-api/src/internal/application"
	"artics-api/src/pkg"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware interface {
	Auth(*fiber.Ctx) error
}

type authMiddleware struct {
	ua application.UserApplication
}

func NewAuthMiddleware(ua application.UserApplication) AuthMiddleware {
	return &authMiddleware{ua}
}

func (am *authMiddleware) Auth(c *fiber.Ctx) error {
	a := c.Get("Authorization")
	token := strings.Replace(a, "Bearer ", "", 1)
	u, err := am.ua.Queries.UserByToken.Handle(pkg.Context{Ctx: c}, token)
	if err != nil {
		return err
	}
	c.Locals("user", u)
	return c.Next()
}
