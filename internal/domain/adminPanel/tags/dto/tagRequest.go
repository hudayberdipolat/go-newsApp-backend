package dto

type CreateTagRequest struct {
	TagName string `json:"tag_name" validate:"required,min=3"`
}

type UpdateTagRequest struct {
	TagName string `json:"tag_name"`
}
