package usecase

import (
	"context"
	"errors"

	"github.com/ayupov-ayaz/todo/pkg/modules/relations"

	"github.com/ayupov-ayaz/todo/pkg/modules/item"

	"github.com/ayupov-ayaz/todo/pkg/services/validator"

	"github.com/ayupov-ayaz/todo/internal/models"
	"go.uber.org/zap"
)

type Service struct {
	itemRepo item.Repository
	listRepo item.UsersListRepository
	validate validator.Validator
	logger   *zap.Logger
}

func NewUseCase(repo item.Repository, listRepo item.UsersListRepository, validate validator.Validator) *Service {
	return &Service{
		itemRepo: repo,
		listRepo: listRepo,
		validate: validate,
		logger:   zap.L().Named("item_service"),
	}
}

func (s *Service) checkListOwner(ctx context.Context, userID, listID int) error {
	userList, err := s.listRepo.GetListUserRelationByListId(ctx, listID)
	if err != nil {
		s.logger.Error("get relation users_lists failed",
			zap.Int("list_id", listID),
			zap.Error(err))

		return err
	}

	if userList.UserID != userID {
		s.logger.Error("list doesn't belongs this user",
			zap.Int("list_id", listID),
			zap.Int("user_id", userID))

		return item.ErrListDoesntBelongsUser
	}

	return nil
}

func (s *Service) checkItemOwner(ctx context.Context, userID, itemID int) error {
	userList, err := s.listRepo.GetListUserRelationByItemId(ctx, itemID)
	if err != nil {
		if errors.Is(err, relations.ErrListNotFound) {
			err = item.ErrItemNotFound
		}

		s.logger.Error("get relation users_lists failed",
			zap.Int("item_id", itemID),
			zap.Error(err))

		return err
	}

	if userList.UserID != userID {
		s.logger.Error("item doesn't belongs this user",
			zap.Int("item_id", itemID),
			zap.Int("user_id", userID))

		return item.ErrItemDoesntBelongUser
	}

	return nil
}

func (s *Service) Create(ctx context.Context, userID, listID int, item models.Item) (int, error) {
	if err := s.validate.Struct(item); err != nil {
		s.logger.Warn("validation failed", zap.Error(err))
		return 0, err
	}

	if err := s.checkListOwner(ctx, userID, listID); err != nil {
		return 0, err
	}

	id, err := s.itemRepo.Create(ctx, listID, item)
	if err != nil {
		s.logger.Error("create list item failed", zap.Error(err))
		return 0, err
	}

	return id, nil
}

func (s *Service) GetAll(ctx context.Context, userID, listID int) ([]models.Item, error) {
	if err := s.checkListOwner(ctx, userID, listID); err != nil {
		return nil, err
	}

	items, err := s.itemRepo.GetAll(ctx, listID)
	if err != nil {
		s.logger.Error("get all items failed",
			zap.Int("user_id", userID),
			zap.Int("list_id", listID),
			zap.Error(err))

		return nil, err
	}

	return items, nil
}

func (s *Service) Get(ctx context.Context, userID, itemID int) (models.Item, error) {
	item, err := s.itemRepo.Get(ctx, userID, itemID)
	if err != nil {
		s.logger.Error("get item by id failed",
			zap.Int("item_id", itemID),
			zap.Error(err))

		return models.Item{}, err
	}

	return item, nil
}

func (s *Service) Delete(ctx context.Context, userID, itemID int) error {
	if err := s.itemRepo.Delete(ctx, userID, itemID); err != nil {
		s.logger.Error("delete item failed",
			zap.Int("user_id", userID),
			zap.Int("item_id", itemID),
			zap.Error(err))

		return err
	}

	return nil
}
