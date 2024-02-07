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
	if err := p.db.Preload("Category").Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (p postRepositoryImp) GetOne(postID int) (*models.Post, error) {
	var post models.Post

	err := p.db.Preload("Category").Preload("PostTags").First(&post, postID).Error
	if err != nil {
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

func (p postRepositoryImp) GetOneTag(tagID int) (*models.Tag, error) {
	var tag models.Tag
	if err := p.db.Where("id=?", tagID).First(&tag).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func (p postRepositoryImp) CreateTagForPost(postTag models.PostTag) error {
	err := p.db.Where("post_id = ?", postTag.PostID).Where("tag_id=?", postTag.TagID).FirstOrCreate(&postTag).Error
	if err != nil {
		return err
	}
	return nil
}

func (p postRepositoryImp) AddLikePost(likePost models.UserLikedPost) error {
	if err := p.db.Create(&likePost).Error; err != nil {
		return err
	}
	return nil
}

func (p postRepositoryImp) AddCommentPost(addComment models.UserCommentPost) error {
	if err := p.db.Create(&addComment).Error; err != nil {
		return err
	}
	return nil
}
