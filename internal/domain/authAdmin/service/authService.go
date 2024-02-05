package service

import (
	dto2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/authAdmin/dto"
)

type AuthAdminService interface {
	LoginAdmin(request dto2.AdminLoginRequest) (*dto2.AuthAdminResponse, error)
}
