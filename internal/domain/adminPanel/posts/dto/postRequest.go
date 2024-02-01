package dto

type CreatePostRequest struct {
	PostTitle  string  `json:"post_title" validate:"required,min=3"`
	PostDesc   string  `json:"post_desc" validate:"required"`
	ImageUrl   *string `json:"image_url" validate:"required"`
	PostStatus string  `json:"post_status" validate:"required,omitempty"`
}

type UpdatePostRequest struct {
	PostTitle  string  `json:"post_title"validate:"required,min=3"`
	PostDesc   string  `json:"post_desc" validate:"required,min=3"`
	OldImage   string  `json:"old_image" validate:"required"`
	ImageUrl   *string `json:"image_url" validate:"required"`
	PostStatus string  `json:"post_status" validate:"required,omitempty"`
}
