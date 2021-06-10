package content

type Browse struct {
	ID        BrowseID
	UserID    BrowseUserID
	ContentID BrowseContentID
}

type BrowseID int
type BrowseUserID int
type BrowseContentID int
