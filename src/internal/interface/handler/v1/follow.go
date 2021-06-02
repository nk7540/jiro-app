package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"artics-api/src/internal/domain"
	"artics-api/src/internal/interface/handler"
	"artics-api/src/internal/usecase"
	"artics-api/src/middleware"
)

// V1FollowHandler - v1 follow handler
type V1FollowHandler interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type v1FollowHandler struct {
	u usecase.FollowUsecase
}

// NewV1FollowHandler - setups v1 follow handler
func NewV1FollowHandler(u usecase.FollowUsecase) V1FollowHandler {
	return &v1FollowHandler{u}
}

func (h *v1FollowHandler) Create(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("user_id"))
	if err != nil {
		handler.ErrorHandling(c, domain.UnableParseJSON.New(err))
		return
	}
	ctx := middleware.GinContextToContext(c)

	if err := h.u.Create(ctx, id); err != nil {
		handler.ErrorHandling(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *v1FollowHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("user_id"))
	if err != nil {
		handler.ErrorHandling(c, domain.UnableParseJSON.New(err))
		return
	}
	ctx := middleware.GinContextToContext(c)

	if err := h.u.Delete(ctx, id); err != nil {
		handler.ErrorHandling(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
