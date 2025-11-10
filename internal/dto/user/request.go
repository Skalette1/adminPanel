package user

type CreateUserRequest struct {
	Username string `json:"username" binding:"required" example:"ivan"`
	Email    string `json:"email" binding:"required,email" example:"ivan@example.com"`
	Password string `json:"password" binding:"required" example:"secret"`
	RoleId   int    `json:"role_id,omitempty" example:"1"`
}

type UpdateUserRequest struct {
	ID       *int    `json:"id,omitempty" example:"1"`
	Username *string `json:"username,omitempty" example:"ivan"`
	Email    *string `json:"email,omitempty" example:"ivan@example.com"`
	Password *string `json:"password,omitempty" example:"secret"`
	RoleId   *int    `json:"role_id,omitempty" example:"1"`
	IsActive *bool   `json:"is_active,omitempty" example:"true"`
}
