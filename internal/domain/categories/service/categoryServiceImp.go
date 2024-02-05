package service

import (
	"errors"
	"github.com/gosimple/slug"
	dto2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/categories/dto"
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

func (c categoryServiceImp) FindAll() ([]dto2.CategoryResponse, error) {
	categories, err := c.categoryRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var categoryResponses []dto2.CategoryResponse

	for _, category := range categories {
		categoryResponse := dto2.NewCategoryResponse(category)
		categoryResponses = append(categoryResponses, categoryResponse)
	}
	return categoryResponses, nil
}

func (c categoryServiceImp) FindOne(categoryID int) (*dto2.CategoryResponse, error) {
	category, err := c.categoryRepo.GetOne(categoryID)
	if err != nil {
		return nil, err
	}
	categoryResponse := dto2.NewCategoryResponse(*category)
	return &categoryResponse, nil
}

func (c categoryServiceImp) CreateCategory(request dto2.CreateCategoryRequest) error {
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

func (c categoryServiceImp) UpdateCategory(categoryID int, request dto2.UpdateCategoryRequest) error {

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
