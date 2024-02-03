package dto

type CreatePostRequest struct {
	PostTitle  string  `json:"post_title" form:"post_title" validate:"required,min=3"`
	PostDesc   string  `json:"post_desc" form:"post_desc" validate:"required"`
	ImageUrl   *string `json:"image_url" form:"image_url" validate:"required"`
	PostStatus string  `json:"post_status" form:"post_status" validate:"required,omitempty"`
	CategoryID int     `json:"category_id" form:"category_id" validate:"required"`
}

type UpdatePostRequest struct {
	PostTitle  string  `json:"post_title" form:"post_title" validate:"required,min=3"`
	PostDesc   string  `json:"post_desc" form:"post_desc" validate:"required,min=3"`
	OldImage   string  `json:"old_image" form:"old_image" validate:"required"`
	ImageUrl   *string `json:"image_url" form:"image_url"`
	PostStatus string  `json:"post_status" form:"post_status" validate:"required,omitempty"`
	CategoryID int     `json:"category_id" form:"category_id" validate:"required" `
}
