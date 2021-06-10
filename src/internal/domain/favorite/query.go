package favorite

type QueryFavorite struct {
	ID        int `boil:"id"`
	UserID    int `boil:"user_id"`
	ContentID int `boil:"content_id"`
}
