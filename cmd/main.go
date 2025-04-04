package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"rest-project/internal/models"
	"rest-project/internal/routes"
)

func main() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:root@localhost:5432/postgres?sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	err = db.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatal("Error on migrating to the DB", err)
	}

	r := gin.Default()
	routes.SetupRoutes(r, db)
	r.Run(":8080")
}
