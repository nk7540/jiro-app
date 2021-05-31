package registry

import (
	"artics-api/src/internal/infrastructure/repository"
	"artics-api/src/internal/infrastructure/service"
	dv "artics-api/src/internal/infrastructure/validation"
	v1 "artics-api/src/internal/interface/handler/v1"
	"artics-api/src/internal/usecase"
	v "artics-api/src/internal/usecase/validation"
	"artics-api/src/lib/firebase"
	"artics-api/src/lib/grpc"
	"artics-api/src/lib/mysql"
)

// Registry - DI container
type Registry struct {
	V1User   v1.V1UserHandler
	V1Follow v1.V1FollowHandler
	// ContentHandler handler.ContentHandler
	// FavoriteHandler handler.FavoriteHandler
	// CategoryHandler handler.CategoryHandler
}

// NewRegistry - imports files in /internal directory
func NewRegistry(
	fa *firebase.Auth, db *mysql.Client, gc *grpc.Client,
) *Registry {
	// Domain Repository
	ur := repository.NewUserRepository(db, fa)
	fr := repository.NewFollowRepository(db)
	cr := repository.NewContentRepository(db)

	// Domain Validator
	udv := dv.NewUserDomainValidator(ur)

	// Domain Service
	us := service.NewUserService(udv, ur, fr, cr)
	cs := service.NewContentService(cr)

	// Request Validator
	rv := v.NewRequestValidator()

	// Usecase
	uu := usecase.NewUserUsecase(rv, ur, us, cs)
	fu := usecase.NewFollowUsecase(fr, us)

	return &Registry{
		V1User:   v1.NewV1UserHandler(uu),
		V1Follow: v1.NewV1FollowHandler(fu),
	}
}
