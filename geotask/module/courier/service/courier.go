package service

import (
	"context"
	"math"

	"github.com/go-redis/redis"
	"gitlab.com/ptflp/geotask/geo"
	"gitlab.com/ptflp/geotask/module/courier/models"
	"gitlab.com/ptflp/geotask/module/courier/storage"
)

// Направления движения курьера
const (
	DirectionUp    = 0
	DirectionDown  = 1
	DirectionLeft  = 2
	DirectionRight = 3
)

const (
	DefaultCourierLat = 59.9311
	DefaultCourierLng = 30.3609
)

type Courierer interface {
	GetCourier(ctx context.Context) (*models.Courier, error)
	MoveCourier(courier models.Courier, direction, zoom int) error
}

type CourierService struct {
	courierStorage storage.CourierStorager
	allowedZone    geo.PolygonChecker
	disabledZones  []geo.PolygonChecker
}

func NewCourierService(courierStorage storage.CourierStorager, allowedZone geo.PolygonChecker, disbledZones []geo.PolygonChecker) Courierer {
	return &CourierService{courierStorage: courierStorage, allowedZone: allowedZone, disabledZones: disbledZones}
}

func (c *CourierService) GetCourier(ctx context.Context) (*models.Courier, error) {
	// получаем курьера из хранилища используя метод GetOne из storage/courier.go
	courier, err := c.courierStorage.GetOne(ctx)
	if err == redis.Nil {
		err = c.courierStorage.Save(ctx, models.Courier{
			Score: 0,
			Location: models.Point{
				Lat: DefaultCourierLat,
				Lng: DefaultCourierLng,
			},
		})
		if err != nil {
			return nil, err
		}
		courier, err = c.courierStorage.GetOne(ctx)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	// проверяем, что курьер находится в разрешенной зоне
	if geo.CheckPointIsAllowed(geo.Point(courier.Location), c.allowedZone, c.disabledZones) == false {
		// если нет, то перемещаем его в случайную точку в разрешенной зоне
		newpoint := models.Point(geo.GetRandomAllowedLocation(c.allowedZone, c.disabledZones))
		// сохраняем новые координаты курьера
		courier.Location = newpoint
	}

	return courier, nil
}

// MoveCourier : direction - направление движения курьера, zoom - зум карты
func (c *CourierService) MoveCourier(courier models.Courier, direction, zoom int) error {
	// точность перемещения зависит от зума карты использовать формулу 0.001 / 2^(zoom - 14)
	// 14 - это максимальный зум карты
	if DirectionUp == direction {
		courier.Location.Lng += 0.001 / math.Pow(2, float64(zoom)-14)
	} else if DirectionDown == direction {
		courier.Location.Lng -= 0.001 / math.Pow(2, float64(zoom)-14)
	} else if DirectionRight == direction {
		courier.Location.Lat += 0.001 / math.Pow(2, float64(zoom)-14)
	} else {
		courier.Location.Lat -= 0.001 / math.Pow(2, float64(zoom)-14)
	}
	// далее нужно проверить, что курьер не вышел за границы зоны
	if geo.CheckPointIsAllowed(geo.Point(courier.Location), c.allowedZone, c.disabledZones) == false {
		// если вышел, то нужно переместить его в случайную точку внутри зоны
		newpoint := models.Point(geo.GetRandomAllowedLocation(c.allowedZone, c.disabledZones))
		courier.Location = newpoint
	}
	// далее сохранить изменения в хранилище

	c.courierStorage.Save(context.TODO(), courier)
	return nil
}
