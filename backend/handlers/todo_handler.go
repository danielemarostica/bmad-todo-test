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

type UpdateTodoRequest struct {
	Completed *bool `json:"completed" binding:"required"`
}

type TodoStore interface {
	Create(text string) (Todo, error)
	List() ([]Todo, error)
	Update(id string, completed bool) (Todo, error)
	Delete(id string) error
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

func ListTodos(store TodoStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		todos, err := store.List()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list todos"})
			return
		}

		if todos == nil {
			todos = []Todo{}
		}

		c.JSON(http.StatusOK, todos)
	}
}

func UpdateTodo(store TodoStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var req UpdateTodoRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "completed field is required"})
			return
		}

		todo, err := store.Update(id, *req.Completed)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
			return
		}

		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(store TodoStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if err := store.Delete(id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
