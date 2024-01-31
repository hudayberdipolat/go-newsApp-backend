package service

import "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/posts/dto"

type PostService interface {
	FindAll() ([]dto.PostResponse, error)
	FindOne(postID int) (*dto.PostResponse, error)
	Create(request dto.CreatePostRequest) error
	Update(postID int, request dto.UpdatePostRequest) error
	Delete(postID int) error
}
