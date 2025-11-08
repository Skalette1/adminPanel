package handlers

import (
	"net/http"
	"strconv"

	"github.com/Skalette1/adminPanel/internal/models"
	"github.com/Skalette1/adminPanel/internal/repository"
	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	Repo *repository.RoleRepository
}

func NewRoleHandler(repo *repository.RoleRepository) *RoleHandler {
	return &RoleHandler{Repo: repo}
}

func (h *RoleHandler) Create(c *gin.Context) {
	var role models.Role
	_ = c.BindJSON(&role)

	id, err := h.Repo.Create(role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *RoleHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	role, err := h.Repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role": role})
}

func (h *RoleHandler) GetAll(c *gin.Context) {
	roles, err := h.Repo.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"roles": roles})
}

func (h *RoleHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var role models.Role
	_ = c.BindJSON(&role)

	err := h.Repo.Update(id, role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role": role})
}

func (h *RoleHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.Repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
