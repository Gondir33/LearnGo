package storage

import (
	"context"
	"petstore/internal/infrastructure/db/dao"
	"petstore/internal/models"
	"strconv"
	"time"
)

type StoreRepository interface {
	Create(ctx context.Context, orderDb models.OrderDB) error
	GetByID(ctx context.Context, id string) (models.OrderDB, error)
	Update(ctx context.Context, orderDb models.OrderDB) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, c dao.Condition) ([]models.OrderDB, error)
}
type StoreStorage struct {
	daoAdapter dao.IfaceDAO
}

func NewStoreStorage(db dao.IfaceDAO) StoreRepository {
	return &StoreStorage{db}
}

func (s *StoreStorage) Create(ctx context.Context, orderDb models.OrderDB) error {
	orderDb.CreatedAt = time.Now().Format(time.DateTime)

	return s.daoAdapter.Create(ctx, &orderDb)
}

func (s *StoreStorage) GetByID(ctx context.Context, id string) (models.OrderDB, error) {
	var list []models.OrderDB
	orderDb := models.OrderDB{}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return models.OrderDB{}, err
	}
	err = s.daoAdapter.List(ctx, &list, &orderDb, dao.Condition{
		Equal: map[string]interface{}{"id": idInt},
		LimitOffset: &dao.LimitOffset{
			Limit:  3,
			Offset: 0,
		},
	})
	if err != nil {
		return models.OrderDB{}, err
	}

	return list[0], nil
}

func (s *StoreStorage) Update(ctx context.Context, orderDb models.OrderDB) error {
	err := s.daoAdapter.Update(ctx, &orderDb, dao.Condition{
		Equal: map[string]interface{}{"id": orderDb.Id},
	})
	return err
}

func (s *StoreStorage) Delete(ctx context.Context, id string) error {
	orderDb, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}

	orderDb.DeletedAt = time.Now().Format(time.DateTime)

	return s.Update(ctx, orderDb)
}

func (s *StoreStorage) List(ctx context.Context, c dao.Condition) ([]models.OrderDB, error) {
	var list []models.OrderDB
	orderDb := models.OrderDB{}
	err := s.daoAdapter.List(ctx, &list, &orderDb, c)
	if err != nil {
		return []models.OrderDB{}, err
	}
	return list, nil
}
