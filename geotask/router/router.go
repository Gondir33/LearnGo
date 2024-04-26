package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/ptflp/geotask/module/courierfacade/controller"
)

type Router struct {
	courier *controller.CourierController
}

func NewRouter(courier *controller.CourierController) *Router {
	return &Router{courier: courier}
}

func (r *Router) CourierAPI(router *gin.RouterGroup) {
	// прописать роуты для courier API
	router.GET("status", r.courier.GetStatus)
	router.GET("ws", r.courier.Websocket)

}

func (r *Router) Swagger(router *gin.RouterGroup) {
	router.GET("/docs/swagger.json", func(c *gin.Context) {
		// Здесь вы можете использовать функцию ServeFile для обслуживания файла
		c.File("/docs/swagger.json")
	})

	router.GET("/swagger", swaggerUI)
}
