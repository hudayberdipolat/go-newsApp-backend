package repository

import "github.com/hudayberdipolat/go-newsApp-backend/internal/models"

type AuthAdminRepository interface {
	GetAdmin(phoneNumber string) (*models.Admin, error)
}
