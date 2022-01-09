package models

type Item struct {
	ID          int    `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description,omitempty"`
	Done        bool   `json:"done"`
}
