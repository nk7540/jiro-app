package application

import (
	"artics-api/src/internal/application/command"
	"artics-api/src/internal/application/query"
)

type UserApplication struct {
	Commands UserCommands
	Queries  UserQueries
}

type UserCommands struct {
	CreateUser      command.CreateUserHandler
	UpdateThumbnail command.UpdateThumbnailHandler
	UpdateUser      command.UpdateUserHandler
	SuspendUser     command.SuspendUserHandler
	Follow          command.FollowHandler
	Unfollow        command.UnfollowHandler
}

type UserQueries struct {
	GetUser    query.GetUserHandler
	Followings query.FollowingsHandler
	Followers  query.FollowersHandler
}
