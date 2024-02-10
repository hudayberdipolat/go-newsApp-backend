package dto

type CreatePostRequest struct {
	PostTitle  string  `json:"post_title" form:"post_title" validate:"required,min=3"`
	PostDesc   string  `json:"post_desc" form:"post_desc" validate:"required"`
	ImageUrl   *string `json:"image_url" form:"image_url" `
	PostStatus string  `json:"post_status,omitempty" form:"post_status,omitempty" `
	CategoryID int     `json:"category_id" form:"category_id" validate:"required"`
}

type UpdatePostRequest struct {
	PostTitle  string  `json:"post_title" form:"post_title" validate:"required,min=3"`
	PostDesc   string  `json:"post_desc" form:"post_desc" validate:"required,min=3"`
	ImageUrl   *string `json:"image_url" form:"image_url"`
	PostStatus string  `json:"post_status" form:"post_status" validate:"required,omitempty"`
	CategoryID int     `json:"category_id" form:"category_id" validate:"required" `
}

type CreateTagForPost struct {
	PostID int `json:"post_id"`
	TagID  int `json:"tag_id" validate:"required"`
}

type AddCommentPostRequest struct {
	PostID      int    `json:"post_id" validate:"required"`
	PostComment string `json:"post_comment" validate:"required"`
}

type AddLike struct {
	PostID   int    `json:"post_id" validate:"required"`
	LikeType string `json:"like_type" validate:"required"`
}
