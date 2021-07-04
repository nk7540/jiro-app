package content

import "artics-api/src/internal/domain/user"

type CommandLike struct {
	User         *user.User
	ContentID    ContentID
	ToUserIDs    FavoriteToUserIDs
	ToCloseUsers FavoriteToCloseUsers
	CommentBody  CommentBody
}

type CommandUnlike struct {
	UserID    FavoriteUserID
	ContentID FavoriteContentID
}

type CommandBrowse struct {
	UserID    BrowseUserID
	ContentID BrowseContentID
}
