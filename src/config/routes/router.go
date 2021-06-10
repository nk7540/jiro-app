package routes

import (
	"github.com/gofiber/fiber/v2"

	"artics-api/src/registry"
)

func Router(app *fiber.App, reg *registry.Registry) {
	app.Post("/v1/users", reg.V1User.Create)

	// Authenticated Routes
	api := app.Group("/v1").Use(reg.AuthMiddleware.Auth)
	{
		// User
		api.Get("/users/:id", reg.V1User.Show)
		api.Patch("/users", reg.V1User.Update)
		api.Delete("/users/suspend", reg.V1User.Suspend)
		api.Get("/users/:id/followings", reg.V1User.Followings)
		api.Get("/users/:id/followers", reg.V1User.Followers)
		api.Post("/follows", reg.V1User.Follow)
		api.Delete("/follows", reg.V1User.Unfollow)

		// Content
		api.Post("/favorites", reg.V1User.Like)
		api.Delete("/favorites", reg.V1User.Unlike)
		api.Post("/browses", reg.V1Browse.Save)
		api.Get("/contents/favorites", reg.V1Content.Favorites)
		api.Get("/contents/:id", reg.V1Content.Show)
	}
}
