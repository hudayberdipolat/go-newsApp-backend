package dto

type CreateTagRequest struct {
	TagName string `json:"tag_name"`
}

type UpdateTagRequest struct {
	TagName string `json:"tag_name"`
}
