package v1

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/interface/handler"
	"artics-api/src/internal/usecase"
	"artics-api/src/internal/usecase/response"
	"artics-api/src/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type V1ContentHandler interface {
	Show(c *gin.Context)
	Favorites(c *gin.Context)
}

type v1ContentHandler struct {
	u usecase.ContentUsecase
}

func NewV1ContentHandler(u usecase.ContentUsecase) V1ContentHandler {
	return &v1ContentHandler{u}
}

func (h *v1ContentHandler) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		handler.ErrorHandling(c, domain.UnableParseJSON.New(err))
		return
	}
	ctx := middleware.GinContextToContext(c)

	content, err := h.u.Show(ctx, id)
	if err != nil {
		handler.ErrorHandling(c, err)
		return
	}

	res := &response.Content{
		ID:    content.ID,
		Title: content.Title,
	}

	c.JSON(http.StatusOK, res)
}

func (h *v1ContentHandler) Favorites(c *gin.Context) {
	userID, err := strconv.Atoi(c.Params.ByName("user_id"))
	if err != nil {
		handler.ErrorHandling(c, domain.UnableParseJSON.New(err))
		return
	}
	ctx := middleware.GinContextToContext(c)
	cs, err := h.u.Favorites(ctx, userID)
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
