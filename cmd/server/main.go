package main

import (
	"log"

	"my-redis-app/internal/handlers"
	"my-redis-app/internal/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	redisClient := storage.NewRedisClient()
	defer redisClient.Close()

	router := gin.Default()

	router.GET("/data/:id", handlers.GetDataHandler(redisClient))

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
