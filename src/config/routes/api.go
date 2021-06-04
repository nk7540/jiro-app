package routes

import (
	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(api fiber.Router) {

}

func v1AuthRoutes(api fiber.Router) {
	api.Post("/users/auth")
}

func v1Routes(api fiber.Router) {
	v1 := api.Group("v1")
	v1AuthRoutes(v1)
}
