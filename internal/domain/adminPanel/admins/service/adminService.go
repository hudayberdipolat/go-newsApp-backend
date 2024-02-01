package service

import "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/admins/dto"

type AdminService interface {
	FindAll() ([]dto.AdminResponse, error)
	FindOne(adminID int) (*dto.AdminResponse, error)
	Create(request dto.CreateAdminRequest) error
	Update(adminID int, request dto.UpdateAdminRequest) error
	Delete(adminID int) error
	ChangePassword(adminID int, changePassword dto.ChangeAdminPassword) error
}
