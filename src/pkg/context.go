package pkg

import (
	"time"

	"github.com/gofiber/fiber/v2"
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
