package graph

import (
	"artics-api/src/internal/graph/model"
	"artics-api/src/internal/models"
	"artics-api/src/utils"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (r *Resolver) createUser(ctx context.Context, input model.CreateUser) (*int, error) {
	// @TODO validation

	uid, err := r.auth.CreateUser(ctx, input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	u := models.User{
		UID:   uid,
		Email: input.Email,
	}
	if err := u.Insert(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}

	return &u.ID, nil
}

func (r *Resolver) updateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	user := r.authUser(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	// @TODO validation
	if input.Thumbnail != nil {
		output, err := r.uploader.Upload(input.Thumbnail.File, input.Thumbnail.Filename)
		if err != nil {
			return nil, err
		}
		user.ThumbnailURL = output.Location
	}

	user.Nickname = *input.Nickname
	user.Profile = *input.Profile

	if _, err := user.Update(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}

	return &model.User{
		ID:           user.ID,
		ThumbnailURL: &user.ThumbnailURL,
		Nickname:     &user.Nickname,
		Profile:      &user.Profile,
	}, nil
}

func (r *Resolver) like(ctx context.Context, input model.Like) (*int, error) {
	user := r.authUser(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	// Create a comment if the body exists
	var cm models.Comment
	if input.CommentBody != "" {
		cm = models.Comment{ContentID: input.ContentID, Body: input.CommentBody}
		if err := cm.Insert(ctx, r.db, boil.Infer()); err != nil {
			return nil, err
		}
	}

	// Create a favorite
	f := models.Favorite{
		UserID:       user.ID,
		ContentID:    input.ContentID,
		CommentID:    null.IntFrom(cm.ID),
		ToUserIds:    utils.IntSliceToString(input.ToUserIds),
		ToCloseUsers: input.ToCloseUsers,
	}
	if err := f.Insert(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}

	// Fetch the favorite content
	c, err := models.FindContent(ctx, r.db, f.ContentID)
	if err != nil {
		return nil, err
	}

	// Notify to the selected users
	for _, userID := range input.ToUserIds {
		if err := r.createNotice(ctx, userID, f, cm, *c); err != nil {
			return nil, err
		}
	}

	// Notify to the close users of current user if specified
	if f.ToCloseUsers {
		us, err := models.Users(
			join(INNER, TableNames.Follow, FollowColumns.FollowingID, TableNames.User, UserColumns.ID),
			FollowWhere.FollowerID.EQ(user.ID),
			FollowWhere.IsClose.EQ(true),
		).All(ctx, r.db)
		if err != nil {
			return nil, err
		}
		for _, u := range us {
			if err := r.createNotice(ctx, u.ID, f, cm, *c); err != nil {
				return nil, err
			}
		}
	}

	return &f.ID, nil
}

func (r *Resolver) unlike(ctx context.Context, contentID int) (*int, error) {
	user := r.authUser(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	f, err := models.Favorites(
		FavoriteWhere.UserID.EQ(user.ID),
		FavoriteWhere.ContentID.EQ(contentID),
	).One(ctx, r.db)
	if err != nil {
		return nil, err
	}
	_, err = f.Delete(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return &f.ID, nil
}

func (r *Resolver) browse(ctx context.Context, contentID int) (*int, error) {
	user := r.authUser(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	b, err := models.Browses(
		BrowseWhere.UserID.EQ(user.ID),
		BrowseWhere.ContentID.EQ(contentID),
	).One(ctx, r.db)

	if err == sql.ErrNoRows {
		b = &models.Browse{
			UserID:    user.ID,
			ContentID: contentID,
		}
		if err := b.Insert(ctx, r.db, boil.Infer()); err != nil {
			return nil, err
		}
	} else if err == nil {
		b.UpdatedAt = time.Now()
		if _, err := b.Update(ctx, r.db, boil.Whitelist("updated_at")); err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	return &b.ID, nil
}

func (r *Resolver) createContent(ctx context.Context, input model.CreateContent) (*int, error) {
	user := r.authUser(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	// @TODO validation
	c := &models.Content{
		UserID:      null.IntFrom(user.ID),
		Title:       input.Title,
		Description: *input.Description,
		Author:      *input.Author,
	}

	if input.Thumbnail != nil {
		output, err := r.uploader.Upload(input.Thumbnail.File, input.Thumbnail.Filename)
		if err != nil {
			return nil, err
		}
		c.ThumbnailURL = output.Location
	}

	if err := c.Insert(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}

	return &c.ID, nil
}

func (r *Resolver) follow(ctx context.Context, followerID int) (*int, error) {
	user := r.authUser(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	f := models.Follow{
		FollowingID: user.ID,
		FollowerID:  followerID,
	}
	if err := f.Insert(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}

	return &f.ID, nil
}

func (r *Resolver) unfollow(ctx context.Context, followerID int) (*int, error) {
	user := r.authUser(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	f, err := models.Follows(
		FollowWhere.FollowingID.EQ(user.ID),
		FollowWhere.FollowerID.EQ(followerID),
	).One(ctx, r.db)
	if err != nil {
		return nil, err
	}
	_, err = f.Delete(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return &f.ID, nil
}

func (r *Resolver) createNotice(ctx context.Context, userID int, f models.Favorite, cm models.Comment, c models.Content) error {
	user := r.authUser(ctx)
	if user == nil {
		return errors.New("access denied")
	}

	n := models.Notice{UserID: userID, Type: "favorite"}
	if err := n.Insert(ctx, r.db, boil.Infer()); err != nil {
		return err
	}
	nf := models.NoticeFavorite{
		NoticeID:            n.ID,
		FavoriteID:          f.ID,
		UserID:              user.ID,
		UserThumbnailURL:    user.ThumbnailURL,
		Header:              "", // @TODO Message
		Body:                cm.Body,
		ContentID:           c.ID,
		ContentThumbnailURL: c.ThumbnailURL,
	}
	return nf.Insert(ctx, r.db, boil.Infer())
}
