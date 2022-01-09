package item

import (
	"context"

	_errors "github.com/ayupov-ayaz/todo/errors"

	"github.com/ayupov-ayaz/todo/pkg/services/validator"

	"github.com/ayupov-ayaz/todo/internal/models"
	"go.uber.org/zap"
)

var (
	ErrListDoesntBelongsUser = _errors.Forbidden("list doesn't belongs user")
)

type Repository interface {
	Create(ctx context.Context, listID int, item models.Item) (int, error)
}

type UsersListRepository interface {
	GetListUserByListId(ctx context.Context, listID int) (models.ListUser, error)
}

type Service struct {
	itemRepo Repository
	listRepo UsersListRepository
	validate validator.Validator
	logger   *zap.Logger
}

func NewService(repo Repository, listRepo UsersListRepository, validate validator.Validator) *Service {
	return &Service{
		itemRepo: repo,
		listRepo: listRepo,
		validate: validate,
		logger:   zap.L().Named("item_service"),
	}
}

func (s *Service) Create(ctx context.Context, userID, listID int, item models.Item) (int, error) {
	if err := s.validate.Struct(item); err != nil {
		s.logger.Warn("validation failed", zap.Error(err))
		return 0, err
	}

	userList, err := s.listRepo.GetListUserByListId(ctx, listID)
	if err != nil {
		s.logger.Error("get relation users_lists failed",
			zap.Int("list_id", listID),
			zap.Error(err))

		return 0, err
	}

	if userList.UserID != userID {
		s.logger.Error("list doesn't belongs this user",
			zap.Int("list_id", listID),
			zap.Int("user_id", userID))

		return 0, ErrListDoesntBelongsUser
	}

	id, err := s.itemRepo.Create(ctx, listID, item)
	if err != nil {
		s.logger.Error("create list item failed", zap.Error(err))
		return 0, err
	}

	return id, nil
}
