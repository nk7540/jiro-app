package content

// Content entity
type Content struct {
	ID           ContentID
	UserID       ContentUserID
	CategoryID   ContentCategoryID
	Title        Title `validate:"required,max=256"`
	ThumbnailURL ThumbnailURL
}

type ContentID int
type ContentUserID int
type ContentCategoryID int
type Title string
type ThumbnailURL string
