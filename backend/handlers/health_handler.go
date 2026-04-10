package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pinger interface {
	Ping() error
}

func HealthHandler(pinger Pinger) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := pinger.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "error", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
