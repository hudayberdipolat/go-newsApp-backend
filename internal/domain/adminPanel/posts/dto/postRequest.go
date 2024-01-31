package dto

type CreatePostRequest struct {
	PostTitle  string  `json:"post_title"`
	PostDesc   string  `json:"post_desc"`
	ImageUrl   *string `json:"image_url"`
	PostStatus string  `json:"post_status"`
}

type UpdatePostRequest struct {
	PostTitle  string  `json:"post_title"`
	PostDesc   string  `json:"post_desc"`
	OldImage   string  `json:"old_image"`
	ImageUrl   *string `json:"image_url"`
	PostStatus string  `json:"post_status"`
}
