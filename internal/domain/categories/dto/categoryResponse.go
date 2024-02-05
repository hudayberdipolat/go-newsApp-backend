package dto

import (
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/posts/dto"
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

type CategoryResponse struct {
	Id             int    `json:"id"`
	CategoryName   string `json:"category_name"`
	CategorySlug   string `json:"category_slug"`
	CategoryStatus string `json:"category_status"`
	CreatedAt      string `json:"created_at"`
	Posts          []dto.PostResponse
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

func NewOneCategoryResponse(category *models.Category) CategoryResponse {
	var postResponses []dto.PostResponse
	for _, post := range category.Posts {
		postResponse := dto.NewPostResponse(post)
		postResponses = append(postResponses, postResponse)
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
