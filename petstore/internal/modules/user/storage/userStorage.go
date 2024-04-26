package storage

import (
	"context"
	"errors"
	"petstore/internal/infrastructure/db/dao"
	"petstore/internal/models"
	"strconv"
	"time"
)

type UsererRepository interface {
	Create(ctx context.Context, userDb models.UserDB) error
	GetByID(ctx context.Context, id string) (models.UserDB, error)
	Update(ctx context.Context, userDb models.UserDB) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, c dao.Condition) ([]models.UserDB, error)
	GetByUsername(ctx context.Context, username string) (models.UserDB, error)
	PutApiKey(ctx context.Context, api_key models.Api_keyDB) error
}
type UserStorage struct {
	daoAdapter dao.IfaceDAO
}

func NewUserStorage(db dao.IfaceDAO) UsererRepository {
	return &UserStorage{db}
}

func (s *UserStorage) Create(ctx context.Context, userDb models.UserDB) error {
	userDb.CreatedAt = time.Now().Format(time.DateTime)

	return s.daoAdapter.Create(ctx, &userDb)
}
func (s *UserStorage) GetByUsername(ctx context.Context, username string) (models.UserDB, error) {
	var list []models.UserDB
	userDb := models.UserDB{}
	err := s.daoAdapter.List(ctx, &list, &userDb, dao.Condition{
		Equal: map[string]interface{}{"username": username},
		LimitOffset: &dao.LimitOffset{
			Limit:  3,
			Offset: 0,
		},
	})
	if err != nil {
		return models.UserDB{}, err
	}
	if len(list) == 0 {
		return models.UserDB{}, errors.New("no such username")
	}
	return list[0], nil
}
func (s *UserStorage) GetByID(ctx context.Context, id string) (models.UserDB, error) {
	var list []models.UserDB
	userDb := models.UserDB{}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return models.UserDB{}, err
	}
	err = s.daoAdapter.List(ctx, &list, &userDb, dao.Condition{
		Equal: map[string]interface{}{"id": idInt},
		LimitOffset: &dao.LimitOffset{
			Limit:  3,
			Offset: 0,
		},
	})
	if err != nil {
		return models.UserDB{}, err
	}

	return list[0], nil
}

func (s *UserStorage) Update(ctx context.Context, userDb models.UserDB) error {
	err := s.daoAdapter.Update(ctx, &userDb, dao.Condition{
		Equal: map[string]interface{}{"id": userDb.Id},
	})
	return err
}

func (s *UserStorage) Delete(ctx context.Context, id string) error {
	userDb, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}

	userDb.DeletedAt = time.Now().Format(time.DateTime)

	return s.Update(ctx, userDb)
}

func (s *UserStorage) List(ctx context.Context, c dao.Condition) ([]models.UserDB, error) {
	var list []models.UserDB
	userDb := models.UserDB{}
	err := s.daoAdapter.List(ctx, &list, &userDb, c)
	if err != nil {
		return []models.UserDB{}, err
	}
	return list, nil
}

func (s UserStorage) PutApiKey(ctx context.Context, api_key models.Api_keyDB) error {
	return s.daoAdapter.Create(ctx, &api_key)
}
