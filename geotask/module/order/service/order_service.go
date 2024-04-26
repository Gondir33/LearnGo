package service

import (
	"context"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"gitlab.com/ptflp/geotask/geo"
	"gitlab.com/ptflp/geotask/module/order/models"
	"gitlab.com/ptflp/geotask/module/order/storage"
)

const (
	minDeliveryPrice = 100.00
	maxDeliveryPrice = 500.00

	maxOrderPrice = 3000.00
	minOrderPrice = 1000.00

	orderMaxAge = 2 * time.Minute
)

type Orderer interface {
	GetByRadius(ctx context.Context, lng, lat, radius float64, unit string) ([]models.Order, error) // возвращает заказы через метод storage.GetByRadius
	Save(ctx context.Context, order models.Order) error                                             // сохраняет заказ через метод storage.Save с заданным временем жизни OrderMaxAge
	GetCount(ctx context.Context) (int, error)                                                      // возвращает количество заказов через метод storage.GetCount
	RemoveOldOrders(ctx context.Context) error                                                      // удаляет старые заказы через метод storage.RemoveOldOrders с заданным временем жизни OrderMaxAge
	GenerateOrder(ctx context.Context) error                                                        // генерирует заказ в случайной точке из разрешенной зоны, с уникальным id, ценой и ценой доставки
}

// OrderService реализация интерфейса Orderer
// в нем должны быть методы GetByRadius, Save, GetCount, RemoveOldOrders, GenerateOrder
// данный сервис отвечает за работу с заказами
type OrderService struct {
	storage       storage.OrderStorager
	allowedZone   geo.PolygonChecker
	disabledZones []geo.PolygonChecker
}

func NewOrderService(storage storage.OrderStorager, allowedZone geo.PolygonChecker, disallowedZone []geo.PolygonChecker) Orderer {
	return &OrderService{storage: storage, allowedZone: allowedZone, disabledZones: disallowedZone}
}

func (s *OrderService) GetByRadius(ctx context.Context, lng, lat, radius float64, unit string) ([]models.Order, error) {
	return s.storage.GetByRadius(ctx, lng, lat, radius, unit)
}
func (s *OrderService) Save(ctx context.Context, order models.Order) error {
	return s.storage.Save(ctx, order, orderMaxAge)
}
func (s *OrderService) GetCount(ctx context.Context) (int, error) {
	return s.storage.GetCount(ctx)
}
func (s *OrderService) RemoveOldOrders(ctx context.Context) error {
	return s.storage.RemoveOldOrders(ctx, orderMaxAge)
}
func (s *OrderService) GenerateOrder(ctx context.Context) error {
	point := geo.GetRandomAllowedLocation(s.allowedZone, s.disabledZones)
	id, err := s.storage.GenerateUniqueID(ctx)
	if err != nil {
		return err
	}

	order := models.Order{
		ID:            id,
		Price:         gofakeit.Float64Range(minOrderPrice, maxOrderPrice),
		DeliveryPrice: gofakeit.Float64Range(minDeliveryPrice, maxDeliveryPrice),
		Lng:           point.Lng,
		Lat:           point.Lat,
		CreatedAt:     time.Now(),
	}
	err = s.Save(ctx, order)
	if err != nil {
		return err
	}
	return nil
}
