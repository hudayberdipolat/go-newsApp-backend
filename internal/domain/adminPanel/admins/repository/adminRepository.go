package repository

import "github.com/hudayberdipolat/go-newsApp-backend/internal/models"

type AdminRepository interface {
	GetAll() ([]models.Admin, error)
	GetOne(adminID int) (*models.Admin, error)
	Create(admin models.Admin) error
	Update(adminID int, admin models.Admin) error
	Delete(adminID int) error
	ChangePassword(adminID int, password string) error
}
