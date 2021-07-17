package graph

import (
	"artics-api/src/internal/graph/model"
	"artics-api/src/internal/models"
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/volatiletech/null/v8"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var TableNames = models.TableNames

var UserColumns = models.UserColumns
var FavoriteColumns = models.FavoriteColumns
var ContentColumns = models.ContentColumns
var FollowColumns = models.FollowColumns
var CommentColumns = models.CommentColumns
var BrowseColumns = models.BrowseColumns
var NoticeColumns = models.NoticeColumns
var NoticeFavoriteColumns = models.NoticeFavoriteColumns
var NoticeFollowedColumns = models.NoticeFollowedColumns

var UserWhere = models.UserWhere
var FavoriteWhere = models.FavoriteWhere
var ContentWhere = models.ContentWhere
var FollowWhere = models.FollowWhere
var BrowseWhere = models.BrowseWhere
var NoticeWhere = models.NoticeWhere

type joinType int

const (
	INNER joinType = iota
	LEFT_OUTER
)

func join(t joinType, joinedTable string, joinedColumn string, joiningTable string, joiningColumn string) QueryMod {
	sql := strings.Join([]string{
		joinedTable,
		" on ",
		joinedTable,
		".",
		joinedColumn,
		"=",
		joiningTable,
		".",
		joiningColumn,
	}, "")
	switch t {
	case INNER:
		return InnerJoin(sql)
	case LEFT_OUTER:
		return LeftOuterJoin(sql)
	default:
		panic("invalid join type")
	}
}

type joinedFavorite struct {
	models.Favorite `boil:"favorite,bind"`
	models.User     `boil:"user,bind"`
	models.Content  `boil:"content,bind"`
	CommentBody     string `boil:"comment_body"`
}

func (r *Resolver) favorites(ctx context.Context, id *int, kind model.FavoriteKind) ([]*model.Favorite, error) {
	user := r.authUser(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	queries := []QueryMod{
		Select(
			TableNames.Favorite+"."+FavoriteColumns.ID,
			TableNames.User+"."+UserColumns.ID,
			TableNames.User+"."+UserColumns.ThumbnailURL,
			TableNames.Content+"."+ContentColumns.ID,
			TableNames.Content+"."+ContentColumns.Title,
			TableNames.Content+"."+ContentColumns.ThumbnailURL,
			TableNames.Comment+".body AS comment_body",
		),
		From(TableNames.Favorite),
		join(INNER, TableNames.User, UserColumns.ID, TableNames.Favorite, FavoriteColumns.UserID),
		join(INNER, TableNames.Content, ContentColumns.ID, TableNames.Favorite, FavoriteColumns.ContentID),
		join(LEFT_OUTER, TableNames.Comment, CommentColumns.ID, TableNames.Favorite, FavoriteColumns.CommentID),
	}

	switch kind {
	// Get a list of the user's favorites such that the content id is equal to that of the current user
	case model.FavoriteKindCommon:
		fs, err := models.Favorites(Select(FavoriteColumns.ID), FavoriteWhere.UserID.EQ(user.ID)).All(ctx, r.db)
		if err != nil {
			return nil, err
		}
		ids := make([]interface{}, len(fs))
		for i, f := range fs {
			ids[i] = f.ID
		}
		queries = append(queries, FavoriteWhere.UserID.EQ(*id))
		queries = append(queries, WhereIn(FavoriteColumns.ID+" in ?", ids))

	// The most favorites' of user
	case model.FavoriteKindMost:
		queries = append(queries, FavoriteWhere.UserID.EQ(*id))
		queries = append(queries, FavoriteWhere.IsMost.EQ(true))

	// Get a full list of the user's favorites
	case model.FavoriteKindOthers:
		queries = append(queries, FavoriteWhere.UserID.EQ(*id))

	// Get a list of favorites the user opts to share with the current user either as a close user or by selecting manually
	case model.FavoriteKindUserFavoriteForYou:
		closeFollows, err := models.Follows(
			Select(FollowColumns.FollowerID),
			FollowWhere.FollowingID.EQ(user.ID),
			FollowWhere.IsClose.EQ(true),
		).All(ctx, r.db)
		if err != nil {
			return nil, err
		}
		// A list of the ids of users who's followed by the current user and cofigure him/her as their close user
		// You can get the favorites' list shared as a close user by querying favorite table by the ids and the toCloseUser option
		ids := make([]interface{}, len(closeFollows))
		for i, f := range closeFollows {
			ids[i] = f.FollowerID
		}
		queries = append(queries, Expr(
			Expr(
				WhereIn(FavoriteColumns.UserID+" in ?", ids),
				FavoriteWhere.ToCloseUsers.EQ(true),
			),
			OrIn(FavoriteColumns.ToUserIds+" in ?", user.ID),
		))

	// Get a list of favorites of the current user's followings
	case model.FavoriteKindUserFavorite:
		follows, err := models.Follows(
			Select(FollowColumns.FollowerID),
			FollowWhere.FollowingID.EQ(user.ID),
		).All(ctx, r.db)
		if err != nil {
			return nil, err
		}
		ids := make([]interface{}, len(follows))
		for i, f := range follows {
			ids[i] = f.FollowerID
		}
		queries = append(queries, WhereIn(FavoriteColumns.UserID+" in ?", ids))
	}

	var fs []joinedFavorite
	if err := models.NewQuery(queries...).Bind(ctx, r.db, &fs); err != nil {
		return nil, err
	}

	resFavorites := make([]*model.Favorite, len(fs))
	for i, f := range fs {
		resFavorites[i] = &model.Favorite{
			ID: f.Favorite.ID,
			User: &model.User{
				ID:           f.User.ID,
				ThumbnailURL: &f.User.ThumbnailURL,
			},
			Content: &model.Content{
				ID:           f.Content.ID,
				Title:        f.Content.Title,
				ThumbnailURL: f.Content.ThumbnailURL,
			},
			CommentBody: f.CommentBody,
		}
	}
	return resFavorites, nil
}

func (r *Resolver) content(ctx context.Context, id int) (*model.Content, error) {
	user := r.authUser(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	c, err := models.FindContent(ctx, r.db, id)
	if err != nil {
		return nil, err
	}

	isLiked, err := models.Favorites(
		FavoriteWhere.ContentID.EQ(id),
		FavoriteWhere.UserID.EQ(user.ID),
	).Exists(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return &model.Content{
		ID:           c.ID,
		Title:        c.Title,
		ThumbnailURL: c.ThumbnailURL,
		Description:  c.Description,
		Author:       c.Author,
		IsLiked:      &isLiked,
	}, nil
}

func (r *Resolver) contents(ctx context.Context, id *int, kind model.ContentKind) ([]*model.Content, error) {
	user := r.authUser(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	var (
		cs  models.ContentSlice
		err error
	)

	switch kind {
	case model.ContentKindRecommended:
		// @TODO implementaion
		cs, err = models.Contents().All(ctx, r.db)

	case model.ContentKindBrowsed:
		cs, err = models.Contents(
			join(INNER, TableNames.Browse, BrowseColumns.ContentID, TableNames.Content, ContentColumns.ID),
			BrowseWhere.UserID.EQ(user.ID),
			OrderBy(TableNames.Browse+"."+BrowseColumns.UpdatedAt+" DESC"),
		).All(ctx, r.db)

	case model.ContentKindUploaded:
		cs, err = models.Contents(
			ContentWhere.UserID.EQ(null.IntFrom(*id)),
			OrderBy(ContentColumns.CreatedAt+" DESC"),
		).All(ctx, r.db)
	}

	if err != nil {
		return nil, err
	}

	resContents := make([]*model.Content, len(cs))
	for i, c := range cs {
		resContents[i] = &model.Content{
			ID:           c.ID,
			Title:        c.Title,
			ThumbnailURL: c.ThumbnailURL,
			Description:  c.Description,
			Author:       c.Author,
		}
	}
	return resContents, nil
}

func (r *Resolver) user(ctx context.Context, id int) (*model.User, error) {
	user := r.authUser(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	u, err := models.FindUser(ctx, r.db, id)
	if err != nil {
		return nil, err
	}

	followed, err := models.Follows(
		Select(FollowColumns.IsClose),
		FollowWhere.FollowingID.EQ(id),
		FollowWhere.FollowerID.EQ(user.ID),
	).One(ctx, r.db)

	isClose := false
	isFollowed := false
	if err == nil {
		isClose = followed.IsClose
		isFollowed = true
	} else if err != sql.ErrNoRows {
		return nil, err
	}

	isFollowing, err := models.Follows(
		FollowWhere.FollowingID.EQ(user.ID),
		FollowWhere.FollowerID.EQ(id),
	).Exists(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:           u.ID,
		Nickname:     &u.Nickname,
		ThumbnailURL: &u.ThumbnailURL,
		Profile:      &u.Profile,
		IsClose:      &isClose,
		IsFollowed:   &isFollowed,
		IsFollowing:  &isFollowing,
	}, nil
}

type joinedUser struct {
	models.User `boil:",bind"`
	IsClose     bool `boil:"is_close"`
}

func (r *Resolver) users(ctx context.Context, kind model.UserKind) ([]*model.User, error) {
	user := r.authUser(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	var us []joinedUser

	switch kind {
	case model.UserKindFollowing:
		userQuery := models.Users(
			InnerJoin(
				TableNames.Follow+" f on f."+FollowColumns.FollowerID+"="+TableNames.User+"."+UserColumns.ID,
			),
			FollowWhere.FollowingID.EQ(user.ID),
		)
		if err := userQuery.Bind(ctx, r.db, &us); err != nil {
			return nil, err
		}
	case model.UserKindFollower:
		userQuery := models.Users(
			Select(TableNames.User+"*", "f."+FollowColumns.IsClose+" AS is_close"),
			InnerJoin(
				TableNames.Follow+" f on f."+FollowColumns.FollowingID+"="+TableNames.User+"."+UserColumns.ID,
			),
			FollowWhere.FollowerID.EQ(user.ID),
		)
		if err := userQuery.Bind(ctx, r.db, &us); err != nil {
			return nil, err
		}
	}

	resUsers := make([]*model.User, len(us))
	for i, u := range us {
		resUsers[i] = &model.User{
			ID:           u.ID,
			Nickname:     &u.Nickname,
			ThumbnailURL: &u.ThumbnailURL,
			Profile:      &u.Profile,
			IsClose:      &u.IsClose,
		}
	}
	return resUsers, nil
}

func (r *Resolver) likedBy(ctx context.Context, contentID int) ([]*model.User, error) {
	us, err := models.Users(
		join(INNER, TableNames.Favorite, FavoriteColumns.UserID, TableNames.User, UserColumns.ID),
		FavoriteWhere.ContentID.EQ(contentID),
	).All(ctx, r.db)

	if err != nil {
		return nil, err
	}

	resUsers := make([]*model.User, len(us))
	for i, u := range us {
		resUsers[i] = &model.User{
			ID:           u.ID,
			ThumbnailURL: &u.ThumbnailURL,
		}
	}
	return resUsers, nil
}

type joinedNotice struct {
	models.Notice         `boil:",bind"`
	models.NoticeFavorite `boil:",bind"`
	models.NoticeFollowed `boil:",bind"`
}

func (r *Resolver) notices(ctx context.Context) ([]model.Notice, error) {
	user := r.authUser(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	query := models.NewQuery(
		Select(TableNames.Notice+".*", TableNames.NoticeFavorite+".*", TableNames.NoticeFollowed+".*"),
		From(TableNames.Notice),
		join(LEFT_OUTER, TableNames.NoticeFavorite, NoticeFavoriteColumns.NoticeID, TableNames.Notice, NoticeColumns.ID),
		join(LEFT_OUTER, TableNames.NoticeFollowed, NoticeFollowedColumns.NoticeID, TableNames.Notice, NoticeColumns.ID),
		NoticeWhere.UserID.EQ(user.ID),
		OrderBy(NoticeColumns.CreatedAt+" DESC"),
	)
	var ns []*joinedNotice
	if err := query.Bind(ctx, r.db, ns); err != nil {
		return nil, err
	}

	resNotices := make([]model.Notice, len(ns))
	for i, n := range ns {
		switch n.Notice.Type {
		case "favorite":
			resNotices[i] = model.NoticeFavorite{
				ID:         n.Notice.ID,
				Type:       n.Notice.Type,
				IsRead:     n.Notice.IsRead,
				CreatedAt:  n.Notice.CreatedAt.String(),
				FavoriteID: n.NoticeFavorite.FavoriteID,
				User: &model.User{
					ID:           n.NoticeFavorite.UserID,
					ThumbnailURL: &n.NoticeFavorite.UserThumbnailURL,
				},
				Header: n.NoticeFavorite.Header,
				Body:   n.NoticeFavorite.Body,
				Content: &model.Content{
					ID:           n.NoticeFavorite.ContentID,
					ThumbnailURL: n.NoticeFavorite.ContentThumbnailURL,
				},
			}
		case "followed":
			resNotices[i] = model.NoticeFollowed{
				ID:        n.Notice.ID,
				Type:      n.Notice.Type,
				IsRead:    n.Notice.IsRead,
				CreatedAt: n.Notice.CreatedAt.String(),
				User: &model.User{
					ID:           n.NoticeFollowed.UserID,
					ThumbnailURL: &n.NoticeFollowed.UserThumbnailURL,
				},
				Body: n.NoticeFollowed.Body,
			}
		}
	}

	return resNotices, nil
}

func (r *Resolver) currentUser(ctx context.Context) (*model.User, error) {
	user := r.authUser(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	return &model.User{
		ID:           user.ID,
		ThumbnailURL: &user.ThumbnailURL,
	}, nil
}

func (r *Resolver) searchUsers(ctx context.Context, words string) ([]*model.User, error) {
	us, err := models.Users(Where(UserColumns.Nickname+" LIKE ?", "%"+words+"%")).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	resUsers := make([]*model.User, len(us))
	for i, u := range us {
		resUsers[i] = &model.User{
			ID:           u.ID,
			ThumbnailURL: &u.ThumbnailURL,
			Nickname:     &u.Nickname,
			Profile:      &u.Profile,
		}
	}

	return resUsers, nil
}

func (r *Resolver) searchContents(ctx context.Context, words string) ([]*model.Content, error) {
	cs, err := models.Contents(Where(ContentColumns.Title+" LIKE ?", "%"+words+"%")).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	resContents := make([]*model.Content, len(cs))
	for i, c := range cs {
		resContents[i] = &model.Content{
			ID:           c.ID,
			ThumbnailURL: c.ThumbnailURL,
			Title:        c.Title,
		}
	}

	return resContents, nil
}
