package repository

import (
	"errors"
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

func (c categoryRepositoryImp) GetAll() ([]models.Category, error) {
	var categories []models.Category
	if err := c.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (c categoryRepositoryImp) GetOne(categoryID int) (*models.Category, error) {
	var category models.Category
	if err := c.db.Where("id=?", categoryID).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c categoryRepositoryImp) Create(category models.Category) error {
	if err := c.db.Create(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("Bu kategoriýa ady eýýäm ulanylýar!!!")
		}
		return err
	}
	return nil
}

func (c categoryRepositoryImp) Update(categoryID int, category models.Category) error {
	var categoryModel models.Category
	if err := c.db.Model(&categoryModel).Where("id=?", categoryID).Updates(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("Bu kategoriýa ady eýýäm ulanylýar!!!")
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
