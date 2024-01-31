package service

import "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/categories/dto"

type CategoryService interface {
	FindAll() ([]dto.CategoryResponse, error)
	FindOne(categoryID int) (*dto.CategoryResponse, error)
	CreateCategory(request dto.CreateCategoryRequest) error
	UpdateCategory(categoryID int, request dto.UpdateCategoryRequest) error
	DeleteCategory(categoryID int) error
}
