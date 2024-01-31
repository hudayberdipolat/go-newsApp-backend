package service

import (
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/tags/dto"
)

type TagService interface {
	FindAll() ([]dto.TagResponse, error)
	FindOne(tagID int) (*dto.TagResponse, error)
	Create(request dto.CreateTagRequest) error
	Update(tagID int, request dto.UpdateTagRequest) error
	Delete(tagID int) error
}
