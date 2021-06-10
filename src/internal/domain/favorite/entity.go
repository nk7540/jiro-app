package favorite

// Favorite entity
type Favorite struct {
	ID        ID
	UserID    UserID
	ContentID ContentID
}

type ID int
type UserID int
type ContentID int
