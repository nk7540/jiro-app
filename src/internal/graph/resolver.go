package graph

import (
	"artics-api/src/config"
	"artics-api/src/internal/middleware"
	"artics-api/src/internal/models"
	"context"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	uploader *config.UploaderConfig
	auth     *config.AuthConfig
	mail     *config.MailConfig
	db       *config.DatabaseConfig
	rpc      *config.RPCConfig
}

func NewResolver(
	uploader *config.UploaderConfig,
	auth *config.AuthConfig,
	mail *config.MailConfig,
	db *config.DatabaseConfig,
	rpc *config.RPCConfig,
) *Resolver {
	return &Resolver{uploader, auth, mail, db, rpc}
}

func (r *Resolver) authUser(ctx context.Context) *models.User {
	token := middleware.ForContext(ctx)
	uid, err := r.auth.VerifyIDToken(ctx, token)
	if err != nil {
		return nil
	}

	u, err := models.Users(models.UserWhere.UID.EQ(uid)).One(ctx, r.db)
	if err != nil {
		return nil
	}

	return u
}
