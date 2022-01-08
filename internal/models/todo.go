package models

type TodoList struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" validate:"required"`
	Description string `json:"description,omitempty" db:"description"`
}
