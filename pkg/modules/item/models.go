package item

import "github.com/ayupov-ayaz/todo/internal/models"

type getAllItemsResponse struct {
	Items []models.Item `json:"items"`
}
