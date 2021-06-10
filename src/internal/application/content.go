package application

import (
	"artics-api/src/internal/application/command"
	"artics-api/src/internal/application/query"
)

type ContentApplication struct {
	Commands ContentCommands
	Queries  ContentQueries
}

type ContentCommands struct {
	Like   command.LikeHandler
	Unlike command.UnlikeHandler
}

type ContentQueries struct {
	Content             query.ContentHandler
	GetFavoriteContents query.GetFavoriteContentsHanlder
}
