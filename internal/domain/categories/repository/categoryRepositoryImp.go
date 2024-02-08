package repository

import (
	"errors"
	"fmt"
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

// functions for frontend

func (c categoryRepositoryImp) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	categoryStatus := "active"
	if err := c.db.Where("category_status=?", categoryStatus).Preload("Posts").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (c categoryRepositoryImp) GetOneCategory(categorySlug string) (*models.Category, error) {
	var category models.Category
	categoryStatus := "active"
	if err := c.db.Select("id", "category_name", "category_slug").Where("category_status=?", categoryStatus).Where("category_slug=?",
		categorySlug).Preload("Posts", func() {

	}).First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	//sqlQuery := `select categories.ID, categories.category_name, categories.category_slug, posts.id, posts.post_title, posts.post_slug
	//				from categories
	//			left join posts on categories.id = posts.category_id
	//			where categories.category_status=? and  categories.category_slug=?`
	//if err := c.db.Raw(sqlQuery, categoryStatus, categorySlug).Scan(&category).Error; err != nil {
	//	return nil, err
	//}

	for _, post := range category.Posts {
		fmt.Printf("post-id  :%d  post_desc :%s \n", post.ID, post)
	}

	return &category, nil
}
