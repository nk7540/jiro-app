package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
)

type key int

var GinContext key

func GinContextToContext(ctx *gin.Context) context.Context {
	return context.WithValue(ctx.Request.Context(), GinContext, ctx)
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(GinContext)
	if ginContext == nil {
		err := xerrors.New("Could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := xerrors.New("gin.Context has wrong type")
		return nil, err
	}

	return gc, nil
}
