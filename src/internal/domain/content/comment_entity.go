package content

import "artics-api/src/internal/domain/user"

type Comment struct {
	ID        CommentID
	UserID    user.UserID
	ContentID ContentID
	Body      CommentBody
}

type CommentID int
type CommentBody string
