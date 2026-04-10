package main

import (
	"log"

	"bmad-todo-test/config"
	"bmad-todo-test/handlers"
	"bmad-todo-test/middleware"
	"bmad-todo-test/models"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db := models.ConnectDB(cfg.MongoURI)

	router := gin.Default()
	middleware.Setup(router, cfg)

	router.GET("/api/health", handlers.HealthHandler(db))

	log.Printf("Server starting on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
