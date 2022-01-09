package models

type Item struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" validate:"required"`
	Description string `json:"description,omitempty" db:"description"`
	Done        bool   `json:"done" db:"done"`
}
