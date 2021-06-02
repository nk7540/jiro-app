package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"artics-api/src/internal/domain"
	"artics-api/src/internal/interface/handler"
	"artics-api/src/internal/usecase"
	"artics-api/src/internal/usecase/request"
	"artics-api/src/middleware"
)

// V1UserHandler - v1 user handler
type V1UserHandler interface {
	Create(ctx *gin.Context)
	Show(ctx *gin.Context)
	Followings(ctx *gin.Context)
	Followers(ctx *gin.Context)
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
	req := &request.CreateUser{}
	if err := c.BindJSON(req); err != nil {
		handler.ErrorHandling(c, domain.UnableParseJSON.New(err))
		return
	}

	ctx := middleware.GinContextToContext(c)
	if err := h.u.Create(ctx, req); err != nil {
		handler.ErrorHandling(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *v1UserHandler) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		handler.ErrorHandling(c, domain.UnableParseJSON.New(err))
		return
	}

	ctx := middleware.GinContextToContext(c)
	res, err := h.u.Show(ctx, id)
	if err != nil {
		handler.ErrorHandling(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *v1UserHandler) Followings(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		handler.ErrorHandling(c, domain.UnableParseJSON.New(err))
		return
	}
	ctx := middleware.GinContextToContext(c)

	res, err := h.u.Followings(ctx, id)
	if err != nil {
		handler.ErrorHandling(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *v1UserHandler) Followers(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		handler.ErrorHandling(c, domain.UnableParseJSON.New(err))
		return
	}
	ctx := middleware.GinContextToContext(c)

	res, err := h.u.Followers(ctx, id)
	if err != nil {
		handler.ErrorHandling(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *v1UserHandler) Update(c *gin.Context) {
	req := &request.UpdateUser{}
	if err := c.Bind(req); err != nil {
		handler.ErrorHandling(c, domain.UnableParseFormData.New(err))
		return
	}

	ctx := middleware.GinContextToContext(c)
	res, err := h.u.Update(ctx, req)
	if err != nil {
		handler.ErrorHandling(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *v1UserHandler) Suspend(c *gin.Context) {
	ctx := middleware.GinContextToContext(c)

	if err := h.u.Suspend(ctx); err != nil {
		handler.ErrorHandling(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
