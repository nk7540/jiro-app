package command

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/content"
	"artics-api/src/pkg"
)

type BrowseHandler struct {
	br content.BrowseRepository
}

func NewBrowseHandler(br content.BrowseRepository) BrowseHandler {
	return BrowseHandler{br}
}

func (h BrowseHandler) Handle(ctx pkg.Context, cmd content.CommandBrowse) error {
	b := &content.Browse{
		UserID:    cmd.UserID,
		ContentID: cmd.ContentID,
	}
	if err := h.br.Save(ctx, b); err != nil {
		return domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	}

	return nil
}
