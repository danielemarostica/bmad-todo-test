package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"bmad-todo-test/handlers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockTodoStore struct {
	todos []handlers.Todo
}

func (m *mockTodoStore) Create(text string) (handlers.Todo, error) {
	todo := handlers.Todo{
		ID:        "fake-id-123",
		Text:      text,
		Completed: false,
		CreatedAt: time.Date(2026, 4, 10, 12, 0, 0, 0, time.UTC),
	}
	m.todos = append(m.todos, todo)
	return todo, nil
}

func setupTodoRouter(store handlers.TodoStore) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/api/v1/todos", handlers.CreateTodo(store))
	return r
}

func TestCreateTodo_ReturnsCreatedTodo_WithValidText(t *testing.T) {
	store := &mockTodoStore{}
	router := setupTodoRouter(store)

	body := strings.NewReader(`{"text": "Buy groceries"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/todos", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var todo handlers.Todo
	err := json.Unmarshal(w.Body.Bytes(), &todo)
	assert.NoError(t, err)
	assert.Equal(t, "Buy groceries", todo.Text)
	assert.False(t, todo.Completed)
	assert.NotEmpty(t, todo.ID)
	assert.False(t, todo.CreatedAt.IsZero())
}

func TestCreateTodo_Returns400_WithEmptyText(t *testing.T) {
	store := &mockTodoStore{}
	router := setupTodoRouter(store)

	body := strings.NewReader(`{"text": ""}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/todos", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"text is required"`)
	assert.Empty(t, store.todos)
}

func TestCreateTodo_Returns400_WithWhitespaceOnlyText(t *testing.T) {
	store := &mockTodoStore{}
	router := setupTodoRouter(store)

	body := strings.NewReader(`{"text": "   "}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/todos", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"text is required"`)
	assert.Empty(t, store.todos)
}
