package dto

type CreateCategoryRequest struct {
	CategoryName   string `json:"category_name"`
	CategoryStatus string `json:"category_status"`
}

type UpdateCategoryRequest struct {
	CategoryName   string `json:"category_name"`
	CategoryStatus string `json:"category_status"`
}
