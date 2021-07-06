package v1

import (
	"artics-api/src/config"
	"artics-api/src/internal/application"
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/notice"
	"artics-api/src/internal/domain/user"
	"artics-api/src/internal/interface/handler/request"
	"artics-api/src/internal/interface/handler/response"
	"artics-api/src/pkg"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/xerrors"
)

type V1NoticeHandler interface {
	Listen(c *fiber.Ctx) error
	List(c *fiber.Ctx) error
}

type v1NoticeHandler struct {
	app       application.NoticeApplication
	websocket *config.WebsocketConfig
}

func NewV1NoticeHandler(app application.NoticeApplication, websocket *config.WebsocketConfig) V1NoticeHandler {
	return &v1NoticeHandler{app, websocket}
}

func (h *v1NoticeHandler) Listen(c *fiber.Ctx) error {
	return h.websocket.New(func(conn *config.WebsocketConn) {
		ctx := pkg.Context{Ctx: c}
		u, err := ctx.CurrentUser()
		if err != nil {
			return
		}
		h.websocket.Conns[u.ID] = conn
		for {
			var req request.Listen
			if err = conn.ReadJSON(&req); err != nil {
				conn.Close()
				delete(h.websocket.Conns, u.ID)
				break
			} else {
				conn.Close()
				delete(h.websocket.Conns, u.ID)
				break
			}
		}
	})(c)
}

func (h *v1NoticeHandler) List(c *fiber.Ctx) error {
	req := &request.ListNotices{}
	if err := c.BodyParser(req); err != nil {
		return domain.UnableParseJSON.New(err)
	}

	ctx := pkg.Context{Ctx: c}
	u, err := ctx.CurrentUser()
	if err != nil {
		return err
	}

	qns, err := h.app.Queries.List.Handle(ctx, notice.QueryArgNotices{
		UserID: u.ID,
		Page:   req.Page,
		Per:    req.Per,
	})
	if err != nil {
		return err
	}

	resNotices := make([]interface{}, 0)
	for i, qn := range qns.Notices {
		res, err := noticeResponse(qn)
		if err != nil {
			return err
		}
		resNotices[i] = res
	}
	res := response.Notices{resNotices}

	return c.JSON(res)
}

func notify(ws *config.WebsocketConfig, qn *notice.QueryNotice) error {
	conn := ws.Conns[user.UserID(qn.UserID)]
	if conn != nil {
		res, err := noticeResponse(qn)
		if err != nil {
			return err
		}
		if err := conn.WriteJSON(res); err != nil {
			conn.Close()
			delete(ws.Conns, user.UserID(qn.UserID))
			return err
		}
	}

	return nil
}

func noticeResponse(qn *notice.QueryNotice) (interface{}, error) {
	t := notice.NoticeType(qn.Type)
	switch t {
	case notice.Favorite:
		return &response.NoticeFavorite{
			ID:     qn.ID,
			Type:   t.String(),
			IsRead: qn.IsRead,
			// NoticeFavorite
			FavoriteID:          qn.Favorite.FavoriteID,
			UserID:              qn.Favorite.UserID,
			UserThumbnailURL:    qn.Favorite.UserThumbnailURL,
			Header:              qn.Favorite.Header,
			Body:                qn.Favorite.Body,
			ContentID:           qn.Favorite.ContentID,
			ContentThumbnailURL: qn.Favorite.ContentThumbnailURL,
		}, nil
	case notice.Followed:
		return &response.NoticeFollowed{
			ID:     qn.ID,
			Type:   t.String(),
			IsRead: qn.IsRead,
			// NoticeFollowed
			UserID:           qn.Followed.UserID,
			UserThumbnailURL: qn.Followed.UserThumbnailURL,
			Body:             qn.Followed.Body,
		}, nil
	default:
		return nil, xerrors.New("invalid notice type")
	}
}
