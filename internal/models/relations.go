package models

type ListUser struct {
	ID     int `db:"id"`
	UserID int `db:"user_id"`
	ListID int `db:"list_id"`
}

type ListItem struct {
	ID     int
	ListID int
	ItemID int
}
