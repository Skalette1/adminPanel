package handlers

import (
	"net/http"
	"strconv"
	"time"

	roledto "github.com/Skalette1/adminPanel/internal/dto/role"
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

// @CreateRoleHandler godoc
// @Summary Create role
// @Description Create a new role
// @Tags roles
// @Accept json
// @Produce json
// @Param user body roledto.CreateRoleRequest true "User object"
// @Success 201 {object} roledto.RoleSuccessResponse
// @Failure 400 {object} roledto.RoleErrorResponse
// @Router /roles [post]
func (h *RoleHandler) CreateRoleHandler(c *gin.Context) {
	var req roledto.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, roledto.RoleErrorResponse{
			Message: "Invalid request body",
			Details: err.Error(),
		})
		return
	}

	role := models.Role{
		Username: req.Username,
	}
	if req.Permission != nil {
		role.Permission = *req.Permission
	}

	id, err := h.Repo.Create(role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, roledto.RoleErrorResponse{
			Message: "Can not create role",
			Details: err.Error(),
		})
		return
	}

	created, err := h.Repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, roledto.RoleErrorResponse{
			Message: "Role created but could not retrieve",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, roledto.RoleSuccessResponse{
		ID:         id,
		Username:   created.Username,
		Permission: created.Permission,
		CreatedAt:  time.Now(),
	})
}

// @GetRoleByIDHandler godoc
// @Summary GetByID role
// @Description Get role by ID
// @Tags roles
// @Accept json
// @Produce json
// @Success 200 {object} roledto.RoleSuccessResponse
// @Failure 404 {object} roledto.RoleErrorResponse
// @Router /roles/{id} [get]
func (h *RoleHandler) GetRoleByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	role, err := h.Repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, roledto.RoleErrorResponse{
			Message: "Can not get role",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, roledto.RoleSuccessResponse{
		ID:         role.ID,
		Username:   role.Username,
		Permission: role.Permission,
	})
}

// @GetAllRolesHandler godoc
// @Summary GetAll roles
// @Description Get roles
// @Tags roles
// @Accept json
// @Produce json
// @Success 200 {object} roledto.RoleSuccessResponse
// @Failure 404 {object} roledto.RoleErrorResponse
// @Router /roles [get]
func (h *RoleHandler) GetAllRolesHandler(c *gin.Context) {
	roles, err := h.Repo.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, roledto.RoleErrorResponse{
			Message: "Can not get all roles",
			Details: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, roledto.RoleSuccessResponse{
		ID:         len(roles),
		Username:   roles[0].Username,
		Permission: roles[0].Permission,
	})
}

// @UpdateRoleHandler godoc
// @Summary Update role
// @Description Update role
// @Tags roles
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 200 {object} roledto.RoleSuccessResponse
// @Failure 400 {object} roledto.RoleErrorResponse
// @Router /roles [put]
func (h *RoleHandler) UpdateRoleHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req roledto.UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, roledto.RoleErrorResponse{
			Message: "Invalid request body",
			Details: err.Error(),
		})
		return
	}

	existing, err := h.Repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, roledto.RoleErrorResponse{
			Message: "Role not found",
			Details: err.Error(),
		})
		return
	}

	if req.Username != nil {
		existing.Username = *req.Username
	}
	if req.Permission != nil {
		existing.Permission = *req.Permission
	}

	err = h.Repo.Update(id, *existing)
	if err != nil {
		c.JSON(http.StatusBadRequest, roledto.RoleErrorResponse{
			Message: "Can not update role",
			Details: err.Error(),
		})
		return
	}

	updated, err := h.Repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, roledto.RoleErrorResponse{
			Message: "Updated but could not retrieve",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, roledto.RoleSuccessResponse{
		ID:         id,
		Username:   updated.Username,
		Permission: updated.Permission,
		UpdatedAt:  time.Now(),
	})
}

// @DeleteRoleHandler godoc
// @Summary Delete role
// @Description Delete role
// @Tags roles
// @Accept json
// @Produce json
// @Success 200 {object} roledto.RoleSuccessResponse
// @Failure 400 {object} roledto.RoleErrorResponse
// @Router /roles [delete]
func (h *RoleHandler) DeleteRoleHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.Repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, roledto.RoleErrorResponse{
			Message: "Can not delete role",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, roledto.RoleSuccessResponse{
		ID: id,
	})
}
