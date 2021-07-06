package response

type NoticeFavorite struct {
	ID                  int    `json:"id"`
	Type                string `json:"type"`
	IsRead              bool   `json:"isRead"`
	FavoriteID          int    `json:"favorite_id"`
	UserID              int    `json:"user_id"`
	UserThumbnailURL    string `json:"user_thumbnail_url"`
	Header              string `json:"header"`
	Body                string `json:"body"`
	ContentID           int    `json:"content_id"`
	ContentThumbnailURL string `json:"content_thumbnail_url"`
}

type NoticeFollowed struct {
	ID               int    `json:"id"`
	Type             string `json:"type"`
	IsRead           bool   `json:"isRead"`
	UserID           int    `json:"user_id"`
	UserThumbnailURL string `json:"user_thumbnail_url"`
	Body             string `json:"body"`
}

type Notices struct {
	Notices []interface{} `json:"notices"`
}
