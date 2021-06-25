package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Conn = websocket.Conn

type WebsocketConfig struct{}

func (c *WebsocketConfig) New(handler func(*Conn)) fiber.Handler {
	return websocket.New(handler) // @TODO Add config
}
