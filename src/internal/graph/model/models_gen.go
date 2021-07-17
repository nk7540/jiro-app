// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

type Notice interface {
	IsNotice()
}

type Content struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	ThumbnailURL string `json:"thumbnailUrl"`
	Description  string `json:"description"`
	Author       string `json:"author"`
	IsLiked      *bool  `json:"isLiked"`
}

type CreateContent struct {
	Title       string          `json:"title"`
	Thumbnail   *graphql.Upload `json:"thumbnail"`
	Description *string         `json:"description"`
	Author      *string         `json:"author"`
}

type CreateUser struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type Favorite struct {
	ID          int      `json:"id"`
	User        *User    `json:"user"`
	Content     *Content `json:"content"`
	CommentBody string   `json:"commentBody"`
}

type Like struct {
	ContentID    int    `json:"contentId"`
	ToUserIds    []int  `json:"toUserIds"`
	ToCloseUsers bool   `json:"toCloseUsers"`
	CommentBody  string `json:"commentBody"`
}

type NoticeFavorite struct {
	ID         int      `json:"id"`
	Type       string   `json:"type"`
	IsRead     bool     `json:"isRead"`
	CreatedAt  string   `json:"createdAt"`
	FavoriteID int      `json:"favoriteId"`
	User       *User    `json:"user"`
	Header     string   `json:"header"`
	Body       string   `json:"body"`
	Content    *Content `json:"content"`
}

func (NoticeFavorite) IsNotice() {}

type NoticeFollowed struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	IsRead    bool   `json:"isRead"`
	CreatedAt string `json:"createdAt"`
	User      *User  `json:"user"`
	Body      string `json:"body"`
}

func (NoticeFollowed) IsNotice() {}

type UpdateContent struct {
	ID          int             `json:"id"`
	Title       *string         `json:"title"`
	Thumbnail   *graphql.Upload `json:"thumbnail"`
	Description *string         `json:"description"`
	Author      *string         `json:"author"`
}

type UpdateUser struct {
	Nickname  *string         `json:"nickname"`
	Thumbnail *graphql.Upload `json:"thumbnail"`
	Profile   *string         `json:"profile"`
}

type User struct {
	ID           int     `json:"id"`
	Email        *string `json:"email"`
	Nickname     *string `json:"nickname"`
	ThumbnailURL *string `json:"thumbnailUrl"`
	Profile      *string `json:"profile"`
	IsClose      *bool   `json:"isClose"`
	IsFollowed   *bool   `json:"isFollowed"`
	IsFollowing  *bool   `json:"isFollowing"`
}

type ContentKind string

const (
	ContentKindRecommended ContentKind = "RECOMMENDED"
	ContentKindBrowsed     ContentKind = "BROWSED"
	ContentKindUploaded    ContentKind = "UPLOADED"
)

var AllContentKind = []ContentKind{
	ContentKindRecommended,
	ContentKindBrowsed,
	ContentKindUploaded,
}

func (e ContentKind) IsValid() bool {
	switch e {
	case ContentKindRecommended, ContentKindBrowsed, ContentKindUploaded:
		return true
	}
	return false
}

func (e ContentKind) String() string {
	return string(e)
}

func (e *ContentKind) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ContentKind(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContentKind", str)
	}
	return nil
}

func (e ContentKind) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FavoriteKind string

const (
	FavoriteKindCommon             FavoriteKind = "COMMON"
	FavoriteKindMost               FavoriteKind = "MOST"
	FavoriteKindOthers             FavoriteKind = "OTHERS"
	FavoriteKindUserFavoriteForYou FavoriteKind = "USER_FAVORITE_FOR_YOU"
	FavoriteKindUserFavorite       FavoriteKind = "USER_FAVORITE"
)

var AllFavoriteKind = []FavoriteKind{
	FavoriteKindCommon,
	FavoriteKindMost,
	FavoriteKindOthers,
	FavoriteKindUserFavoriteForYou,
	FavoriteKindUserFavorite,
}

func (e FavoriteKind) IsValid() bool {
	switch e {
	case FavoriteKindCommon, FavoriteKindMost, FavoriteKindOthers, FavoriteKindUserFavoriteForYou, FavoriteKindUserFavorite:
		return true
	}
	return false
}

func (e FavoriteKind) String() string {
	return string(e)
}

func (e *FavoriteKind) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FavoriteKind(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FavoriteKind", str)
	}
	return nil
}

func (e FavoriteKind) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserKind string

const (
	UserKindFollowing UserKind = "FOLLOWING"
	UserKindFollower  UserKind = "FOLLOWER"
)

var AllUserKind = []UserKind{
	UserKindFollowing,
	UserKindFollower,
}

func (e UserKind) IsValid() bool {
	switch e {
	case UserKindFollowing, UserKindFollower:
		return true
	}
	return false
}

func (e UserKind) String() string {
	return string(e)
}

func (e *UserKind) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserKind(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserKind", str)
	}
	return nil
}

func (e UserKind) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
