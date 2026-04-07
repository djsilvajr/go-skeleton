package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/djsilvajr/go-skeleton/internal/config"
	"github.com/djsilvajr/go-skeleton/internal/domain/user/service"
	"github.com/djsilvajr/go-skeleton/internal/response"
)

type AuthHandler struct {
	svc service.UserService
	cfg *config.Config
}

func NewAuthHandler(svc service.UserService, cfg *config.Config) *AuthHandler {
	return &AuthHandler{svc: svc, cfg: cfg}
}

// Login godoc
// @Summary     Authenticate and return JWT
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       body body loginRequest true "Credentials"
// @Success     200  {object}  map[string]interface{}
// @Failure     401  {object}  map[string]interface{}
// @Router      /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusUnprocessableEntity, "Validation error", gin.H{"validation": err.Error()})
		return
	}

	user, err := h.svc.ValidateCredentials(req.Email, req.Password)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Invalid credentials", nil)
		return
	}

	token, err := generateToken(user.ID, string(user.Role), h.cfg)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not generate token", nil)
		return
	}

	response.Success(c, http.StatusOK, gin.H{"token": token, "type": "Bearer"})
}

// Register godoc
// @Summary     Register a new user
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       body body registerRequest true "User data"
// @Success     201  {object}  map[string]interface{}
// @Failure     422  {object}  map[string]interface{}
// @Router      /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusUnprocessableEntity, "Validation error", gin.H{"validation": err.Error()})
		return
	}

	user, err := h.svc.Create(req.Name, req.Email, req.Password)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	token, _ := generateToken(user.ID, string(user.Role), h.cfg)
	response.Success(c, http.StatusCreated, gin.H{"user": user, "token": token, "type": "Bearer"})
}

// --- helpers ---

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func generateToken(userID uint, role string, cfg *config.Config) (string, error) {
	claims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.JWTExpireHour) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}

// --- request structs ---

type loginRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type registerRequest struct {
	Name     string `json:"name"     binding:"required,min=2"`
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
