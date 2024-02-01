package repository

import "github.com/hudayberdipolat/go-newsApp-backend/internal/models"

type RoleRepository interface {
	GetAll() ([]models.Role, error)
	GetOne(roleID int) (*models.Role, error)
	Create(role models.Role) error
	Update(roleID int, role models.Role) error
	Delete(roleID int) error
}
