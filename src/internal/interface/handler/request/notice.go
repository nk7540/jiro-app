package request

type Listen struct {
	close string `json:"close"`
}

type ListNotices struct {
	Page int `json:"page"`
	Per  int `json:"per"`
}
