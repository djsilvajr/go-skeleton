package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/djsilvajr/go-skeleton/internal/domain/user/service"
	"github.com/djsilvajr/go-skeleton/internal/response"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) List(c *gin.Context) {
	users, err := h.svc.List()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not retrieve users", nil)
		return
	}
	response.Success(c, http.StatusOK, users)
}

func (h *UserHandler) Show(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid id", nil)
		return
	}

	user, err := h.svc.GetByID(id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "User not found", nil)
		return
	}

	response.Success(c, http.StatusOK, user)
}

func (h *UserHandler) Store(c *gin.Context) {
	var req storeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusUnprocessableEntity, "Validation error", gin.H{"validation": err.Error()})
		return
	}

	user, err := h.svc.Create(req.Name, req.Email, req.Password)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	response.Success(c, http.StatusCreated, user)
}

func (h *UserHandler) Update(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid id", nil)
		return
	}

	var req updateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusUnprocessableEntity, "Validation error", gin.H{"validation": err.Error()})
		return
	}

	user, err := h.svc.Update(id, req.Name, req.Email)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, user)
}

// Destroy performs a soft-delete and returns {"data": {"deleted": true}}.
func (h *UserHandler) Destroy(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid id", nil)
		return
	}

	if err := h.svc.Delete(id); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	response.Deleted(c)
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
