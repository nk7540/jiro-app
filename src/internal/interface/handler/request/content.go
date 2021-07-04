package request

type Like struct {
	ContentID    int    `json:"contentID"`
	ToUserIDs    []int  `json:"toUserIDs"`
	ToCloseUsers bool   `json:"toCloseUsers"`
	CommentBody  string `json:"commentBody"`
}
