package service

import "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/authAdmin/dto"

type AuthAdminService interface {
	LoginAdmin(request dto.AdminLoginRequest) (*dto.AuthAdminResponse, error)
}
