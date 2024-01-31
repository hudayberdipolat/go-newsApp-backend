package service

import (
	"errors"
	"github.com/gosimple/slug"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/categories/dto"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/categories/repository"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
	"time"
)

type categoryServiceImp struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return categoryServiceImp{
		categoryRepo: repo,
	}
}

func (c categoryServiceImp) FindAll() ([]dto.CategoryResponse, error) {
	categories, err := c.categoryRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var categoryResponses []dto.CategoryResponse

	for _, category := range categories {
		categoryResponse := dto.NewCategoryResponse(category)
		categoryResponses = append(categoryResponses, categoryResponse)
	}
	return categoryResponses, nil
}

func (c categoryServiceImp) FindOne(categoryID int) (*dto.CategoryResponse, error) {
	category, err := c.categoryRepo.GetOne(categoryID)
	if err != nil {
		return nil, err
	}
	categoryResponse := dto.NewCategoryResponse(*category)
	return &categoryResponse, nil
}

func (c categoryServiceImp) CreateCategory(request dto.CreateCategoryRequest) error {
	createCategory := models.Category{
		CategoryName:   request.CategoryName,
		CategorySlug:   slug.Make(request.CategoryName),
		CategoryStatus: request.CategoryStatus,
		CreatedAt:      time.Now(),
	}
	if err := c.categoryRepo.Create(createCategory); err != nil {
		return err
	}
	return nil
}

func (c categoryServiceImp) UpdateCategory(categoryID int, request dto.UpdateCategoryRequest) error {

	updateCategory, err := c.categoryRepo.GetOne(categoryID)
	if err != nil {
		return errors.New("category not found")
	}
	updateCategory.CategoryName = request.CategoryName
	updateCategory.CategorySlug = slug.Make(request.CategoryName)
	updateCategory.CategoryStatus = request.CategoryStatus
	if errUpdate := c.categoryRepo.Update(categoryID, *updateCategory); errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (c categoryServiceImp) DeleteCategory(categoryID int) error {
	_, err := c.categoryRepo.GetOne(categoryID)
	if err != nil {
		return errors.New("category not found")
	}
	return nil
}
