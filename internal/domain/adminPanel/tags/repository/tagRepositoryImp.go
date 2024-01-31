package repository

import (
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
	"gorm.io/gorm"
)

type tagRepositoryImp struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return tagRepositoryImp{
		db: db,
	}
}

func (t tagRepositoryImp) GetAll() ([]models.Tag, error) {
	var tags []models.Tag
	if err := t.db.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t tagRepositoryImp) GetOne(tagID int) (*models.Tag, error) {
	var tag models.Tag
	if err := t.db.Where("id =?", tagID).First(&tag).Error; err != nil {
		return nil, err
	}

	return &tag, nil
}

func (t tagRepositoryImp) Create(tag models.Tag) error {
	if err := t.db.Create(&tag).Error; err != nil {
		return err
	}
	return nil
}

func (t tagRepositoryImp) Update(tagID int, tag models.Tag) error {
	var updateTag models.Tag
	if err := t.db.Model(&updateTag).Where("id=?", tagID).Updates(&tag).Error; err != nil {
		return err
	}
	return nil
}

func (t tagRepositoryImp) Delete(tagID int) error {
	var tag models.Tag
	if err := t.db.Where("id =?", tagID).Unscoped().Delete(&tag).Error; err != nil {
		return err
	}
	return nil
}
