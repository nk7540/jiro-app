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
	CreateUser  command.CreateUserHandler
	UpdateUser  command.UpdateUserHandler
	SuspendUser command.SuspendUserHandler
	Follow      command.FollowHandler
	Unfollow    command.UnfollowHandler
}

type UserQueries struct {
	UserByToken query.UserByTokenHandler
	GetUser     query.GetUserHandler
	Followings  query.FollowingsHandler
	Followers   query.FollowersHandler
}
