package storage

import (
	"context"
	"repository/internal/infrastructure/db/dao"
	"repository/internal/models"
	"strconv"
	"time"
)

type (
	UserStorage struct {
		daoAdapter dao.IfaceDAO
	}

	UserRepository interface {
		Create(ctx context.Context, user models.User) error
		GetByID(ctx context.Context, id string) (models.User, error)
		Update(ctx context.Context, user models.User) error
		Delete(ctx context.Context, id string) error
		List(ctx context.Context, c dao.Condition) ([]models.User, error)
		// Другие методы, необходимые для работы с пользователями
	}
)

func NewUserStorage(db dao.IfaceDAO) UserRepository {
	return &UserStorage{db}
}

func (s *UserStorage) Create(ctx context.Context, user models.User) error {
	return s.daoAdapter.Create(ctx, &user)
}

func (s *UserStorage) GetByID(ctx context.Context, id string) (models.User, error) {
	var list []models.User
	user := models.User{}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return models.User{}, err
	}
	err = s.daoAdapter.List(ctx, &list, &user, dao.Condition{
		Equal: map[string]interface{}{"id": idInt},
		LimitOffset: &dao.LimitOffset{
			Limit:  3,
			Offset: 0,
		},
	})
	if err != nil {
		return models.User{}, err
	}

	return list[0], nil
}

func (s *UserStorage) Update(ctx context.Context, user models.User) error {
	err := s.daoAdapter.Update(ctx, &user, dao.Condition{
		Equal: map[string]interface{}{"id": user.Id},
	})
	return err
}

func (s *UserStorage) Delete(ctx context.Context, id string) error {
	user, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}

	user.DeletedAt = time.Now().Format(time.DateTime)

	return s.Update(ctx, user)
}

func (s *UserStorage) List(ctx context.Context, c dao.Condition) ([]models.User, error) {
	var list []models.User
	user := models.User{}
	err := s.daoAdapter.List(ctx, &list, &user, c)
	if err != nil {
		return []models.User{}, err
	}
	return list, nil
}
