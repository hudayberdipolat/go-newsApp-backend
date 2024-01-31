package dto

import (
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
)

type PostResponse struct {
	ID         int    `json:"id"`
	PostTitle  string `json:"post_title"`
	PostSlug   string `json:"post_slug"`
	PostDesc   string `json:"post_desc"`
	ClickCount int    `json:"click_count"`
	PostStatus string `json:"post_status"`
	ImageUrl   string `json:"image_url"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func NewPostResponse(post models.Post) PostResponse {
	return PostResponse{
		ID:         post.ID,
		PostTitle:  post.PostTitle,
		PostSlug:   post.PostSlug,
		PostDesc:   post.PostDesc,
		ClickCount: post.ClickCount,
		PostStatus: post.PostStatus,
		ImageUrl:   *post.ImageUrl,
		CreatedAt:  post.CreatedAt.Format("01-02-2006"),
		UpdatedAt:  post.UpdatedAt.Format("01-02-2006"),
	}
}
