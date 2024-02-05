package repository

import "github.com/hudayberdipolat/go-newsApp-backend/internal/models"

type TagRepository interface {
	GetAll() ([]models.Tag, error)
	GetOne(tagID int) (*models.Tag, error)
	Create(tag models.Tag) error
	Update(tagID int, tag models.Tag) error
	Delete(tagID int) error
}
