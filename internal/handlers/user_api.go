package handlers

import (
	"net/http"
	"strconv"

	"github.com/Skalette1/adminPanel/internal/dto"
	"github.com/Skalette1/adminPanel/internal/models"
	"github.com/Skalette1/adminPanel/internal/repository"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}

// @Summary Create user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 201 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /users [post]
func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Bad Request",
			Details: err.Error(),
		})
		return
	}

	id, err := h.Repo.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Internal Server Error",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse{
		Message: "User Created",
		Data:    id,
	})
}

// GetUserByIDHandler godoc
// @Summary Get user by ID
// @Description Get user details by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dto.SuccessResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetUserByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.Repo.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Internal Server Error",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "User Found",
		Data:    user,
	})
}

// UpdateUserByIDHandler godoc
// @Summary Update user by ID
// @Description Update user details by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUserByIDHandler(c *gin.Context) {
	_, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Bad Request",
			Details: err.Error(),
		})
		return
	}
	if err := h.Repo.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Internal Server Error",
			Details: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "User Updated",
		Data:    user,
	})
}

// DeleteUserByIDHandler godoc
// @Summary Delete user by ID
// @Description Delete user details by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUserByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Bad Request",
			Details: err.Error(),
		})
		return
	}
	if err := h.Repo.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Internal Server Error",
			Details: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "User Deleted",
		Data:    id,
	})
}

// @Summary Get users
// @Description Get users
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 200 {object} dto.SuccessResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /users [get]
func (h *UserHandler) GetAllUsersHandler(c *gin.Context) {
	users, err := h.Repo.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Internal Server Error",
			Details: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Users Found",
		Data:    users,
	})
}
