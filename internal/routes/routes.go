package routes

import (
	"github.com/gin-gonic/gin"
	"order-project/internal/auth"
	"order-project/internal/db"
	"order-project/internal/delivery"
	"order-project/internal/middleware"
	"order-project/internal/repository"
	"order-project/internal/services"
)

func SetupRoutes(r *gin.Engine) {
	// Auth routes
	r.POST("/api/v1/auth/login", auth.Login)
	r.POST("/api/v1/auth/register", auth.Register)

	protected := r.Group("/api/v1")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/me", auth.Me)

		orderRepo := repository.NewOrderRepository(db.DB)
		orderService := service.NewOrderService(orderRepo)
		orderHandler := delivery.NewOrderHandler(orderService)

		orders := protected.Group("orders")
		{
			orders.GET("/", orderHandler.GetAllOrders)
			orders.POST("/", orderHandler.CreateOrder)
			orders.GET(":id", orderHandler.GetOrder)
			orders.PUT(":id", orderHandler.UpdateOrder)
			orders.DELETE(":id", orderHandler.DeleteOrder)
		}
	}
}
