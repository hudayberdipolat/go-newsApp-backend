package service

import (
	"errors"
	"github.com/gosimple/slug"
	dto "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/categories/dto"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/categories/repository"
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

func (c categoryServiceImp) FindAll() ([]dto.CategoryAllResponse, error) {
	categories, err := c.categoryRepo.GetAll()
	if err != nil {
		return nil, err
	}
	categoryResponses := dto.NewAllCategoryResponse(categories)
	return categoryResponses, nil
}

func (c categoryServiceImp) FindOne(categoryID int) (*dto.CategoryResponse, error) {
	category, err := c.categoryRepo.GetOne(categoryID)
	if err != nil {
		return nil, err
	}
	categoryResponse := dto.NewOneCategoryResponse(category)
	return &categoryResponse, nil
}

func (c categoryServiceImp) CreateCategory(request dto.CreateCategoryRequest) error {
	if request.CategoryStatus == "" {
		request.CategoryStatus = "passive"
	}
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
	if errUpdate := c.categoryRepo.Update(updateCategory.ID, *updateCategory); errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (c categoryServiceImp) DeleteCategory(categoryID int) error {
	deleteCategory, err := c.categoryRepo.GetOne(categoryID)
	if err != nil {
		return errors.New("category not found")
	}
	if errDelete := c.categoryRepo.Delete(deleteCategory.ID); errDelete != nil {
		return errDelete
	}
	return nil
}

// functions for frontend

func (c categoryServiceImp) GetAllCategories() ([]dto.GetAllCategoriesResponse, error) {
	categories, err := c.categoryRepo.GetAllCategories()
	if err != nil {
		return nil, err
	}
	getAllCategoryResponses := dto.NewGetAllCategoriesResponse(categories)
	return getAllCategoryResponses, nil
}

func (c categoryServiceImp) GetOneCategory(categorySlug string) (*dto.GetCategoryResponse, error) {
	category, err := c.categoryRepo.GetOneCategory(categorySlug)
	if err != nil {
		return nil, err
	}
	categoryResponse := dto.NewGetCategoryResponse(category)
	return &categoryResponse, nil
}
