package order

import (
	"context"
	"log"
	"time"

	"gitlab.com/ptflp/geotask/module/order/service"
)

const (
	orderCleanInterval = 5 * time.Second
)

// OrderCleaner воркер, который удаляет старые заказы
// используя метод orderService.RemoveOldOrders()
type OrderCleaner struct {
	orderService service.Orderer
}

func NewOrderCleaner(orderService service.Orderer) *OrderCleaner {
	return &OrderCleaner{orderService: orderService}
}

func (o *OrderCleaner) Run() {
	// исользовать горутину и select
	// внутри горутины нужно использовать time.NewTicker()
	// и вызывать метод orderService.RemoveOldOrders()
	// если при удалении заказов произошла ошибка, то нужно вывести ее в лог
	go func() {
		ticker := time.NewTicker(orderCleanInterval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				err := o.orderService.RemoveOldOrders(context.TODO())
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()
}
