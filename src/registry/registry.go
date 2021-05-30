package registry

import (
	v1 "artics-api/src/internal/interface/handler/v1"
	"artics-api/src/lib/firebase"
	"artics-api/src/lib/grpc"
	"artics-api/src/lib/mysql"
)

// Registry - DI container
type Registry struct {
	V1User v1.V1UserHandler
	V1Follow v1.V1FollowHandler
	// ContentHandler handler.ContentHandler
	// FavoriteHandler handler.FavoriteHandler
	// CategoryHandler handler.CategoryHandler
}

// NewRegistry - imports files in /internal directory
func NewRegistry(
	fa *firebase.Auth, db *mysql.Client, gc *grpc.Client,
) *Registry {
	v1User := v1UserInjection(fa, db)
	v1Follow := v1FollowInjection(fa, db)

	return &Registry{
		V1User: v1User,
		V1Follow: v1Follow,
	}
}
