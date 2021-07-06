package registry

import (
	"artics-api/src/config"
	"artics-api/src/internal/application"
	"artics-api/src/internal/application/command"
	"artics-api/src/internal/application/query"
	"artics-api/src/internal/infrastructure/repository"
	v1 "artics-api/src/internal/interface/handler/v1"
	"artics-api/src/internal/interface/middleware"
)

// Registry - DI container
type Registry struct {
	AuthMiddleware middleware.AuthMiddleware
	V1User         v1.V1UserHandler
	V1Content      v1.V1ContentHandler
	V1Notice       v1.V1NoticeHandler
}

// NewRegistry - imports files in /internal directory
func NewRegistry(
	uploader *config.UploaderConfig,
	auth *config.AuthConfig,
	mail *config.MailConfig,
	db *config.DatabaseConfig,
	rpc *config.RPCConfig,
	websocket *config.WebsocketConfig,
) *Registry {
	// Domain Repository
	ur := repository.NewUserRepository(db, auth, uploader)
	fr := repository.NewFollowRepository(db)
	cr := repository.NewContentRepository(db)
	cmr := repository.NewCommentRepository(db)
	fvr := repository.NewFavoriteRepository(db)
	br := repository.NewBrowseRepository(db)
	nr := repository.NewNoticeRepository(db)

	// Commands and Queries
	ua := application.UserApplication{
		Commands: application.UserCommands{
			CreateUser:  command.NewCreateUserHandler(ur),
			UpdateUser:  command.NewUpdateUserHandler(ur),
			SuspendUser: command.NewSuspendUserHandler(ur),
			Follow:      command.NewFollowHandler(fr, nr),
			Unfollow:    command.NewUnfollowHandler(fr),
		},
		Queries: application.UserQueries{
			UserByToken: query.NewUserByTokenHandler(ur),
			GetUser:     query.NewGetUserHandler(ur),
			Followings:  query.NewFollowingsHandler(ur),
			Followers:   query.NewFollowersHandler(ur),
		},
	}

	ca := application.ContentApplication{
		Commands: application.ContentCommands{
			Like:   command.NewLikeHandler(cr, fvr, cmr, nr, ur),
			Unlike: command.NewUnlikeHandler(fvr),
			Browse: command.NewBrowseHandler(br),
		},
		Queries: application.ContentQueries{
			Content:             query.NewContentHandler(cr),
			GetFavoriteContents: query.NewGetFavoriteContentsHandler(cr),
		},
	}

	na := application.NoticeApplication{
		Queries: application.NoticeQueries{
			List: query.NewNoticesHandler(nr),
		},
	}

	return &Registry{
		AuthMiddleware: middleware.NewAuthMiddleware(ua),
		V1User:         v1.NewV1UserHandler(ua, websocket),
		V1Content:      v1.NewV1ContentHandler(ca, websocket),
		V1Notice:       v1.NewV1NoticeHandler(na, websocket),
	}
}
