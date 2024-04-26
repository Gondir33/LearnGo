package service

import (
	"context"
	"log"

	cservice "gitlab.com/ptflp/geotask/module/courier/service"
	cfm "gitlab.com/ptflp/geotask/module/courierfacade/models"
	oservice "gitlab.com/ptflp/geotask/module/order/service"
)

const (
	CourierVisibilityRadius = 2800 // 2.8km
)

type CourierFacer interface {
	MoveCourier(ctx context.Context, direction, zoom int) // отвечает за движение курьера по карте direction - направление движения, zoom - уровень зума
	GetStatus(ctx context.Context) cfm.CourierStatus      // отвечает за получение статуса курьера и заказов вокруг него
}

// CourierFacade фасад для курьера и заказов вокруг него (для фронта)
type CourierFacade struct {
	courierService cservice.Courierer
	orderService   oservice.Orderer
}

func NewCourierFacade(courierService cservice.Courierer, orderService oservice.Orderer) CourierFacer {
	return &CourierFacade{courierService: courierService, orderService: orderService}
}

func (c *CourierFacade) MoveCourier(ctx context.Context, direction, zoom int) {
	courier, err := c.courierService.GetCourier(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	err = c.courierService.MoveCourier(*courier, direction, zoom)
	if err != nil {
		log.Println(err)
	}
}
func (c *CourierFacade) GetStatus(ctx context.Context) cfm.CourierStatus {
	courier, err := c.courierService.GetCourier(ctx)
	if err != nil {
		log.Println(err)
		return cfm.CourierStatus{}
	}
	orders, err := c.orderService.GetByRadius(ctx, courier.Location.Lng, courier.Location.Lat, CourierVisibilityRadius, "km")
	if err != nil {
		log.Println(err)
		return cfm.CourierStatus{}
	}
	return cfm.CourierStatus{
		Courier: *courier,
		Orders:  orders,
	}
}
