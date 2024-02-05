package service

import (
	dto "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/categories/dto"
)

type CategoryService interface {
	FindAll() ([]dto.CategoryAllResponse, error)
	FindOne(categoryID int) (*dto.CategoryResponse, error)
	CreateCategory(request dto.CreateCategoryRequest) error
	UpdateCategory(categoryID int, request dto.UpdateCategoryRequest) error
	DeleteCategory(categoryID int) error
}
