package v1

import (
	"artics-api/src/config"
	"artics-api/src/internal/domain/content"
	"artics-api/src/internal/interface/handler/request"
	"artics-api/src/internal/interface/handler/response"
	"artics-api/src/pkg"

	"github.com/gofiber/fiber/v2"
)

type V1NoticeHandler interface {
	Listen(c *fiber.Ctx) error
}

type v1NoticeHandler struct {
	websocket *config.WebsocketConfig
}

func NewV1NoticeHandler(websocket *config.WebsocketConfig) V1NoticeHandler {
	return &v1NoticeHandler{websocket}
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

func notify(ws *config.WebsocketConfig, n *content.Notice) error {
	conn := ws.Conns[n.UserID]
	if conn != nil {
		res := &response.Notice{
			ID:    int(n.ID),
			Title: string(n.Title),
			Body:  string(n.Body),
		}
		if err := conn.WriteJSON(res); err != nil {
			conn.Close()
			delete(ws.Conns, n.UserID)
			return err
		}
	}

	return nil
}
