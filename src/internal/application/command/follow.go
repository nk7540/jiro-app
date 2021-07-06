package command

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/notice"
	"artics-api/src/internal/domain/user"
	"artics-api/src/pkg"
)

type FollowHandler struct {
	fr user.FollowRepository
	nr notice.NoticeRepository
}

func NewFollowHandler(fr user.FollowRepository, nr notice.NoticeRepository) FollowHandler {
	return FollowHandler{fr, nr}
}

func (h FollowHandler) Handle(ctx pkg.Context, cmd user.CommandFollow) (*notice.QueryNotice, error) {
	f := &user.Follow{
		FollowingID: user.FollowingID(cmd.User.ID),
		FollowerID:  cmd.FollowerID,
	}

	if err := h.fr.Create(ctx, f); err != nil {
		return nil, domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	}

	p, err := ctx.Printer()
	if err != nil {
		return nil, err
	}
	n := notice.NewFollowedNotice(p, cmd.FollowerID, cmd.User)
	if err := h.nr.Create(ctx, n); err != nil {
		return nil, domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	}

	return &notice.QueryNotice{
		ID:     int(n.ID),
		UserID: int(n.UserID),
		Type:   int(n.Type),
		IsRead: bool(n.IsRead),
		Followed: &notice.QueryNoticeFollowed{
			UserID:           int(n.Followed.UserID),
			UserThumbnailURL: string(n.Followed.UserThumbnailURL),
			Body:             string(n.Followed.Body),
		},
	}, nil
}
