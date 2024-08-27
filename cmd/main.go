package main

import (
	"log"

	"go-products/interfaces/api/controllers"
	"go-products/internal/infrastructure"
	"go-products/internal/infrastructure/db"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	server := gin.Default()

	config, err := infrastructure.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	dbConnection, err := db.ConnectDB(config.GetDBConnectionString())
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer dbConnection.Close()

	repository := db.NewProductRepositoryDB(dbConnection)
	controller := controllers.NewProductController(repository)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", controller.GetAllHandler)
	server.GET("/products/:id", controller.GetByIDHandler)
	server.POST("/products", controller.SaveHandler)
	server.PUT("/products/:id", controller.UpdateHandler)
	server.DELETE("/products/:id", controller.DeleteHandler)

	server.Run(":8000")
}
