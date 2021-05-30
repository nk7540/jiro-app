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

func v1UserInjection(fa *firebase.Auth, db *mysql.Client) v1.V1UserHandler {
	fr := repository.NewFollowRepository(db)
	ur := repository.NewUserRepository(db, fa)
	us := service.NewUserService(ur, fr)
	urv := validation.NewUserRequestValidator()
	uu := usecase.NewUserUsecase(urv, ur, us, fr)
	uh := v1.NewV1UserHandler(uu)

	return uh
}
