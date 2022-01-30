package item

import (
	"context"

	"github.com/ayupov-ayaz/todo/internal/models"
)

type Repository interface {
	Create(ctx context.Context, listID int, item models.Item) (int, error)
	GetAll(ctx context.Context, listID int) ([]models.Item, error)
	Get(ctx context.Context, userID, itemID int) (models.Item, error)
	Delete(ctx context.Context, userID, itemID int) error
}

type UsersListRepository interface {
	GetListUserByListId(ctx context.Context, listID int) (models.ListUser, error)
	GetListUserByItemId(ctx context.Context, itemID int) (models.ListUser, error)
}

type UseCase interface {
	Create(ctx context.Context, userID, listID int, item models.Item) (int, error)
	GetAll(ctx context.Context, userID, listID int) ([]models.Item, error)
	Get(ctx context.Context, userID, itemID int) (models.Item, error)
	Delete(ctx context.Context, userID, itemID int) error
}
