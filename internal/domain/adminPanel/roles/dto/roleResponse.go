package dto

import "github.com/hudayberdipolat/go-newsApp-backend/internal/models"

type RoleResponse struct {
	ID       int    `json:"id"`
	RoleName string `json:"role_name"`
}

func NewRoleResponse(role models.Role) RoleResponse {
	return RoleResponse{
		ID:       role.ID,
		RoleName: role.RoleName,
	}
}
