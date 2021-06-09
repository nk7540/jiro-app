package pkg

import (
	"artics-api/src/config"
	"artics-api/src/internal/domain/user"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/xerrors"
)

type emptyCtx int

type Context struct {
	*fiber.Ctx
}

func (Context) Deadline() (deadline time.Time, ok bool) {
	return
}

func (Context) Done() <-chan struct{} {
	return nil
}

func (Context) Err() error {
	return nil
}

func (Context) Value(key interface{}) interface{} {
	return nil
}

func (c *Context) CurrentUser() (*user.User, error) {
	u, ok := c.Locals("user").(*user.User)
	if !ok {
		return nil, xerrors.New("failed to get user from fiber context")
	}

	return u, nil
}

func (c *Context) Printer() (*config.I18nConfig, error) {
	p, ok := c.Locals("i18n").(*config.I18nConfig)
	if !ok {
		return nil, xerrors.New("failed to get printer from fiber context")
	}

	return p, nil
}
