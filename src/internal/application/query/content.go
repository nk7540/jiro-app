package query

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/content"
	"artics-api/src/pkg"
)

type ContentHandler struct {
	cr content.ContentRepository
}

func NewContentHandler(cr content.ContentRepository) ContentHandler {
	return ContentHandler{cr}
}

func (h ContentHandler) Handle(ctx pkg.Context, id content.ContentID) (*content.QueryDetailContent, error) {
	c, err := h.cr.GetOrNone(ctx, id)
	if err != nil {
		return nil, domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	} else if c == nil {
		return nil, domain.NotFound.New(pkg.NewRepositoryError(err))
	}

	return c, nil
}
