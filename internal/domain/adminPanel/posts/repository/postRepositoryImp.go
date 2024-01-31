package repository

import (
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
	"gorm.io/gorm"
)

type postRepositoryImp struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return postRepositoryImp{
		db: db,
	}
}

func (p postRepositoryImp) GetAll() ([]models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p postRepositoryImp) GetOne(postID int) (*models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p postRepositoryImp) Create(post models.Post) error {
	//TODO implement me
	panic("implement me")
}

func (p postRepositoryImp) Update(postID int, post models.Post) error {
	//TODO implement me
	panic("implement me")
}

func (p postRepositoryImp) Delete(postID int) error {
	//TODO implement me
	panic("implement me")
}
