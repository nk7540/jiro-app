package content

// Favorite entity
type Favorite struct {
	ID        FavoriteID
	UserID    FavoriteUserID
	ContentID FavoriteContentID
}

type FavoriteID int
type FavoriteUserID int
type FavoriteContentID int
