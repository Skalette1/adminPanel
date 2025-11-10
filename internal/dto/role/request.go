package role

type CreateRoleRequest struct {
	Username   string  `json:"username" binding:"required" example:"admin"`
	Permission *string `json:"permission,omitempty" example:"full_access"`
}

type UpdateRoleRequest struct {
	Username   *string `json:"username,omitempty" example:"admin"`
	Permission *string `json:"permission,omitempty" example:"full_access"`
}
