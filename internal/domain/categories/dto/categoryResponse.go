package dto

import "github.com/hudayberdipolat/go-newsApp-backend/internal/models"

type CategoryResponse struct {
	Id             int    `json:"id"`
	CategoryName   string `json:"category_name"`
	CategorySlug   string `json:"category_slug"`
	CategoryStatus string `json:"category_status"`
	CreatedAt      string `json:"created_at"`
}

func NewCategoryResponse(category models.Category) CategoryResponse {
	return CategoryResponse{
		Id:             category.ID,
		CategoryName:   category.CategoryName,
		CategorySlug:   category.CategorySlug,
		CategoryStatus: category.CategoryStatus,
		CreatedAt:      category.CreatedAt.Format("01-02-2006"),
	}
}
