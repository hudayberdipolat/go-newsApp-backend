package service

import (
	dto2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/tags/dto"
)

type TagService interface {
	FindAll() ([]dto2.TagResponse, error)
	FindOne(tagID int) (*dto2.TagResponse, error)
	Create(request dto2.CreateTagRequest) error
	Update(tagID int, request dto2.UpdateTagRequest) error
	Delete(tagID int) error
}
