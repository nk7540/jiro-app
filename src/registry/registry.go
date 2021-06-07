package registry

import (
	"artics-api/src/config"
	"artics-api/src/internal/infrastructure/repository"
	"artics-api/src/internal/infrastructure/service"
	dv "artics-api/src/internal/infrastructure/validation"
	v1 "artics-api/src/internal/interface/handler/v1"
	"artics-api/src/internal/interface/middleware"
	"artics-api/src/internal/usecase"
	v "artics-api/src/internal/usecase/validation"
)

// Registry - DI container
type Registry struct {
	AuthMiddleware middleware.AuthMiddleware
	V1User         v1.V1UserHandler
	V1Follow       v1.V1FollowHandler
	V1Content      v1.V1ContentHandler
	V1Favorite     v1.V1FavoriteHandler
	V1Browse       v1.V1BrowseHandler
	// CategoryHandler handler.CategoryHandler
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

	// Request Validator
	rv := v.NewRequestValidator()

	// Usecase
	uu := usecase.NewUserUsecase(rv, us, cs)
	fu := usecase.NewFollowUsecase(fr, us)
	cu := usecase.NewContentUsecase(cs)
	fvu := usecase.NewFavoriteUsecase(us, fvs)
	bu := usecase.NewBrowseUsecase(us, bs)

	return &Registry{
		AuthMiddleware: middleware.NewAuthMiddleware(uu),
		V1User:         v1.NewV1UserHandler(uu),
		V1Follow:       v1.NewV1FollowHandler(fu),
		V1Content:      v1.NewV1ContentHandler(cu),
		V1Favorite:     v1.NewV1FavoriteHandler(fvu),
		V1Browse:       v1.NewV1BrowseHandler(bu),
	}
}
