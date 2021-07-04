package content

import "artics-api/src/internal/domain/user"

// Favorite entity
type Favorite struct {
	ID           FavoriteID
	UserID       user.UserID
	ContentID    ContentID
	CommentID    CommentID
	ToUserIDs    FavoriteToUserIDs
	ToCloseUsers FavoriteToCloseUsers
}

type FavoriteID int
type FavoriteUserID int
type FavoriteContentID int
type FavoriteCommentID int
type FavoriteToUserIDs []user.UserID
type FavoriteToCloseUsers bool
