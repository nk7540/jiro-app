package middleware

import (
	"artics-api/src/internal/usecase"
	"artics-api/src/pkg"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware interface {
	Auth(*fiber.Ctx) error
}

type authMiddleware struct {
	uu usecase.UserUsecase
}

func NewAuthMiddleware(uu usecase.UserUsecase) AuthMiddleware {
	return &authMiddleware{uu}
}

func (am *authMiddleware) Auth(c *fiber.Ctx) error {
	a := c.Get("Authorization")
	token := strings.Replace(a, "Bearer ", "", 1)
	u, err := am.uu.Auth(pkg.Context{Ctx: c}, token)
	if err != nil {
		return err
	}
	c.Locals("user", u)
	return c.Next()
}
