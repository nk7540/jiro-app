package command

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/content"
	"artics-api/src/internal/domain/user"
	"artics-api/src/pkg"
)

type LikeHandler struct {
	cr  content.ContentRepository
	fr  content.FavoriteRepository
	cmr content.CommentRepository
	nr  content.NoticeRepository
	ur  user.UserRepository
}

func NewLikeHandler(
	cr content.ContentRepository, fr content.FavoriteRepository, cmr content.CommentRepository,
	nr content.NoticeRepository, ur user.UserRepository,
) LikeHandler {
	return LikeHandler{cr, fr, cmr, nr, ur}
}

func (h LikeHandler) Handle(ctx pkg.Context, cmd content.CommandLike) ([]*content.Notice, error) {
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
	c, err := h.cr.GetOrNone(ctx, cmd.ContentID)
	if err != nil || c == nil {
		return nil, domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	}

	notices := make([]*content.Notice, 0)
	p, err := ctx.Printer()
	if err != nil {
		return nil, err
	}

	// Notify to the selected users
	for _, userID := range cmd.ToUserIDs {
		n := content.NewFavoriteNotice(p, userID, cmd.User.Nickname, content.Title(c.Title), cmd.CommentBody)
		if err := h.nr.Create(ctx, n); err != nil {
			return nil, domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
		}
		notices = append(notices, n)
	}

	// Notify to the close users of current user if specified
	if cmd.ToCloseUsers {
		us, err := h.ur.CloseUsers(ctx, cmd.User.ID)
		if err != nil {
			return nil, domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
		}
		for _, u := range us.Users {
			n := content.NewFavoriteNotice(p, user.UserID(u.ID), cmd.User.Nickname, content.Title(c.Title), cmd.CommentBody)
			if err := h.nr.Create(ctx, n); err != nil {
				return nil, domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
			}
			notices = append(notices, n)
		}
	}

	return notices, nil
}
