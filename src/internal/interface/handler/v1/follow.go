package v1

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"artics-api/src/internal/domain"
	"artics-api/src/internal/usecase"
	"artics-api/src/pkg"
)

// V1FollowHandler - v1 follow handler
type V1FollowHandler interface {
	Create(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type v1FollowHandler struct {
	u usecase.FollowUsecase
}

// NewV1FollowHandler - setups v1 follow handler
func NewV1FollowHandler(u usecase.FollowUsecase) V1FollowHandler {
	return &v1FollowHandler{u}
}

func (h *v1FollowHandler) Create(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	if err := h.u.Create(pkg.Context{Ctx: c}, id); err != nil {
		return err
	}

	return c.JSON(nil)
}

func (h *v1FollowHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	if err := h.u.Delete(pkg.Context{Ctx: c}, id); err != nil {
		return err
	}

	return c.JSON(nil)
}
