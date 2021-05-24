package config

import (
	"github.com/gin-gonic/gin"

	"artics-api/src/registry"
)

func Router(reg *registry.Registry) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		// ユーザー
		api.POST("/users", reg.V1User.Create)
		// api.POST("/users/login", reg.UserHandler.Login)
		api.GET("/users/:id", reg.V1User.Show)
		api.PATCH("/users", reg.V1User.Update)
		// api.DELETE("/users/leave", reg.UserHandler.Leave)

		// // フォロー
		// api.POST("/follows", reg.FollowHandler.CreateFollow)
		// api.GET("/follows", reg.FollowHandler.GetFollows)
		// api.DELETE("/follows", reg.FollowHandler.DestroyFollow)

		// // コンテンツ
		// api.POST("/contents", reg.ContentHandler.CreateContent)
		// api.GET("/contents", reg.ContentHandler.GetContents)
		// api.GET("/contents/:id", reg.ContentHandler.GetContent)
		// api.PATCH("/contents/:id", reg.ContentHandler.UpdateContent)
		// api.DELETE("/contents/:id", reg.ContentHandler.DestroyContent)

		// // お気に入り
		// api.POST("/favorites", reg.FavoriteHandler.CreateFavorite)
		// api.GET("/favorites", reg.FavoriteHandler.GetFavorites)
		// api.DELETE("/favorites", reg.FavoriteHandler.DestroyFavorite)

		// // カテゴリ
		// api.GET("/categories", reg.CategoryHandler.GetCategories)
	}

	return r
}
