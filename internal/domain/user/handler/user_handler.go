package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/djsilvajr/go-skeleton/internal/domain/user/service"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// List godoc
// @Summary     List users
// @Tags        users
// @Security    BearerAuth
// @Produce     json
// @Success     200  {array}   model.User
// @Router      /users [get]
func (h *UserHandler) List(c *gin.Context) {
	users, err := h.svc.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// Show godoc
// @Summary     Get user by ID
// @Tags        users
// @Security    BearerAuth
// @Param       id   path int true "User ID"
// @Produce     json
// @Success     200  {object}  model.User
// @Failure     404  {object}  map[string]string
// @Router      /users/{id} [get]
func (h *UserHandler) Show(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	user, err := h.svc.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Store godoc
// @Summary     Create user
// @Tags        users
// @Security    BearerAuth
// @Accept      json
// @Produce     json
// @Param       body body storeRequest true "User payload"
// @Success     201  {object}  model.User
// @Failure     422  {object}  map[string]string
// @Router      /users [post]
func (h *UserHandler) Store(c *gin.Context) {
	var req storeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	user, err := h.svc.Create(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": user})
}

// Update godoc
// @Summary     Update user
// @Tags        users
// @Security    BearerAuth
// @Accept      json
// @Produce     json
// @Param       id   path int true "User ID"
// @Param       body body updateRequest true "User payload"
// @Success     200  {object}  model.User
// @Router      /users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req updateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	user, err := h.svc.Update(id, req.Name, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Destroy godoc
// @Summary     Delete user (admin only)
// @Tags        users
// @Security    BearerAuth
// @Param       id   path int true "User ID"
// @Produce     json
// @Success     204
// @Failure     403  {object}  map[string]string
// @Router      /users/{id} [delete]
func (h *UserHandler) Destroy(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.svc.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// --- request structs ---

type storeRequest struct {
	Name     string `json:"name"     binding:"required,min=2"`
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type updateRequest struct {
	Name  string `json:"name"  binding:"required,min=2"`
	Email string `json:"email" binding:"required,email"`
}

func parseID(c *gin.Context) (uint, error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	return uint(id), err
}
