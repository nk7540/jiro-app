package config

import (
	"artics-api/src/internal/domain/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type WebsocketConn = websocket.Conn

type WebsocketConfig struct {
	Conns map[user.UserID]*WebsocketConn
}

func (c *WebsocketConfig) New(handler func(*WebsocketConn)) fiber.Handler {
	return websocket.New(handler) // @TODO Add config
}
