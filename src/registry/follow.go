package registry

import (
	"artics-api/src/internal/infrastructure/repository"
	"artics-api/src/internal/infrastructure/service"
	v1 "artics-api/src/internal/interface/handler/v1"
	"artics-api/src/internal/usecase"
	"artics-api/src/internal/usecase/validation"
	"artics-api/src/lib/firebase"
	"artics-api/src/lib/mysql"
)

func v1FollowInjection(fa *firebase.Auth, db *mysql.Client) v1.V1FollowHandler {
	us := service.NewUserService(fa)
	fr := repository.NewFollowRepository(db)
	fu := usecase.NewFollowUsecase(us, fr)
	fh := v1.NewV1FollowHandler(fu)

	return fh
}
