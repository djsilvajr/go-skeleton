package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/djsilvajr/go-skeleton/internal/config"
	"github.com/djsilvajr/go-skeleton/internal/domain/user/handler"
	"github.com/djsilvajr/go-skeleton/internal/domain/user/model"
	"github.com/djsilvajr/go-skeleton/internal/domain/user/service"
)

// stubUserService satisfies service.UserService for handler tests.
type stubUserService struct {
	service.UserService
	createFn            func(name, email, password string) (*model.User, error)
	validateFn          func(email, password string) (*model.User, error)
}

func (s *stubUserService) Create(name, email, password string) (*model.User, error) {
	return s.createFn(name, email, password)
}
func (s *stubUserService) ValidateCredentials(email, password string) (*model.User, error) {
	return s.validateFn(email, password)
}

func setupAuthRouter(svc service.UserService, cfg *config.Config) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h := handler.NewAuthHandler(svc, cfg)
	r.POST("/api/v1/auth/login", h.Login)
	r.POST("/api/v1/auth/register", h.Register)
	return r
}

func TestAuthHandler_Login_Success(t *testing.T) {
	cfg := &config.Config{JWTSecret: "test-secret", JWTExpireHour: 1}
	svc := &stubUserService{
		validateFn: func(email, password string) (*model.User, error) {
			return &model.User{ID: 1, Email: email, Role: model.RoleUser}, nil
		},
	}

	r := setupAuthRouter(svc, cfg)
	body, _ := json.Marshal(map[string]string{"email": "alice@example.com", "password": "secret"})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d — body: %s", w.Code, w.Body.String())
	}

	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp["token"] == "" {
		t.Error("expected JWT token in response")
	}
}

func TestAuthHandler_Login_InvalidCredentials(t *testing.T) {
	cfg := &config.Config{JWTSecret: "test-secret", JWTExpireHour: 1}
	svc := &stubUserService{
		validateFn: func(email, password string) (*model.User, error) {
			return nil, service.ErrInvalidCredentials
		},
	}

	r := setupAuthRouter(svc, cfg)
	body, _ := json.Marshal(map[string]string{"email": "x@x.com", "password": "wrong"})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", w.Code)
	}
}
