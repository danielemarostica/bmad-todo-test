package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateTodoRequest struct {
	Text string `json:"text"`
}

type TodoStore interface {
	Create(text string) (Todo, error)
}

func CreateTodo(store TodoStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateTodoRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		trimmed := strings.TrimSpace(req.Text)
		if trimmed == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "text is required"})
			return
		}

		todo, err := store.Create(trimmed)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create todo"})
			return
		}

		c.JSON(http.StatusCreated, todo)
	}
}
