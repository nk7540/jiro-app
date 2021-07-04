package response

type Contents struct {
	Contents []*Content `json:"contents"`
}

type Content struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Notice struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
