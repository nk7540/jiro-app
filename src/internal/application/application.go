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
	UpdateUser      command.UpdateUserHandler
	SuspendUser     command.SuspendUserHandler
	Follow          command.FollowHandler
	Unfollow        command.UnfollowHandler

	// Content
	Like   command.LikeHandler
	Unlike command.UnlikeHandler
}

type Queries struct {
	// User
	GetUser    query.GetUserHandler
	Followings query.FollowingsHandler
	Followers  query.FollowersHandler

	// Content
	GetFavoriteContents query.GetFavoriteContentsHanlder
}
