package repository

import "github.com/hudayberdipolat/go-newsApp-backend/internal/models"

type PostRepository interface {
	GetAll() ([]models.Post, error)
	GetOne(postID int) (*models.Post, error)
	Create(post models.Post) error
	Update(postID int, post models.Post) error
	Delete(postID int) error
	GetOneTag(tagID int) (*models.Tag, error)
	CreateTagForPost(postTag models.PostTag) error

	// functions for frontend

	GetAllPosts(page, pageSize int) ([]models.Post, error)
	GetOnePost(postSlug string) (*models.Post, error)
	// add like functions
	AddLikePost(likePost models.UserLikedPost) error
	CheckLikePost(userID, postID int) *models.UserLikedPost

	AddCommentPost(addComment models.UserCommentPost) error
	GetPostWithIDAndPostSlug(postID int, postSlug string) (int, error)
}
