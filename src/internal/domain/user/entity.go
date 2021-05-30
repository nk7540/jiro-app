package user

// User entity
type User struct {
	ID             string
	Nickname       string
	Email          string
	Password       string
	FollowingCount int
	FollowerCount  int
}
