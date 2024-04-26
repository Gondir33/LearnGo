package storage

import (
	"context"
	"petstore/internal/infrastructure/db/dao"
	"petstore/internal/models"
	"strconv"
	"time"
)

type PeterRepository interface {
	Create(ctx context.Context, petDb models.PetDB) error
	GetByID(ctx context.Context, id string) (models.PetDB, error)
	Update(ctx context.Context, petDb models.PetDB) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, c dao.Condition) ([]models.PetDB, error)
	GetApiKey(ctx context.Context, api_key string) bool
}
type PetStorage struct {
	daoAdapter dao.IfaceDAO
}

func NewPetStorage(db dao.IfaceDAO) PeterRepository {
	return &PetStorage{db}
}

func (s *PetStorage) Create(ctx context.Context, petDb models.PetDB) error {
	petDb.CreatedAt = time.Now().Format(time.DateTime)

	return s.daoAdapter.Create(ctx, &petDb)
}

func (s *PetStorage) GetByID(ctx context.Context, id string) (models.PetDB, error) {
	var list []models.PetDB
	petDb := models.PetDB{}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return models.PetDB{}, err
	}
	err = s.daoAdapter.List(ctx, &list, &petDb, dao.Condition{
		Equal: map[string]interface{}{"id": idInt},
		LimitOffset: &dao.LimitOffset{
			Limit:  3,
			Offset: 0,
		},
	})
	if err != nil {
		return models.PetDB{}, err
	}

	return list[0], nil
}

func (s *PetStorage) Update(ctx context.Context, petDb models.PetDB) error {
	err := s.daoAdapter.Update(ctx, &petDb, dao.Condition{
		Equal: map[string]interface{}{"id": petDb.Id},
	})
	return err
}

func (s *PetStorage) Delete(ctx context.Context, id string) error {
	petDb, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}

	petDb.DeletedAt = time.Now().Format(time.DateTime)

	return s.Update(ctx, petDb)
}

func (s *PetStorage) List(ctx context.Context, c dao.Condition) ([]models.PetDB, error) {
	var list []models.PetDB
	petDb := models.PetDB{}
	err := s.daoAdapter.List(ctx, &list, &petDb, c)
	if err != nil {
		return []models.PetDB{}, err
	}
	return list, nil
}

func (s *PetStorage) GetApiKey(ctx context.Context, api_key string) bool {
	var list []models.Api_keyDB
	s.daoAdapter.List(ctx, &list, &models.Api_keyDB{}, dao.Condition{
		LimitOffset: &dao.LimitOffset{Limit: 3, Offset: 0},
		Equal:       map[string]interface{}{"api_key": api_key},
	})
	if len(list) > 0 {
		return true
	}
	return false
}
