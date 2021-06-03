package registry

import (
	"artics-api/src/internal/infrastructure/repository"
	"artics-api/src/internal/infrastructure/service"
	dv "artics-api/src/internal/infrastructure/validation"
	v1 "artics-api/src/internal/interface/handler/v1"
	"artics-api/src/internal/usecase"
	v "artics-api/src/internal/usecase/validation"
	"artics-api/src/lib/awssdk"
	"artics-api/src/lib/firebase"
	"artics-api/src/lib/gmail"
	"artics-api/src/lib/grpc"
	"artics-api/src/lib/mysql"
)

// Registry - DI container
type Registry struct {
	V1User     v1.V1UserHandler
	V1Follow   v1.V1FollowHandler
	V1Content  v1.V1ContentHandler
	V1Favorite v1.V1FavoriteHandler
	V1Browse   v1.V1BrowseHandler
	// CategoryHandler handler.CategoryHandler
}

// NewRegistry - imports files in /internal directory
func NewRegistry(
	au *awssdk.Uploader, fa *firebase.Auth, gm *gmail.Client, db *mysql.Client, gc *grpc.Client,
) *Registry {
	// Domain Repository
	ur := repository.NewUserRepository(db, fa)
	fr := repository.NewFollowRepository(db)
	cr := repository.NewContentRepository(db)
	fvr := repository.NewFavoriteRepository(db)
	br := repository.NewBrowseRepository(db)
	flr := repository.NewFileRepository(au)

	// Domain Validator
	udv := dv.NewUserDomainValidator(ur)

	// Domain Service
	us := service.NewUserService(gm, udv, ur, fr, cr, flr)
	cs := service.NewContentService(cr)
	fvs := service.NewFavoriteService(fvr)
	bs := service.NewBrowseService(br)

	// Request Validator
	rv := v.NewRequestValidator()

	// Usecase
	uu := usecase.NewUserUsecase(rv, ur, us, cs)
	fu := usecase.NewFollowUsecase(fr, us)
	cu := usecase.NewContentUsecase(cs)
	fvu := usecase.NewFavoriteUsecase(us, fvs)
	bu := usecase.NewBrowseUsecase(us, bs)

	return &Registry{
		V1User:     v1.NewV1UserHandler(uu),
		V1Follow:   v1.NewV1FollowHandler(fu),
		V1Content:  v1.NewV1ContentHandler(cu),
		V1Favorite: v1.NewV1FavoriteHandler(fvu),
		V1Browse:   v1.NewV1BrowseHandler(bu),
	}
}
