package v1

import (
	"artics-api/src/internal/interface/handler"
	"artics-api/src/internal/usecase"
	"artics-api/src/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type V1FavoriteHandler interface {
	Create(c *gin.Context)
}

type v1FavoriteHandler struct {
	u usecase.FavoriteUsecase
}

func NewV1FavoriteHandler(u usecase.FavoriteUsecase) V1FavoriteHandler {
	return &v1FavoriteHandler{u}
}

func (h *v1FavoriteHandler) Create(c *gin.Context) {
	contentID := c.Params.ByName("content_id")
	ctx := middleware.GinContextToContext(c)

	if err := h.u.Create(ctx, contentID); err != nil {
		handler.ErrorHandling(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
