package repository

import (
	"errors"
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
	var posts []models.Post
	if err := p.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (p postRepositoryImp) GetOne(postID int) (*models.Post, error) {
	var post models.Post
	if err := p.db.Where("id =?", postID).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (p postRepositoryImp) Create(post models.Post) error {
	if err := p.db.Create(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("Bu post Title adyny ulanyp bilmeyarsiniz!!!")
		}
		return err
	}
	return nil
}

func (p postRepositoryImp) Update(postID int, post models.Post) error {
	var postModel models.Post
	if err := p.db.Model(&postModel).Where("id =?", postID).Updates(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("Bu post Title adyny ulanyp bilmeyarsiniz!!!")
		}
		return err
	}
	return nil
}

func (p postRepositoryImp) Delete(postID int) error {
	var post models.Post
	if err := p.db.Where("id=?", postID).Unscoped().Delete(&post).Error; err != nil {
		return err
	}
	return nil
}
