package content

type QueryDetailContent struct {
	ID          int    `boil:"id"`
	UserID      int    `boil:"user_id"`
	CategoryID  int    `boil:"content_id"`
	Title       string `boil:"title"`
	Description string `boil:"description"`
}

type QueryContent struct {
	ID         int
	UserID     int
	CategoryID int
	Title      string
}

type QueryFavorite struct {
	ID        int `boil:"id"`
	UserID    int `boil:"user_id"`
	ContentID int `boil:"content_id"`
}
