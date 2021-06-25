package v1

import (
	"artics-api/src/config"
	"artics-api/src/internal/application"
	"artics-api/src/internal/application/query"
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/content"
	"artics-api/src/internal/interface/handler/response"
	"artics-api/src/pkg"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type V1ContentHandler interface {
	Show(c *fiber.Ctx) error
	Favorites(c *fiber.Ctx) error
	Like(c *fiber.Ctx) error
	Unlike(c *fiber.Ctx) error
	Browse(c *fiber.Ctx) error
	ListenFavorite(c *fiber.Ctx) error
}

type v1ContentHandler struct {
	app       application.ContentApplication
	websocket *config.WebsocketConfig
}

func NewV1ContentHandler(app application.ContentApplication, websocket *config.WebsocketConfig) V1ContentHandler {
	return &v1ContentHandler{app, websocket}
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

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	cs, err := h.app.Queries.GetFavoriteContents.Handle(pkg.Context{Ctx: c}, query.GetFavoriteContents{
		UserID: userID,
		Limit:  limit,
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
	res := &response.Contents{Contents: resContents}

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

func (h *v1ContentHandler) Browse(c *fiber.Ctx) error {
	contentID, err := c.ParamsInt("content_id")
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	ctx := pkg.Context{Ctx: c}
	u, err := ctx.CurrentUser()
	if err != nil {
		return err
	}

	if err := h.app.Commands.Browse.Handle(ctx, content.CommandBrowse{
		UserID:    content.BrowseUserID(u.ID),
		ContentID: content.BrowseContentID(contentID),
	}); err != nil {
		return err
	}

	return c.JSON(nil)
}

func (h *v1ContentHandler) ListenFavorite(c *fiber.Ctx) error {
	return h.websocket.New(func(conn *websocket.Conn) {
		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = conn.ReadMessage(); err != nil {
				break
			}
			log.Printf("recv: %s", msg)

			if err = conn.WriteMessage(mt, msg); err != nil {
				break
			}
		}
	})(c)
}
