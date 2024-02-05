package service

import (
	dto2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/roles/dto"
)

type RoleService interface {
	FindAll() ([]dto2.RoleResponse, error)
	FindOne(roleID int) (*dto2.RoleResponse, error)
	Create(request dto2.CreateRoleRequest) error
	Update(roleID int, request dto2.UpdateRoleRequest) error
	Delete(roleID int) error
}
