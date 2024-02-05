package dto

type CreateRoleRequest struct {
	RoleName string `json:"role_name" validate:"required"`
}

type UpdateRoleRequest struct {
	RoleName string `json:"role_name" validate:"required"`
}
