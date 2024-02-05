package service

import (
	dto2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/admins/dto"
)

type AdminService interface {
	FindAll() ([]dto2.AdminResponse, error)
	FindOne(adminID int) (*dto2.AdminResponse, error)
	Create(request dto2.CreateAdminRequest) error
	Update(adminID int, request dto2.UpdateAdminRequest) error
	Delete(adminID int) error
	ChangePassword(adminID int, changePassword dto2.ChangeAdminPassword) error
}
