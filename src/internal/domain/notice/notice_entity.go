package notice

import (
	"artics-api/src/config"
	"artics-api/src/internal/domain/content"
	"artics-api/src/internal/domain/user"
)

type Notice struct {
	ID       NoticeID
	UserID   user.UserID
	Type     NoticeType
	IsRead   NoticeIsRead
	Favorite NoticeFavorite
	Followed NoticeFollowed
}

type NoticeFavorite struct {
	NoticeID            NoticeID
	FavoriteID          content.FavoriteID
	UserID              user.UserID
	UserThumbnailURL    user.ThumbnailURL
	Header              NoticeFavoriteHeader
	Body                NoticeFavoriteBody
	ContentID           content.ContentID
	ContentThumbnailURL content.ThumbnailURL
}

type NoticeFollowed struct {
	NoticeID         NoticeID
	UserID           user.UserID
	UserThumbnailURL user.ThumbnailURL
	Body             NoticeFollowedBody
}

type NoticeID int
type NoticeType int
type NoticeIsRead bool
type NoticeFavoriteHeader string
type NoticeFavoriteBody string
type NoticeFollowedBody string

const (
	Favorite NoticeType = iota
	Followed
)

func (t NoticeType) String() string {
	switch t {
	case Favorite:
		return "favorite"
	case Followed:
		return "followed"
	default:
		return ""
	}
}

const (
	FavoriteNoticeTitle = "%s likes: %s"
	FollowedNoticeBody  = "%s follows you"
)

func NewFavoriteNotice(
	p *config.I18nConfig, userID user.UserID, u *user.User, c *content.Content, body content.CommentBody,
) *Notice {
	return &Notice{
		UserID: userID,
		Type:   Favorite,
		Favorite: NoticeFavorite{
			// @TODO FavoriteID
			UserID:              u.ID,
			UserThumbnailURL:    u.ThumbnailURL,
			Header:              NoticeFavoriteHeader(p.Sprintf(FavoriteNoticeTitle, u.Nickname, c.Title)),
			Body:                NoticeFavoriteBody(body),
			ContentID:           c.ID,
			ContentThumbnailURL: c.ThumbnailURL,
		},
	}
}

func NewFollowedNotice(p *config.I18nConfig, followerID user.FollowerID, u *user.User) *Notice {
	return &Notice{
		UserID: user.UserID(followerID),
		Type:   Followed,
		Followed: NoticeFollowed{
			UserID:           u.ID,
			UserThumbnailURL: u.ThumbnailURL,
			Body:             NoticeFollowedBody(p.Sprintf(FollowedNoticeBody, u.Nickname)),
		},
	}
}
