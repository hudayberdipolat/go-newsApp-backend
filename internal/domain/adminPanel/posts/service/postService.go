package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/posts/dto"
	"github.com/hudayberdipolat/go-newsApp-backend/pkg/config"
)

type PostService interface {
	FindAll() ([]dto.PostResponse, error)
	FindOne(postID int) (*dto.PostResponse, error)
	Create(ctx *fiber.Ctx, config config.Config, request dto.CreatePostRequest) error
	Update(ctx *fiber.Ctx, config config.Config, postID int, request dto.UpdatePostRequest) error
	Delete(postID int) error
}
