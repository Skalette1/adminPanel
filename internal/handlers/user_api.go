package handlers

import (
	"net/http"
	"strconv"
	"time"

	userdto "github.com/Skalette1/adminPanel/internal/dto/user"
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

// CreateUserHandler godoc
// @Summary Create user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body userdto.CreateUserRequest true "User object"
// @Success 201 {object} userdto.UserCreatedResponse
// @Failure 400 {object} userdto.UserErrorResponse
// @Failure 404 {object} userdto.UserNotFoundResponse
// @Failure 501 {object} userdto.UserInternalErrorResponse
// @Router /users [post]
func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	var req userdto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, userdto.UserErrorResponse{
			Message: "invalid request",
			Details: err.Error(),
		})
	}
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		RoleId:   req.RoleId,
	}

	id, err := h.Repo.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, userdto.UserErrorResponse{
			Message: "Can not create user",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, userdto.UserSuccessResponse{
		ID:        id,
		Username:  user.Username,
		Email:     user.Email,
		RoleId:    user.RoleId,
		IsActive:  true,
		CreatedAt: time.Now(),
	})
}

// GetUserByIDHandler godoc
// @Summary Get user by ID
// @Description Get user details by ID
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} userdto.UserSuccessResponse
// @Failure 400 {object} userdto.UserErrorResponse
// @Failure 404 {object} userdto.UserNotFoundResponse
// @Failure 501 {object} userdto.UserInternalErrorResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetUserByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.Repo.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, userdto.UserErrorResponse{
			Message: "Can not get user",
			Details: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, userdto.UserSuccessResponse{
		ID:       id,
		Username: user.Username,
		Email:    user.Email,
		RoleId:   user.RoleId,
		IsActive: true,
	})
}

// GetAllUsersHandler godoc
// @Summary Get user by ID
// @Description Get user details by ID
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} userdto.UserSuccessResponse
// @Failure 400 {object} userdto.UserErrorResponse
// @Failure 404 {object} userdto.UserNotFoundResponse
// @Failure 501 {object} userdto.UserInternalErrorResponse
// @Router /users/ [get]
func (h *UserHandler) GetAllUsersHandler(c *gin.Context) {
	users, err := h.Repo.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, userdto.UserErrorResponse{
			Message: "Can not get users",
			Details: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, userdto.UserSuccessResponse{
		ID:       len(users),
		Username: users[0].Username,
		Email:    users[0].Email,
		RoleId:   users[0].RoleId,
		IsActive: true,
	})
}

// UpdateUserByIDHandler godoc
// @Summary Get user by ID
// @Description Get user details by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} userdto.UserSuccessResponse
// @Failure 400 {object} userdto.UserErrorResponse
// @Failure 404 {object} userdto.UserNotFoundResponse
// @Failure 501 {object} userdto.UserInternalErrorResponse
// @Router /users [put]
func (h *UserHandler) UpdateUserByIDHandler(c *gin.Context) {
	_, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, userdto.UserErrorResponse{
			Message: "Invalid id",
			Details: err.Error(),
		})
		return
	}
	var req userdto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, userdto.UserErrorResponse{
			Message: "Invalid request body",
			Details: err.Error(),
		})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	existing, err := h.Repo.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, userdto.UserErrorResponse{
			Message: "User not found",
			Details: err.Error(),
		})
		return
	}

	switch {
	case req.Username != nil:
		existing.Username = *req.Username
	case req.Email != nil:
		existing.Email = *req.Email
	case req.Password != nil:
		existing.Password = *req.Password
	case req.RoleId != nil:
		existing.RoleId = *req.RoleId
	case req.IsActive != nil:
		existing.IsActive = *req.IsActive
	}

	if err := h.Repo.UpdateUser(existing); err != nil {
		c.JSON(http.StatusBadRequest, userdto.UserErrorResponse{
			Message: "Can not update user",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, userdto.UserSuccessResponse{
		ID:        id,
		UpdatedAt: time.Now(),
	})
}

// DeleteUserByIDHandler godoc
// @Summary Get user by ID
// @Description Get user details by ID
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} userdto.UserSuccessResponse
// @Failure 400 {object} userdto.UserErrorResponse
// @Failure 404 {object} userdto.UserNotFoundResponse
// @Failure 501 {object} userdto.UserInternalErrorResponse
// @Router /users [delete]
func (h *UserHandler) DeleteUserByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, userdto.UserErrorResponse{
			Message: "Invalid id",
			Details: err.Error(),
		})
		return
	}
	if err := h.Repo.DeleteUser(id); err != nil {
		c.JSON(http.StatusBadRequest, userdto.UserErrorResponse{
			Message: "Can not delete user",
			Details: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, userdto.UserSuccessResponse{
		ID: id,
	})
}
