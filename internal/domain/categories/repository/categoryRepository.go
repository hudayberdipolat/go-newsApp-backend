package repository

import "github.com/hudayberdipolat/go-newsApp-backend/internal/models"

type CategoryRepository interface {
	GetAll() ([]models.Category, error)
	GetOne(categoryID int) (*models.Category, error)
	Create(category models.Category) error
	Update(categoryID int, category models.Category) error
	Delete(categoryID int) error

	// functions for frontend

	GetAllCategories() ([]models.Category, error)
}
