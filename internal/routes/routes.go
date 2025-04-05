package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"rest-project/internal/delivery"
	"rest-project/internal/repository"
	"rest-project/internal/services"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := delivery.NewOrderHandler(orderService)

	orders := r.Group("orders")
	{
		orders.GET("/", orderHandler.GetAllOrders)
		orders.POST("/", orderHandler.CreateOrder)
		orders.GET(":id", orderHandler.GetOrder)
		orders.PUT(":id", orderHandler.UpdateOrder)
		orders.DELETE(":id", orderHandler.DeleteOrder)
	}

}
