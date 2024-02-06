package models

type Tag struct {
	ID      int    `json:"id"`
	TagName string `json:"tag_name"`
	//Posts   []Post `json:"posts" gorm:"many2many:post_tags"`
}

func (*Tag) TableName() string {
	return "tags"
}
