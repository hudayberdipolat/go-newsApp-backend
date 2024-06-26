package repository

import (
	"errors"
	"log"

	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
	"gorm.io/gorm"
)

type categoryRepositoryImp struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return categoryRepositoryImp{
		db: db,
	}
}

func (c categoryRepositoryImp) Edit(categoryID int) (*models.Category, error) {
	var category models.Category
	if err := c.db.First(&category, categoryID).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c categoryRepositoryImp) GetAll() ([]models.Category, error) {
	var categories []models.Category
	if err := c.db.Preload("Posts").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (c categoryRepositoryImp) GetOne(categoryID int) (*models.Category, error) {
	var category models.Category
	if err := c.db.Where("id=?", categoryID).Preload("Posts").First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c categoryRepositoryImp) FindCategory(categoryID int) (*models.Category, error) {
	var category models.Category
	if err := c.db.First(&category, categoryID).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c categoryRepositoryImp) Create(category models.Category) error {
	if err := c.db.Create(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("this category name already exists")
		}
		return err
	}
	return nil
}

func (c categoryRepositoryImp) Update(categoryID int, category models.Category) error {
	var categoryModel models.Category
	if err := c.db.Model(&categoryModel).Where("id=?", categoryID).Updates(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("this category name already exists")
		}
		return err
	}
	return nil
}

func (c categoryRepositoryImp) Delete(categoryID int) error {
	var category models.Category
	if err := c.db.Where("id=?", categoryID).Unscoped().Delete(&category).Error; err != nil {
		return err
	}
	return nil
}

// functions for frontend

func (c categoryRepositoryImp) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	categoryStatus := "active"
	// var postCount int64
	if err := c.db.Select("id", "category_name", "category_slug").Where("category_status=?", categoryStatus).Preload("Posts").Find(&categories).Error; err != nil {
		return nil, err
	}
	for _, category := range categories {
		log.Println(category.CreatedAt)
	}
	return categories, nil
}

func (c categoryRepositoryImp) GetOneCategory(categorySlug string) (*models.Category, error) {
	var category models.Category
	activeStatus := "active"

	if err := c.db.Table("categories").
		Select("categories.id, categories.category_name, categories.category_slug").
		Where("category_status = ?", activeStatus).
		Where("category_slug = ?", categorySlug).
		Preload("Posts").First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	return &category, nil
}

func (c categoryRepositoryImp) GetCategoryBySlug(categorySlug string) (*models.Category, error) {
	var category models.Category
	if err := c.db.Where("category_slug=?", categorySlug).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
