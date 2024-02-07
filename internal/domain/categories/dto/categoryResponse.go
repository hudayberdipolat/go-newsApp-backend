package dto

import (
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
)

type CategoryAllResponse struct {
	Id             int    `json:"id"`
	CategoryName   string `json:"category_name"`
	CategorySlug   string `json:"category_slug"`
	CategoryStatus string `json:"category_status"`
	CreatedAt      string `json:"created_at"`
	PostCount      int    `json:"post_count"`
}

func NewAllCategoryResponse(categories []models.Category) []CategoryAllResponse {
	var categoryAllResponses []CategoryAllResponse
	for _, category := range categories {
		postCount := 0
		for i := 0; i < len(category.Posts); i++ {
			postCount = postCount + 1
		}
		categoryAllResponse := CategoryAllResponse{
			Id:             category.ID,
			CategoryName:   category.CategoryName,
			CategorySlug:   category.CategorySlug,
			CategoryStatus: category.CategoryStatus,
			CreatedAt:      category.CreatedAt.Format("01-02-2006"),
			PostCount:      postCount,
		}
		categoryAllResponses = append(categoryAllResponses, categoryAllResponse)
	}
	return categoryAllResponses
}

type CategoryResponse struct {
	Id             int    `json:"id"`
	CategoryName   string `json:"category_name"`
	CategorySlug   string `json:"category_slug"`
	CategoryStatus string `json:"category_status"`
	CreatedAt      string `json:"created_at"`
	Posts          []postResponse
}

func NewOneCategoryResponse(category *models.Category) CategoryResponse {
	var postResponses []postResponse
	for _, post := range category.Posts {
		onePostResponse := newPostResponse(post)
		postResponses = append(postResponses, onePostResponse)
	}
	return CategoryResponse{
		Id:             category.ID,
		CategoryName:   category.CategoryName,
		CategorySlug:   category.CategorySlug,
		CategoryStatus: category.CategoryStatus,
		CreatedAt:      category.CreatedAt.Format("01-02-2006"),
		Posts:          postResponses,
	}
}

type postResponse struct {
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

func newPostResponse(post models.Post) postResponse {
	return postResponse{
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

type GetAllCategoriesResponse struct {
	Id           int    `json:"id"`
	CategoryName string `json:"category_name"`
	CategorySlug string `json:"category_slug"`
	PostCount    int    `json:"post_count"`
}

func NewGetAllCategoriesResponse(categories []models.Category) []GetAllCategoriesResponse {
	var getAllCategories []GetAllCategoriesResponse
	for _, category := range categories {
		postCount := 0
		for i := 0; i < len(category.Posts); i++ {
			postCount = postCount + 1
		}
		categoryResponse := GetAllCategoriesResponse{
			Id:           category.ID,
			CategoryName: category.CategoryName,
			CategorySlug: category.CategorySlug,
			PostCount:    postCount,
		}
		getAllCategories = append(getAllCategories, categoryResponse)
	}
	return getAllCategories
}
