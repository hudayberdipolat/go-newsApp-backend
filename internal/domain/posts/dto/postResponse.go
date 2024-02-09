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
	TagCount     int                  `json:"tag_count" `
}

func NewAllPostResponse(posts []models.Post) []AllPostResponse {
	var allPostResponses []AllPostResponse
	tagCount := 0
	for _, post := range posts {

		for i := 0; i < len(post.PostTags); i++ {
			tagCount = tagCount + 1
		}
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
			TagCount:   tagCount,
			PostCategory: postCategoryResponse{
				Id:           post.Category.ID,
				CategoryName: post.Category.CategoryName,
				CategorySlug: post.Category.CategorySlug,
			},
		}
		allPostResponses = append(allPostResponses, postResponse)
	}
	return allPostResponses
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

func NewOnePostResponse(post *models.Post) OnePostResponse {
	var postTagResponses []postTagResponse
	for _, postTag := range post.PostTags {
		onePostTagResponse := postTagResponse{
			Id:      postTag.ID,
			TagName: postTag.TagName,
		}
		postTagResponses = append(postTagResponses, onePostTagResponse)
	}
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
			CategorySlug: post.Category.CategorySlug,
		},
		PostTags: postTagResponses,
	}
}

type postCategoryResponse struct {
	Id           int    `json:"id"`
	CategoryName string `json:"category_name"`
	CategorySlug string `josn:"category_slug"`
}

type postTagResponse struct {
	Id      int    `json:"id"`
	TagName string `json:"tag_name"`
}

// frontend for responses

type GetAllPostsResponse struct {
	ID           int                  `json:"id"`
	PostTitle    string               `json:"post_title"`
	PostSlug     string               `json:"post_slug"`
	ClickCount   int                  `json:"click_count"`
	ImageUrl     *string              `json:"image_url"`
	CreatedAt    string               `json:"created_at"`
	PostCategory postCategoryResponse `json:"post_category"`
}

type GetOnePostResponse struct {
	ID           int                  `json:"id"`
	PostTitle    string               `json:"post_title"`
	PostSlug     string               `json:"post_slug"`
	PostDesc     string               `json:"post_desc"`
	ClickCount   int                  `json:"click_count"`
	ImageUrl     *string              `json:"image_url"`
	CreatedAt    string               `json:"created_at"`
	PostCategory postCategoryResponse `json:"post_category"`
	PostTags     []postTagResponse    `json:"post_tags"`
}

func NewGetAllPostsResponse(posts []models.Post) []GetAllPostsResponse {
	var allPostResponses []GetAllPostsResponse
	for _, post := range posts {
		getPostResponse := GetAllPostsResponse{
			ID:         post.ID,
			PostTitle:  post.PostTitle,
			PostSlug:   post.PostSlug,
			ClickCount: post.ClickCount,
			ImageUrl:   post.ImageUrl,
			CreatedAt:  post.CreatedAt.Format("01-02-2006"),
			PostCategory: postCategoryResponse{
				Id:           post.Category.ID,
				CategoryName: post.Category.CategoryName,
				CategorySlug: post.Category.CategorySlug,
			},
		}

		allPostResponses = append(allPostResponses, getPostResponse)
	}
	return allPostResponses
}

func NewGetOnePostResponse(post *models.Post) GetOnePostResponse {
	var postTagResponses []postTagResponse
	for _, postTag := range post.PostTags {
		onePostTagResponse := postTagResponse{
			Id:      postTag.ID,
			TagName: postTag.TagName,
		}
		postTagResponses = append(postTagResponses, onePostTagResponse)
	}
	return GetOnePostResponse{
		ID:         post.ID,
		PostTitle:  post.PostTitle,
		PostSlug:   post.PostSlug,
		PostDesc:   post.PostDesc,
		ClickCount: post.ClickCount,
		ImageUrl:   post.ImageUrl,
		CreatedAt:  post.CreatedAt.Format("01-02-2006"),
		PostCategory: postCategoryResponse{
			Id:           post.Category.ID,
			CategoryName: post.Category.CategoryName,
			CategorySlug: post.Category.CategorySlug,
		},
		PostTags: postTagResponses,
	}
}
