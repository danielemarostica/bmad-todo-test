package middleware

import (
	"bmad-todo-test/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, cfg *config.Config) {
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{cfg.CORSOrigin},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Content-Type"},
	}))

	router.Use(gin.Recovery())
}
