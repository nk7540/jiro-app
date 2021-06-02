package content

// Content entity
type Content struct {
	ID         int
	UserID     int
	CategoryID int
	Title      string `validate:"required,max=256"`
}
