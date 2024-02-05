package dto

import "github.com/hudayberdipolat/go-newsApp-backend/internal/models"

type TagResponse struct {
	ID      int    `json:"id"`
	TagName string `json:"tag_name"`
}

func NewTagResponse(tag models.Tag) TagResponse {
	return TagResponse{
		ID:      tag.ID,
		TagName: tag.TagName,
	}
}
