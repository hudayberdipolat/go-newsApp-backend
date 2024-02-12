package repository

import (
	"errors"
	"log"

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
	if err := p.db.Preload("Category").Preload("PostTags").Preload("Comments").Find(&posts).Error; err != nil {
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

// functions for frontend

// get all posts for frontend
func (p postRepositoryImp) GetAllPosts() ([]models.Post, error) {
	var allPosts []models.Post
	activeStatus := "active"
	if err := p.db.Select("id, post_title, post_slug, image_url, category_id,click_count, created_at ").
		Where("post_status=?", activeStatus).Preload("Category").Find(&allPosts).Error; err != nil {
		return nil, err
	}

	return allPosts, nil

}

// get one post for frontend

func (p postRepositoryImp) GetOnePost(postSlug string) (*models.Post, error) {
	var post models.Post
	activeStatus := "active"
	if err := p.db.Where("post_status=?", activeStatus).Where("post_slug=?", postSlug).
		Preload("Category").Preload("PostTags").First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

// add like post functions

func (p postRepositoryImp) AddLikePost(likePost models.UserLikedPost) error {

	// get post liked by user

	userLikePost := p.CheckLikePost(likePost.UserID, likePost.PostID)
	log.Println("user like post id --->>>", userLikePost.ID)
	if userLikePost.ID != 0 {
		if userLikePost.LikeType == likePost.LikeType {
			// delete user like
			if err := p.db.Delete(&userLikePost).Error; err != nil {
				return err
			}
			return nil
		} else if userLikePost.LikeType != likePost.LikeType {
			// update like
			if err := p.db.Model(&userLikePost).Updates(&likePost).Error; err != nil {
				return err
			}
			return nil
		}
	}

	// create new like
	if err := p.db.Create(&likePost).Error; err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			log.Println(err)
		}
		return err
	}
	return nil
}
func (p postRepositoryImp) CheckLikePost(userID, postID int) *models.UserLikedPost {
	var likePost models.UserLikedPost
	p.db.Where("user_id=?", userID).Where("post_id=?", postID).First(&likePost)
	return &likePost
}

// end add like post functions

func (p postRepositoryImp) AddCommentPost(addComment models.UserCommentPost) error {
	if err := p.db.Create(&addComment).Error; err != nil {

		return err
	}
	return nil
}

// get slug with postID

func (p postRepositoryImp) GetPostWithIDAndPostSlug(postID int, postSlug string) (int, error) {
	var post models.Post
	if err := p.db.Select("posts.id").Where("post_slug=?", postSlug).Where("id=?", postID).First(&post).Error; err != nil {
		return 0, err
	}
	return post.ID, nil
}
