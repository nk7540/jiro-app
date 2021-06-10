package v1

import (
	"artics-api/src/internal/application"
	"artics-api/src/internal/application/query"
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/content"
	"artics-api/src/internal/usecase/response"
	"artics-api/src/pkg"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type V1ContentHandler interface {
	Show(c *fiber.Ctx) error
	Favorites(c *fiber.Ctx) error
	Like(c *fiber.Ctx) error
	Unlike(c *fiber.Ctx) error
}

type v1ContentHandler struct {
	app application.ContentApplication
}

func NewV1ContentHandler(app application.ContentApplication) V1ContentHandler {
	return &v1ContentHandler{app}
}

func (h *v1ContentHandler) Show(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	content, err := h.app.Queries.Content.Handle(pkg.Context{Ctx: c}, content.ContentID(id))
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

	cs, err := h.app.Queries.GetFavoriteContents.Handle(pkg.Context{Ctx: c}, query.GetFavoriteContents{
		UserID: userID,
		Limit:  10,
	})
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

func (h *v1ContentHandler) Like(c *fiber.Ctx) error {
	contentID, err := strconv.Atoi(c.Query("content_id"))
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	ctx := pkg.Context{Ctx: c}
	u, err := ctx.CurrentUser()
	if err != nil {
		return err
	}

	if err := h.app.Commands.Like.Handle(ctx, content.CommandLike{
		UserID:    content.FavoriteUserID(u.ID),
		ContentID: content.FavoriteContentID(contentID),
	}); err != nil {
		return err
	}

	return c.JSON(nil)
}

func (h *v1ContentHandler) Unlike(c *fiber.Ctx) error {
	contentID, err := strconv.Atoi(c.Query("content_id"))
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	ctx := pkg.Context{Ctx: c}
	u, err := ctx.CurrentUser()
	if err != nil {
		return err
	}

	if err := h.app.Commands.Unlike.Handle(ctx, content.CommandUnlike{
		UserID:    content.FavoriteUserID(u.ID),
		ContentID: content.FavoriteContentID(contentID),
	}); err != nil {
		return err
	}

	return c.JSON(nil)
}
