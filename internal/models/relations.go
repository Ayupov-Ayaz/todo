package models

type ListUser struct {
	ID     int
	UserID int
	ListID int
}

type ListItem struct {
	ID     int
	ListID int
	ItemID int
}
