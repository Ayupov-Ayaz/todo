package list

import (
	"context"

	"go.uber.org/zap"

	"github.com/ayupov-ayaz/todo/pkg/services/validator"

	"github.com/ayupov-ayaz/todo/internal/models"
)

type TodoListRepository interface {
	Create(ctx context.Context, userID int, list models.TodoList) (int, error)
	GetAll(ctx context.Context, userID int) ([]models.TodoList, error)
	Get(ctx context.Context, userID int, listID int) (models.TodoList, error)
	Update(ctx context.Context, userID, listID int, input UpdateTodoList) error
	Delete(ctx context.Context, userID, listID int) error
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

func (s *Service) GetAll(ctx context.Context, userID int) ([]models.TodoList, error) {
	lists, err := s.repo.GetAll(ctx, userID)
	if err != nil {
		s.logger.Error("get lists failed",
			zap.Int("user_id", userID),
			zap.Error(err))

		return nil, err
	}

	return lists, nil
}

func (s *Service) Get(ctx context.Context, userID, listID int) (models.TodoList, error) {
	list, err := s.repo.Get(ctx, userID, listID)
	if err != nil {
		s.logger.Error("get list failed",
			zap.Int("user_id", userID),
			zap.Int("list_id", listID),
			zap.Error(err))

		return models.TodoList{}, err
	}

	return list, nil
}

func (s *Service) Update(ctx context.Context, userID, listID int, input UpdateTodoList) error {
	if err := input.Validate(); err != nil {
		s.logger.Warn("validation failed", zap.Error(err))
		return err
	}

	if err := s.repo.Update(ctx, userID, listID, input); err != nil {
		s.logger.Error("failed to update the list",
			zap.Int("user_id", userID),
			zap.Int("list_id", listID),
			zap.Error(err))

		return err
	}

	return nil
}

func (s *Service) Delete(ctx context.Context, userID, listID int) error {
	if err := s.repo.Delete(ctx, userID, listID); err != nil {
		s.logger.Error("failed to delete the list",
			zap.Int("user_id", userID),
			zap.Int("list_id", listID),
			zap.Error(err))

		return err
	}

	return nil
}
