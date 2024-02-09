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

	GetAllPosts() ([]models.Post, error)
	GetOnePost(postSlug string) (*models.Post, error)
	AddLikePost(likePost models.UserLikedPost) error
	AddCommentPost(addComment models.UserCommentPost) error
	GetPostWithIDAndPostSlug(postID int, postSlug string) (*models.Post, error)
}
