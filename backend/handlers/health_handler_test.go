package handlers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"bmad-todo-test/handlers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockPinger struct {
	err error
}

func (m *mockPinger) Ping() error {
	return m.err
}

func setupRouter(pinger handlers.Pinger) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/api/health", handlers.HealthHandler(pinger))
	return r
}

func TestHealthEndpoint_ReturnsOK_WhenMongoReachable(t *testing.T) {
	router := setupRouter(&mockPinger{err: nil})

	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"status":"ok"`)
}

func TestHealthEndpoint_Returns503_WhenMongoUnreachable(t *testing.T) {
	router := setupRouter(&mockPinger{err: errors.New("connection refused")})

	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
	assert.Contains(t, w.Body.String(), `"status":"error"`)
}
