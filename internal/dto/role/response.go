package role

import "time"

type RoleSuccessResponse struct {
	ID         int       `json:"id" example:"1"`
	Username   string    `json:"username" example:"Beslan"`
	Permission string    `json:"permission" example:"full_access"`
	CreatedAt  time.Time `json:"created_at,omitempty" example:"2025-01-01T00:00:00Z"`
	UpdatedAt  time.Time `json:"updated_at,omitempty" example:"2025-01-01T01:00:00Z"`
}

type RoleErrorResponse struct {
	Message string `json:"message" example:"status bad request"`
	Details string `json:"details" example:"status bad request"`
}
