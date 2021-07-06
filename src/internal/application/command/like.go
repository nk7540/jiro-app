package command

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/content"
	"artics-api/src/internal/domain/notice"
	"artics-api/src/internal/domain/user"
	"artics-api/src/pkg"
)

type LikeHandler struct {
	cr  content.ContentRepository
	fr  content.FavoriteRepository
	cmr content.CommentRepository
	nr  notice.NoticeRepository
	ur  user.UserRepository
}

func NewLikeHandler(
	cr content.ContentRepository, fr content.FavoriteRepository, cmr content.CommentRepository,
	nr notice.NoticeRepository, ur user.UserRepository,
) LikeHandler {
	return LikeHandler{cr, fr, cmr, nr, ur}
}

func (h LikeHandler) Handle(ctx pkg.Context, cmd content.CommandLike) ([]*notice.QueryNotice, error) {
	// Create a comment if the body exists
	var (
		commentID content.CommentID
		err       error
	)
	if cmd.CommentBody != "" {
		c := &content.Comment{
			UserID:    cmd.User.ID,
			ContentID: cmd.ContentID,
			Body:      cmd.CommentBody,
		}
		commentID, err = h.cmr.Create(ctx, c)
		if err != nil {
			return nil, domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
		}
	}

	// Create a favorite
	f := &content.Favorite{
		UserID:       cmd.User.ID,
		ContentID:    cmd.ContentID,
		CommentID:    commentID,
		ToUserIDs:    cmd.ToUserIDs,
		ToCloseUsers: cmd.ToCloseUsers,
	}
	if err := h.fr.Create(ctx, f); err != nil {
		return nil, domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	}

	// Fetch the favorite content
	qc, err := h.cr.GetOrNone(ctx, cmd.ContentID)
	if err != nil || qc == nil {
		return nil, domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	}
	c := &content.Content{
		ID:           content.ContentID(qc.ID),
		Title:        content.Title(qc.Title),
		ThumbnailURL: content.ThumbnailURL(qc.ThumbnailURL),
	}

	notices := make([]*notice.QueryNotice, 0)
	p, err := ctx.Printer()
	if err != nil {
		return nil, err
	}

	// Notify to the selected users
	for _, userID := range cmd.ToUserIDs {
		n := notice.NewFavoriteNotice(p, userID, cmd.User, c, cmd.CommentBody)
		if err := h.nr.Create(ctx, n); err != nil {
			return nil, domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
		}
		notices = append(notices, toQueryNotice(n))
	}

	// Notify to the close users of current user if specified
	if cmd.ToCloseUsers {
		us, err := h.ur.Followings(ctx, cmd.User.ID, true)
		if err != nil {
			return nil, domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
		}
		for _, u := range us.Users {
			n := notice.NewFavoriteNotice(p, user.UserID(u.ID), cmd.User, c, cmd.CommentBody)
			if err := h.nr.Create(ctx, n); err != nil {
				return nil, domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
			}
			notices = append(notices, toQueryNotice(n))
		}
	}

	return notices, nil
}

func toQueryNotice(n *notice.Notice) *notice.QueryNotice {
	return &notice.QueryNotice{
		ID:     int(n.ID),
		UserID: int(n.UserID),
		Type:   int(n.Type),
		IsRead: bool(n.IsRead),
		Favorite: &notice.QueryNoticeFavorite{
			FavoriteID:          int(n.Favorite.FavoriteID),
			UserID:              int(n.Favorite.UserID),
			UserThumbnailURL:    string(n.Favorite.UserThumbnailURL),
			Header:              string(n.Favorite.Header),
			Body:                string(n.Favorite.Body),
			ContentID:           int(n.Favorite.ContentID),
			ContentThumbnailURL: string(n.Favorite.ContentThumbnailURL),
		},
	}
}
