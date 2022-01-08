package list

import (
	"context"

	"go.uber.org/zap"

	"github.com/ayupov-ayaz/todo/pkg/services/validator"

	"github.com/ayupov-ayaz/todo/internal/models"
)

type TodoListRepository interface {
	Create(ctx context.Context, userID int, list models.TodoList) (int, error)
}

type Service struct {
	repo     TodoListRepository
	validate validator.Validator
	logger   *zap.Logger
}

func NewService(repo TodoListRepository, validator validator.Validator) *Service {
	return &Service{
		repo:     repo,
		validate: validator,
		logger:   zap.L().Named("list_service"),
	}
}

func (s *Service) Create(ctx context.Context, userID int, list models.TodoList) (int, error) {
	if err := s.validate.Struct(list); err != nil {
		s.logger.Warn("validation failed", zap.Error(err))
		return 0, err
	}

	id, err := s.repo.Create(ctx, userID, list)
	if err != nil {
		s.logger.Error("create list failed", zap.Error(err))
		return 0, err
	}

	return id, nil
}
