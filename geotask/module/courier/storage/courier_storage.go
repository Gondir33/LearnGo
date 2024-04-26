package storage

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	"gitlab.com/ptflp/geotask/module/courier/models"
)

type CourierStorager interface {
	Save(ctx context.Context, courier models.Courier) error // сохранить курьера по ключу courier
	GetOne(ctx context.Context) (*models.Courier, error)    // получить курьера по ключу courier
}

type CourierStorage struct {
	storage *redis.Client
}

func NewCourierStorage(storage *redis.Client) CourierStorager {
	return &CourierStorage{storage: storage}
}

func (s *CourierStorage) Save(ctx context.Context, courier models.Courier) error {
	bytes, err := json.Marshal(courier)
	if err != nil {
		return err
	}
	cmd := s.storage.Set("courier", string(bytes), 5*time.Minute)
	return cmd.Err()
}
func (s *CourierStorage) GetOne(ctx context.Context) (*models.Courier, error) {
	var courier models.Courier

	bytes, err := s.storage.Get("courier").Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &courier)
	if err != nil {
		return nil, err
	}
	return &courier, nil
}
