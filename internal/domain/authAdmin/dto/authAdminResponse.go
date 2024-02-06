package dto

import (
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
)

type AuthAdminResponse struct {
	ID          int    `json:"id"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	AdminRole   string `json:"admin_role"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	AccessToken string `json:"access_token"`
}

func NewAuthAdminResponse(admin models.Admin, accessToken string) AuthAdminResponse {
	return AuthAdminResponse{
		ID:          admin.ID,
		FullName:    admin.FullName,
		PhoneNumber: admin.PhoneNumber,
		AdminRole:   admin.AdminRole,
		CreatedAt:   admin.CreatedAt.Format("01-02-2006"),
		UpdatedAt:   admin.UpdatedAt.Format("01-02-2006"),
		AccessToken: accessToken,
	}
}
