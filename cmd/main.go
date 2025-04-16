package main

import (
	"github.com/gin-gonic/gin"
	"order-project/internal/db"
	"order-project/internal/routes"
)

func main() {
	db.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
