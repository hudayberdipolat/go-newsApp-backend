package dto

import (
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
)

type AllPostResponse struct {
	ID           int                  `json:"id"`
	PostTitle    string               `json:"post_title"`
	PostSlug     string               `json:"post_slug"`
	PostDesc     string               `json:"post_desc"`
	ClickCount   int                  `json:"click_count"`
	PostStatus   string               `json:"post_status"`
	CategoryID   int                  `json:"category_id"`
	ImageUrl     *string              `json:"image_url"`
	CreatedAt    string               `json:"created_at"`
	UpdatedAt    string               `json:"updated_at"`
	PostCategory postCategoryResponse `json:"post_category"`
	TagCount     int                  `json:"tag_count"`
}

type OnePostResponse struct {
	ID           int                  `json:"id"`
	PostTitle    string               `json:"post_title"`
	PostSlug     string               `json:"post_slug"`
	PostDesc     string               `json:"post_desc"`
	ClickCount   int                  `json:"click_count"`
	PostStatus   string               `json:"post_status"`
	CategoryID   int                  `json:"category_id"`
	ImageUrl     *string              `json:"image_url"`
	CreatedAt    string               `json:"created_at"`
	UpdatedAt    string               `json:"updated_at"`
	PostCategory postCategoryResponse `json:"post_category"`
	PostTags     []postTagResponse    `json:"post_tags"`
}

type postCategoryResponse struct {
	Id           int    `json:"id"`
	CategoryName string `json:"category_name"`
}

type postTagResponse struct {
	Id      int    `json:"id"`
	TagName string `json:"tag_name"`
}

func NewAllPostResponse(posts []models.Post) []AllPostResponse {
	var allPostResponses []AllPostResponse
	for _, post := range posts {
		postResponse := AllPostResponse{
			ID:         post.ID,
			PostTitle:  post.PostTitle,
			PostSlug:   post.PostSlug,
			PostDesc:   post.PostDesc,
			ClickCount: post.ClickCount,
			PostStatus: post.PostStatus,
			CategoryID: post.CategoryID,
			ImageUrl:   post.ImageUrl,
			CreatedAt:  post.CreatedAt.Format("01-02-2006"),
			UpdatedAt:  post.UpdatedAt.Format("01-02-2006"),
			TagCount:   0,
			PostCategory: postCategoryResponse{
				Id:           post.Category.ID,
				CategoryName: post.Category.CategoryName,
			},
		}
		allPostResponses = append(allPostResponses, postResponse)
	}
	return allPostResponses
}

func NewOnePostResponse(post *models.Post) OnePostResponse {
	return OnePostResponse{
		ID:         post.ID,
		PostTitle:  post.PostTitle,
		PostSlug:   post.PostSlug,
		PostDesc:   post.PostDesc,
		ClickCount: post.ClickCount,
		PostStatus: post.PostStatus,
		CategoryID: post.CategoryID,
		ImageUrl:   post.ImageUrl,
		CreatedAt:  post.CreatedAt.Format("01-02-2006"),
		UpdatedAt:  post.UpdatedAt.Format("01-02-2006"),
		PostCategory: postCategoryResponse{
			Id:           post.Category.ID,
			CategoryName: post.Category.CategoryName,
		},
		PostTags: nil,
	}
}
