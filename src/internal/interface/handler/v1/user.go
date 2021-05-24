package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"artics-api/src/internal/domain"
	"artics-api/src/internal/interface/handler"
	"artics-api/src/internal/usecase"
	"artics-api/src/internal/usecase/request"
	"artics-api/src/internal/usecase/response"
	"artics-api/src/middleware"
)

// V1UserHandler - v1 user handler
type V1UserHandler interface {
	Create(ctx *gin.Context)
	Show(ctx *gin.Context)
}

type v1UserHandler struct {
	u usecase.UserUsecase
}

// NewV1UserHandler - setups v1 user handler
func NewV1UserHandler(u usecase.UserUsecase) V1UserHandler {
	return &v1UserHandler{u}
}

func (h *v1UserHandler) Create(c *gin.Context) {
	// Request
	ctx := middleware.GinContextToContext(c)
	req := &request.CreateUser{}
	if err := ctx.bindJSON(req); err != nil {
		handler.ErrorHandling(ctx, domain.UnableParseJSON.New(err))
		return
	}

	// Response
	_, err := h.u.Create(ctx, req)
	if err != nil {
		handler.ErrorHandling(ctx, err)
		return
	}

	res := &response.CreateUser{
		resultCode: "success",
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *v1UserHandler) Show(c *gin.Context) {
	// Request
	id := c.Params.ByName("id")
	ctx := middleware.GinContextToContext(c)

	// Response
	u, err := h.u.Show(ctx, id)
	if err != nil {
		handler.ErrorHandling(ctx, err)
		return
	}

	res := &response.ShowUser{
		ID:       u.ID,
		Nickname: u.Nickname,
		Email:    u.Email,
	}

	ctx.JSON(http.StatusOK, res)
}
