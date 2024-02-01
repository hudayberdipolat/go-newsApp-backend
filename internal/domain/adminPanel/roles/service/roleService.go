package service

import "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/roles/dto"

type RoleService interface {
	FindAll() ([]dto.RoleResponse, error)
	FindOne(roleID int) (*dto.RoleResponse, error)
	Create(request dto.CreateRoleRequest) error
	Update(roleID int, request dto.UpdateRoleRequest) error
	Delete(roleID int) error
}
