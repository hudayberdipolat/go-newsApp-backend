package service

import (
	dto2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/categories/dto"
)

type CategoryService interface {
	FindAll() ([]dto2.CategoryResponse, error)
	FindOne(categoryID int) (*dto2.CategoryResponse, error)
	CreateCategory(request dto2.CreateCategoryRequest) error
	UpdateCategory(categoryID int, request dto2.UpdateCategoryRequest) error
	DeleteCategory(categoryID int) error
}
