package v1

import (
	"artics-api/src/internal/interface/handler"
	"artics-api/src/internal/usecase"
	"artics-api/src/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type V1BrowseHandler interface {
	Save(c *gin.Context)
}

type v1BrowseHandler struct {
	u usecase.BrowseUsecase
}

func NewV1BrowseHandler(u usecase.BrowseUsecase) V1BrowseHandler {
	return &v1BrowseHandler{u}
}

func (h *v1BrowseHandler) Save(c *gin.Context) {
	contentID := c.Params.ByName("content_id")
	ctx := middleware.GinContextToContext(c)

	if err := h.u.Save(ctx, contentID); err != nil {
		handler.ErrorHandling(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
