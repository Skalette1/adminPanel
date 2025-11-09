package handlers

import (
	"net/http"
	"strconv"

	"github.com/Skalette1/adminPanel/internal/dto"
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

// @Create godoc
// @Summary Create role
// @Description Create a new rple
// @Tags roles
// @Accept json
// @Produce json
// @Param role body models.Role true "Role object"
// @Success 201 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /roles [post]
func (h *RoleHandler) CreateRole(c *gin.Context) {
	var role models.Role
	_ = c.BindJSON(&role)

	user, err := h.Repo.Create(role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Can not create role",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Created role",
		Data:    user,
	})
}

// @Create godoc
// @Summary GetByID role
// @Description Get role by ID
// @Tags roles
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 200 {object} dto.SuccessResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /roles/{id} [get]
func (h *RoleHandler) GetRoleByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	role, err := h.Repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Can not get role",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Get role",
		Data:    role,
	})
}

// @Create godoc
// @Summary GetAll roles
// @Description Get roles
// @Tags roles
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 200 {object} dto.SuccessResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /roles [get]
func (h *RoleHandler) GetAllRoles(c *gin.Context) {
	roles, err := h.Repo.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Can not get all roles",
			Details: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Get all roles",
		Data:    roles,
	})
}

// @Create godoc
// @Summary Update role
// @Description Update role
// @Tags roles
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /roles [put]
func (h *RoleHandler) UpdateRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var role models.Role
	_ = c.BindJSON(&role)

	err := h.Repo.Update(id, role)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Can not update role",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Updated role",
		Data:    role,
	})
}

// @Create godoc
// @Summary Delete role
// @Description Delete role
// @Tags roles
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /roles [delete]
func (h *RoleHandler) DeleteRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.Repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Can not delete role",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Deleted role",
		Data:    nil,
	})
}

//
