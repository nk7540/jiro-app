package notice

import (
	"artics-api/src/internal/domain/user"
)

type QueryArgNotices struct {
	UserID user.UserID
	Page   int
	Per    int
}

type QueryArgListNotice struct {
	UserID user.UserID
	Limit  int
	Offset int
}

type QueryNoticeFavorite struct {
	FavoriteID          int    `boil:"favorite_id"`
	UserID              int    `boil:"user_id"`
	UserThumbnailURL    string `boil:"user_thumbnail_url"`
	Header              string `boil:"header"`
	Body                string `boil:"body"`
	ContentID           int    `boil:"content_id"`
	ContentThumbnailURL string `boil:"content_thumbnail_url"`
}

type QueryNoticeFollowed struct {
	UserID           int    `boil:"user_id"`
	UserThumbnailURL string `boil:"user_thumbnail_url"`
	Body             string `boil:"body"`
}

type QueryNotice struct {
	ID       int                  `boil:"id"`
	UserID   int                  `boil:"user_id"`
	Type     int                  `boil:"type"`
	IsRead   bool                 `boil:"is_read"`
	Favorite *QueryNoticeFavorite `boil:"n_favorite,bind"`
	Followed *QueryNoticeFollowed `boil:"n_followed,bind"`
}

type QueryNotices struct {
	Notices []*QueryNotice
}
