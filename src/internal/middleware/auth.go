package middleware

import (
	"artics-api/src/config"
	"context"
	"net/http"
	"strings"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var tokenCtxKey = &contextKey{"token"}

type contextKey struct {
	token string
}

// Middleware decodes the share session cookie and packs the session into context
func Auth(db *config.DatabaseConfig, auth *config.AuthConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			a := r.Header.Get("Authorization")
			token := strings.Replace(a, "Bearer ", "", 1)

			// put it in context
			ctx = context.WithValue(ctx, tokenCtxKey, token)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) string {
	raw, _ := ctx.Value(tokenCtxKey).(string)
	return raw
}
