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

		// Follow
		api.Post("/follows", reg.V1Follow.Create)
		api.Delete("/follows", reg.V1Follow.Delete)

		// Content
		api.Get("/contents/favorites", reg.V1Content.Favorites)
		api.Get("/contents/:id", reg.V1Content.Show)

		// Favorite
		api.Post("/favorites", reg.V1Favorite.Create)
		api.Delete("/favorites", reg.V1Favorite.Delete)

		// Browse
		api.Post("/browses", reg.V1Browse.Save)
	}
}
