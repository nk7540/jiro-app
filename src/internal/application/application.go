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
	// User
	CreateUser      command.CreateUserHandler
	UpdateThumbnail command.UpdateThumbnailHandler
	Update          command.UpdateUserHandler
	Suspend         command.SuspendUserHandler
	Follow          command.FollowHandler
	Unfollow        command.UnfollowHandler

	// Content
}

type Queries struct {
	// User
	GetUser    query.GetUserHandler
	Followings query.FollowingsHandler
	Followers  query.FollowersHandler

	// Content
	GetFavoriteContents query.GetFavoriteContentsHanlder
}
