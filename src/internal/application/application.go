package application

import (
	"artics-api/src/internal/application/command"
	"artics-api/src/internal/application/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateUser      command.CreateUserHandler
	UpdateThumbnail command.UpdateThumbnailHandler
	Update          command.UpdateUserHandler
}

type Queries struct {
	GetUser             query.GetUserHandler
	GetFavoriteContents query.GetFavoriteContentsHanlder
	Followings          query.FollowingsHandler
	Followers           query.FollowersHandler
}
