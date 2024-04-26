package service

import (
	"context"
	"petstore/internal/infrastructure/db/dao"
	"petstore/internal/models"
	"petstore/internal/modules/store/storage"
	"strconv"

	"github.com/go-chi/jwtauth"
)

type Storere interface {
	Create(order models.Order) (models.Order, error)
	Find(orderId int) (models.Order, error)
	Delete(orderId int) error
	GetInventory() (map[string]int, error)
}

type Store struct {
	token *jwtauth.JWTAuth
	storage.StoreRepository
}

func fromOrderToOrderDB(order models.Order) models.OrderDB {
	return models.OrderDB{
		Id:       order.Id,
		PetId:    order.PetId,
		Quantity: order.Quantity,
		ShipDate: order.ShipDate,
		Status:   order.Status,
		Complete: order.Complete,
	}
}

func fromOrderDBToOrder(order models.OrderDB) models.Order {
	return models.Order{
		Id:       order.Id,
		PetId:    order.PetId,
		Quantity: order.Quantity,
		ShipDate: order.ShipDate,
		Status:   order.Status,
		Complete: order.Complete,
	}
}
func NewStoreService(token *jwtauth.JWTAuth, StoreRep storage.StoreRepository) Storere {
	return &Store{
		token:           token,
		StoreRepository: StoreRep,
	}
}

func (s *Store) Create(order models.Order) (models.Order, error) {
	orderDb := fromOrderToOrderDB(order)
	err := s.StoreRepository.Create(context.Background(), orderDb)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}
func (s *Store) Find(orderId int) (models.Order, error) {
	orederDb, err := s.StoreRepository.GetByID(context.Background(), strconv.Itoa(orderId))
	if err != nil {
		return models.Order{}, err
	}
	return fromOrderDBToOrder(orederDb), nil
}
func (s *Store) Delete(orderId int) error {
	return s.StoreRepository.Delete(context.Background(), strconv.Itoa(orderId))
}
func (s *Store) GetInventory() (map[string]int, error) {
	mapStatus := make(map[string]int, 10)

	arrOrdersDb, err := s.StoreRepository.List(context.Background(), dao.Condition{
		LimitOffset: &dao.LimitOffset{
			Limit: 10000,
		},
	})
	if err != nil {
		return map[string]int{}, err
	}
	for i, orderDb := range arrOrdersDb {
		if orderDb.Status == "placed" {
			mapStatus["additionalProp"+strconv.Itoa(i)] = 1
		} else if orderDb.Status == "approved" {
			mapStatus["additionalProp"+strconv.Itoa(i)] = 2
		} else {
			mapStatus["additionalProp"+strconv.Itoa(i)] = 3
		}
	}
	return mapStatus, nil
}
