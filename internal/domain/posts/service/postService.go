package service

import (
	"github.com/gofiber/fiber/v2"
	dto2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/posts/dto"
	"github.com/hudayberdipolat/go-newsApp-backend/pkg/config"
)

type PostService interface {
	FindAll() ([]dto2.PostResponse, error)
	FindOne(postID int) (*dto2.PostResponse, error)
	Create(ctx *fiber.Ctx, config config.Config, request dto2.CreatePostRequest) error
	Update(ctx *fiber.Ctx, config config.Config, postID int, request dto2.UpdatePostRequest) error
	Delete(postID int) error
}
