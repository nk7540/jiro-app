package response

type Contents struct {
	Contents []*Content `json:"contents"`
}

type Content struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
