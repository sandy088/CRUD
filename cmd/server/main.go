package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"saaster.tech/crud/config"
	"saaster.tech/crud/internal/handler"
	repository "saaster.tech/crud/internal/repositories"
	route "saaster.tech/crud/internal/routes"
	"saaster.tech/crud/internal/services"
	"saaster.tech/crud/pkg/db"
)

func main() {
	cfg := config.LoadConfig()

	//initialize mongoDB
	client := db.ConnectMongoDB(cfg.MongoURI)
	defer client.Disconnect(context.TODO())

	//initialize repositories, services and handlers
	userRepo := repository.NewUserRepository(client, cfg.MongoDatabase)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	//setup Gin
	r := gin.Default()
	route.RegisterRoutes(r, userHandler)

	//start server
	log.Println("Starting server on port 3000")
	err := r.Run(":3000")

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
