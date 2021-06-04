package v1

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/usecase"
	"artics-api/src/pkg"

	"github.com/gofiber/fiber/v2"
)

type V1FavoriteHandler interface {
	Create(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type v1FavoriteHandler struct {
	u usecase.FavoriteUsecase
}

func NewV1FavoriteHandler(u usecase.FavoriteUsecase) V1FavoriteHandler {
	return &v1FavoriteHandler{u}
}

func (h *v1FavoriteHandler) Create(c *fiber.Ctx) error {
	contentID, err := c.ParamsInt("content_id")
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	if err := h.u.Create(pkg.Context{Ctx: c}, contentID); err != nil {
		return err
	}

	return c.JSON(nil)
}

func (h *v1FavoriteHandler) Delete(c *fiber.Ctx) error {
	contentID, err := c.ParamsInt("content_id")
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	if err := h.u.Delete(pkg.Context{Ctx: c}, contentID); err != nil {
		return err
	}

	return c.JSON(nil)
}
