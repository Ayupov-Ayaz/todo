package list

import (
	"context"

	"github.com/ayupov-ayaz/todo/internal/models"
)

type Repository interface {
	Create(ctx context.Context, userID int, list models.TodoList) (int, error)
	GetAll(ctx context.Context, userID int) ([]models.TodoList, error)
	Get(ctx context.Context, userID int, listID int) (models.TodoList, error)
	Update(ctx context.Context, userID, listID int, input UpdateTodoList) error
	Delete(ctx context.Context, userID, listID int) error
}

type UseCase interface {
	Create(ctx context.Context, userID int, list models.TodoList) (int, error)
	GetAll(ctx context.Context, userID int) ([]models.TodoList, error)
	Get(ctx context.Context, userID, listID int) (models.TodoList, error)
	Update(ctx context.Context, userID, listID int, list UpdateTodoList) error
	Delete(ctx context.Context, userID, listID int) error
}
