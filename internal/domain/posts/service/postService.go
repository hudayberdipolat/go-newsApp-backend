package service

import (
	"github.com/gofiber/fiber/v2"
	dto "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/posts/dto"
	"github.com/hudayberdipolat/go-newsApp-backend/pkg/config"
)

type PostService interface {
	FindAll() ([]dto.AllPostResponse, error)
	FindOne(postID int) (*dto.OnePostResponse, error)
	Create(ctx *fiber.Ctx, config config.Config, request dto.CreatePostRequest) error
	Update(ctx *fiber.Ctx, config config.Config, postID int, request dto.UpdatePostRequest) error
	Delete(postID int) error
	CreateTagForPost(createPostTag dto.CreateTagForPost) error

	// functions for frontend

	GetAllPosts(category string, page, pageSize int) ([]dto.GetAllPostsResponse, error)
	GetOnePost(postSlug string) (*dto.GetOnePostResponse, error)
	AddLikePost(userID int, postSlug string, addLike dto.AddLike) error
	AddCommentPost(userID int, postSlug string, addComment dto.AddCommentPostRequest) error
}
