package mock_pkg

import (
	"artics-api/src/config"
	"artics-api/src/internal/domain/user"
	"artics-api/src/pkg"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func NewMockContext(u *user.User) pkg.Context {
	c := &fiber.Ctx{}

	if u != nil {
		c.Locals("user", u)
	}

	tag := language.MustParse("ja")
	i18n := config.I18nConfig{Printer: message.NewPrinter(tag)}
	c.Locals("i18n", i18n)

	return pkg.Context{Ctx: c}
}
