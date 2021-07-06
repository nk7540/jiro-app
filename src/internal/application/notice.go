package application

import (
	"artics-api/src/internal/application/query"
)

type NoticeApplication struct {
	Queries NoticeQueries
}

type NoticeQueries struct {
	List query.NoticesHandler
}
