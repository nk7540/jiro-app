package content

// Content entity
type Content struct {
	ID         string
	UserID     string
	CategoryID string
	Title      string `validate:"required,max=256"`
}
