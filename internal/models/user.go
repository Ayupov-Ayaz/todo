package models

type User struct {
	ID       int    `json:"-"`
	Name     string `json:"name" validate:"required,min=2"`
	Username string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=6"`
}
