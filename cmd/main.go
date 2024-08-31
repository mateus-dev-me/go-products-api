package main

import (
	"log"

	"go-products/interfaces/api"
	"go-products/interfaces/api/controllers"
	"go-products/internal/infrastructure"
	"go-products/internal/infrastructure/db"

	_ "github.com/lib/pq"
)

func main() {
	config, err := infrastructure.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	dbConnection, err := db.ConnectDB(config.GetDBConnectionString())
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer dbConnection.Close()

	repository := db.NewProductRepositoryDB(dbConnection)      // Domain
	controller := controllers.NewProductController(repository) // Adpters

	server := api.SetupRouter(controller)
	server.Run(":8000")
}
