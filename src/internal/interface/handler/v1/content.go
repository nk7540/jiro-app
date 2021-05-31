package v1

import (
	"artics-api/src/internal/interface/handler"
	"artics-api/src/internal/usecase"
	"artics-api/src/internal/usecase/response"
	"artics-api/src/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type V1ContentHandler interface {
	Favorites(c *gin.Context)
}

type v1ContentHandler struct {
	u usecase.ContentUsecase
}

func NewV1ContentHandler(u usecase.ContentUsecase) V1ContentHandler {
	return &v1ContentHandler{u}
}

func (h *v1ContentHandler) Favorites(c *gin.Context) {
	userId := c.Params.ByName("userId")
	ctx := middleware.GinContextToContext(c)
	cs, err := h.u.Favorites(ctx, userId)
	if err != nil {
		handler.ErrorHandling(c, err)
	}

	resContents := make([]*response.Content, len(cs))
	for i, c := range cs {
		resContents[i] = &response.Content{
			ID:    c.ID,
			Title: c.Title,
		}
	}
	res := &response.Contents{resContents}

	c.JSON(http.StatusOK, res)
}
