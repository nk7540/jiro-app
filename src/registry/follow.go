package registry

import (
	"artics-api/src/internal/infrastructure/repository"
	"artics-api/src/internal/infrastructure/service"
	v1 "artics-api/src/internal/interface/handler/v1"
	"artics-api/src/internal/usecase"
	"artics-api/src/lib/firebase"
	"artics-api/src/lib/mysql"
)

func v1FollowInjection(fa *firebase.Auth, db *mysql.Client) v1.V1FollowHandler {
	ur := repository.NewUserRepository(db, fa)
	fr := repository.NewFollowRepository(db)
	us := service.NewUserService(ur, fr)
	fu := usecase.NewFollowUsecase(fr, us)
	fh := v1.NewV1FollowHandler(fu)

	return fh
}
