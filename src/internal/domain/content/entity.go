package content

// Content entity
type Content struct {
	UserID     string
	CategoryID string
	Title      string `validate:"required,max=256"`
}
