package handlers_test

import (
	"encoding/json"
	"errors"
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

func (m *mockTodoStore) List() ([]handlers.Todo, error) {
	return m.todos, nil
}

func (m *mockTodoStore) Delete(id string) error {
	for i, todo := range m.todos {
		if todo.ID == id {
			m.todos = append(m.todos[:i], m.todos[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

func (m *mockTodoStore) Update(id string, completed bool) (handlers.Todo, error) {
	for i, todo := range m.todos {
		if todo.ID == id {
			m.todos[i].Completed = completed
			return m.todos[i], nil
		}
	}
	return handlers.Todo{}, errors.New("not found")
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
	r.GET("/api/v1/todos", handlers.ListTodos(store))
	r.POST("/api/v1/todos", handlers.CreateTodo(store))
	r.PATCH("/api/v1/todos/:id", handlers.UpdateTodo(store))
	r.DELETE("/api/v1/todos/:id", handlers.DeleteTodo(store))
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

func TestListTodos_ReturnsTodos_WhenTheyExist(t *testing.T) {
	store := &mockTodoStore{
		todos: []handlers.Todo{
			{ID: "1", Text: "First", Completed: false, CreatedAt: time.Date(2026, 4, 10, 12, 0, 0, 0, time.UTC)},
			{ID: "2", Text: "Second", Completed: true, CreatedAt: time.Date(2026, 4, 10, 13, 0, 0, 0, time.UTC)},
		},
	}
	router := setupTodoRouter(store)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/todos", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var todos []handlers.Todo
	err := json.Unmarshal(w.Body.Bytes(), &todos)
	assert.NoError(t, err)
	assert.Len(t, todos, 2)
	assert.Equal(t, "First", todos[0].Text)
	assert.Equal(t, "Second", todos[1].Text)
}

func TestListTodos_ReturnsEmptyArray_WhenNoTodos(t *testing.T) {
	store := &mockTodoStore{}
	router := setupTodoRouter(store)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/todos", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "[]", strings.TrimSpace(w.Body.String()))
}

func TestUpdateTodo_ReturnsUpdatedTodo_WithValidID(t *testing.T) {
	store := &mockTodoStore{
		todos: []handlers.Todo{
			{ID: "abc123", Text: "Test todo", Completed: false, CreatedAt: time.Date(2026, 4, 10, 12, 0, 0, 0, time.UTC)},
		},
	}
	router := setupTodoRouter(store)

	body := strings.NewReader(`{"completed": true}`)
	req := httptest.NewRequest(http.MethodPatch, "/api/v1/todos/abc123", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var todo handlers.Todo
	err := json.Unmarshal(w.Body.Bytes(), &todo)
	assert.NoError(t, err)
	assert.True(t, todo.Completed)
	assert.Equal(t, "Test todo", todo.Text)
}

func TestUpdateTodo_Returns404_WithInvalidID(t *testing.T) {
	store := &mockTodoStore{}
	router := setupTodoRouter(store)

	body := strings.NewReader(`{"completed": true}`)
	req := httptest.NewRequest(http.MethodPatch, "/api/v1/todos/nonexistent", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"todo not found"`)
}

func TestDeleteTodo_Returns204_WithValidID(t *testing.T) {
	store := &mockTodoStore{
		todos: []handlers.Todo{
			{ID: "abc123", Text: "Test todo", Completed: false, CreatedAt: time.Date(2026, 4, 10, 12, 0, 0, 0, time.UTC)},
		},
	}
	router := setupTodoRouter(store)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/todos/abc123", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Empty(t, store.todos)
}

func TestDeleteTodo_Returns404_WithInvalidID(t *testing.T) {
	store := &mockTodoStore{}
	router := setupTodoRouter(store)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/todos/nonexistent", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"todo not found"`)
}
