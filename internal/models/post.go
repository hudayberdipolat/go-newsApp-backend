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
}

func (*Post) TableName() string {
	return "posts"
}
