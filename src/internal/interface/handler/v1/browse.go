package v1

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/usecase"
	"artics-api/src/pkg"

	"github.com/gofiber/fiber/v2"
)

type V1BrowseHandler interface {
	Save(c *fiber.Ctx) error
}

type v1BrowseHandler struct {
	u usecase.BrowseUsecase
}

func NewV1BrowseHandler(u usecase.BrowseUsecase) V1BrowseHandler {
	return &v1BrowseHandler{u}
}

func (h *v1BrowseHandler) Save(c *fiber.Ctx) error {
	contentID, err := c.ParamsInt("content_id")
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	if err := h.u.Save(pkg.Context{Ctx: c}, contentID); err != nil {
		return err
	}

	return c.JSON(nil)
}
