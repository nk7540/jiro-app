package content

import (
	"artics-api/src/config"
	"artics-api/src/internal/domain/user"
)

type Notice struct {
	ID     NoticeID
	UserID user.UserID
	Title  NoticeTitle
	Body   NoticeBody
}

type NoticeID int
type NoticeTitle string
type NoticeBody string

const (
	FavoriteNoticeTitle = "%s likes: %s"
	FollowedNoticeBody  = "%s follows you"
)

func NewFavoriteNotice(
	p *config.I18nConfig, userID user.UserID, nickname user.Nickname, title Title, body CommentBody,
) *Notice {
	return &Notice{
		UserID: userID,
		Title:  NoticeTitle(p.Sprintf(FavoriteNoticeTitle, nickname, title)),
		Body:   NoticeBody(body),
	}
}

func NewFollowedNotice(p *config.I18nConfig, followerID user.FollowerID, nickname user.Nickname) *Notice {
	return &Notice{
		UserID: user.UserID(followerID),
		Title:  "",
		Body:   NoticeBody(p.Sprintf(FollowedNoticeBody, nickname)),
	}
}
