package models

type Tag struct {
	ID      int    `json:"id"`
	TagName string `json:"tag_name"`
}

func (*Tag) TableName() string {
	return "tags"
}
