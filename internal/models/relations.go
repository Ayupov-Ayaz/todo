package models

type ListUserRelation struct {
	ID     int
	UserID int `db:"user_id"`
	ListID int `db:"list_id"`
}

type ListItem struct {
	ID     int
	ListID int
	ItemID int
}
