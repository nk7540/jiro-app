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
	Update(ctx *gin.Context)
	Suspend(ctx *gin.Context)
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
	req := &request.CreateUser{}
	if err := c.BindJSON(req); err != nil {
		handler.ErrorHandling(c, domain.UnableParseJSON.New(err))
		return
	}

	// Response
	ctx := middleware.GinContextToContext(c)
	if err := h.u.Create(ctx, req); err != nil {
		handler.ErrorHandling(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *v1UserHandler) Show(c *gin.Context) {
	// Request
	id := c.Params.ByName("id")

	// Response
	ctx := middleware.GinContextToContext(c)
	u, err := h.u.Show(ctx, id)
	if err != nil {
		handler.ErrorHandling(c, err)
		return
	}

	res := &response.ShowUser{
		ID:       u.ID,
		Nickname: u.Nickname,
		Email:    u.Email,
	}

	c.JSON(http.StatusOK, res)
}

func (h *v1UserHandler) Update(c *gin.Context) {
	// Request
	req := &request.UpdateUser{}
	if err := c.BindJSON(req); err != nil {
		handler.ErrorHandling(c, domain.UnableParseJSON.New(err))
		return
	}

	// Response
	ctx := middleware.GinContextToContext(c)
	u, err := h.u.Update(ctx, req)
	if err != nil {
		handler.ErrorHandling(c, err)
		return
	}

	res := &response.UpdateUser{
		ID:       u.ID,
		Nickname: u.Nickname,
		Email:    u.Email,
	}

	c.JSON(http.StatusOK, res)
}

func (h *v1UserHandler) Suspend(c *gin.Context) {
	// Request
	ctx := middleware.GinContextToContext(c)

	// Response
	if err := h.u.Suspend(ctx); err != nil {
		handler.ErrorHandling(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
