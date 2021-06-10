package registry

import (
	"artics-api/src/config"
	"artics-api/src/internal/application"
	"artics-api/src/internal/application/command"
	"artics-api/src/internal/application/query"
	"artics-api/src/internal/infrastructure/repository"
	"artics-api/src/internal/infrastructure/service"
	dv "artics-api/src/internal/infrastructure/validation"
	v1 "artics-api/src/internal/interface/handler/v1"
	"artics-api/src/internal/interface/middleware"
)

// Registry - DI container
type Registry struct {
	AuthMiddleware middleware.AuthMiddleware
	V1User         v1.V1UserHandler
	V1Content      v1.V1ContentHandler
}

// NewRegistry - imports files in /internal directory
func NewRegistry(
	uploader *config.UploaderConfig,
	auth *config.AuthConfig,
	mail *config.MailConfig,
	db *config.DatabaseConfig,
	rpc *config.RPCConfig,
) *Registry {
	// Domain Repository
	ur := repository.NewUserRepository(db, auth)
	fr := repository.NewFollowRepository(db)
	cr := repository.NewContentRepository(db)
	fvr := repository.NewFavoriteRepository(db)
	br := repository.NewBrowseRepository(db)
	flr := repository.NewFileRepository(uploader)

	// Domain Validator
	udv := dv.NewUserDomainValidator(ur)

	// Domain Service
	us := service.NewUserService(udv, ur, fr, cr, flr)
	cs := service.NewContentService(cr)
	fvs := service.NewFavoriteService(fvr)
	bs := service.NewBrowseService(br)

	// Commands and Queries
	ua := application.UserApplication{
		Commands: application.UserCommands{
			CreateUser:      command.NewCreateUserHandler(ur),
			UpdateThumbnail: command.NewUpdateThumbnailHandler(ur),
			UpdateUser:      command.NewUpdateUserHandler(ur),
			SuspendUser:     command.NewSuspendUserHandler(ur),
			Follow:          command.NewFollowHandler(fr),
			Unfollow:        command.NewUnfollowHandler(fr),
		},
		Queries: application.UserQueries{
			GetUser:    query.NewGetUserHandler(ur),
			Followings: query.NewFollowingsHandler(ur),
			Followers:  query.NewFollowersHandler(ur),
		},
	}

	ca := application.ContentApplication{
		Commands: application.ContentCommands{
			Like:   command.NewLikeHandler(fvr),
			Unlike: command.NewUnlikeHandler(fvr),
		},
		Queries: application.ContentQueries{
			Content:             query.NewContentHandler(cr),
			GetFavoriteContents: query.NewGetFavoriteContentsHandler(cr),
		},
	}
	// Usecase

	return &Registry{
		AuthMiddleware: middleware.NewAuthMiddleware(uu),
		V1User:         v1.NewV1UserHandler(ua),
		V1Content:      v1.NewV1ContentHandler(ca),
	}
}
