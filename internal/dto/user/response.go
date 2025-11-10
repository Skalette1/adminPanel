package user

import "time"

type UserSuccessResponse struct {
	ID        int       `json:"id" example:"1"`
	Username  string    `json:"username" example:"beslan"`
	Email     string    `json:"email" example:"dudaev.beslan@bk.ru"`
	RoleId    int       `json:"role_id" example:"1"`
	IsActive  bool      `json:"is_active" example:"true"`
	CreatedAt time.Time `json:"created_at,omitempty" example:"2025-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2025-01-01T01:00:00Z"`
}

type UserErrorResponse struct {
	Message string `json:"message" example:"status bad request"`
	Details string `json:"details" example:"status bad request"`
}

type UserCreatedResponse struct {
	ID int `json:"id" example:"1"`
}

type UserNotFoundResponse struct {
	Message string `json:"message" example:"user not found"`
}

type UserInternalErrorResponse struct {
	Message string `json:"message" example:"internal server error"`
}
