package response

type Contents struct {
	Contents []*Content `json:"contents"`
}

type Content struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
