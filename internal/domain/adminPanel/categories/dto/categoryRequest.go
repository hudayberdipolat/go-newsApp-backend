package dto

type CreateCategoryRequest struct {
	CategoryName   string `json:"category_name" validate:"required,min=3" `
	CategoryStatus string `json:"category_status,omitempty"`
}

type UpdateCategoryRequest struct {
	CategoryName   string `json:"category_name" validate:"required,min=3"`
	CategoryStatus string `json:"category_status,omitempty"`
}
