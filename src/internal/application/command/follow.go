package command

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/content"
	"artics-api/src/internal/domain/user"
	"artics-api/src/pkg"
)

type FollowHandler struct {
	fr user.FollowRepository
	nr content.NoticeRepository
}

func NewFollowHandler(fr user.FollowRepository, nr content.NoticeRepository) FollowHandler {
	return FollowHandler{fr, nr}
}

func (h FollowHandler) Handle(ctx pkg.Context, cmd user.CommandFollow) (*content.Notice, error) {
	f := &user.Follow{
		FollowingID: user.FollowingID(cmd.User.ID),
		FollowerID:  cmd.FollowerID,
	}

	if err := h.fr.Create(ctx, f); err != nil {
		return nil, domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	}

	n := content.NewFollowedNotice(cmd.FollowerID, cmd.User.Nickname)
	if err := h.nr.Create(ctx, n); err != nil {
		return nil, domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	}

	return n, nil
}
