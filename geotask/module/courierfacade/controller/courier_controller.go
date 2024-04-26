package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/ptflp/geotask/module/courierfacade/service"
)

type CourierController struct {
	courierService service.CourierFacer
}

func NewCourierController(courierService service.CourierFacer) *CourierController {
	return &CourierController{courierService: courierService}
}

// @Summary	Получение статуса сервиса
// @Tags		status
// @Accept		json
// @Produce	json
// @Success	200	{object}	models.CourierStatus
// @Router		/api/status [get]
func (c *CourierController) GetStatus(ctx *gin.Context) {
	// установить задержку в 50 миллисекунд
	time.Sleep(50 * time.Millisecond)
	// получить статус курьера из сервиса courierService используя метод GetStatus
	status := c.courierService.GetStatus(ctx)
	// отправить статус курьера в ответ
	ctx.JSON(http.StatusOK, status)
}

func (c *CourierController) MoveCourier(m webSocketMessage) {
	var cm CourierMove
	// получить данные из m.Data и десериализовать их в структуру CourierMove

	err := json.Unmarshal(m.Data.([]byte), &cm)
	if err != nil {
		log.Println(err)
		return
	}
	// вызвать метод MoveCourier у courierService
	c.courierService.MoveCourier(context.Background(), cm.Direction, cm.Zoom)
}
