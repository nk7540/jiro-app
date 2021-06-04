package v1

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/usecase"
	"artics-api/src/internal/usecase/response"
	"artics-api/src/pkg"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type V1ContentHandler interface {
	Show(c *fiber.Ctx) error
	Favorites(c *fiber.Ctx) error
}

type v1ContentHandler struct {
	u usecase.ContentUsecase
}

func NewV1ContentHandler(u usecase.ContentUsecase) V1ContentHandler {
	return &v1ContentHandler{u}
}

func (h *v1ContentHandler) Show(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	content, err := h.u.Show(pkg.Context{Ctx: c}, id)
	if err != nil {
		return err
	}

	res := &response.Content{
		ID:    content.ID,
		Title: content.Title,
	}

	return c.JSON(res)
}

func (h *v1ContentHandler) Favorites(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	cs, err := h.u.Favorites(pkg.Context{Ctx: c}, userID)
	if err != nil {
		return err
	}

	resContents := make([]*response.Content, len(cs))
	for i, c := range cs {
		resContents[i] = &response.Content{
			ID:    c.ID,
			Title: c.Title,
		}
	}
	res := &response.Contents{resContents}

	return c.JSON(res)
}
