package models

import "time"

type Post struct {
	ID         int       `json:"id"`
	PostTitle  string    `json:"post_title"`
	PostSlug   string    `json:"post_slug"`
	PostDesc   string    `json:"post_desc"`
	ClickCount int       `json:"click_count"`
	PostStatus string    `json:"post_status"`
	ImageUrl   *string   `json:"image_url"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Category   Category  `json:"category"`
	PostTags   []Tag     `json:"post_tags" gorm:"many2many:post_tags;"`
}

type PostTag struct {
	ID     int `json:"id"`
	PostID int `json:"post_id"`
	TagID  int `json:"tag_id"`
}

func (*PostTag) TableName() string {
	return "post_tags"
}

func (*Post) TableName() string {
	return "posts"
}
