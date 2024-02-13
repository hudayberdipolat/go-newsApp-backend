package models

import "time"

type Post struct {
	ID         int               `json:"id"`
	PostTitle  string            `json:"post_title"`
	PostSlug   string            `json:"post_slug"`
	PostDesc   string            `json:"post_desc"`
	ClickCount int               `json:"click_count"`
	PostStatus string            `json:"post_status"`
	ImageUrl   *string           `json:"image_url"`
	CategoryID int               `json:"category_id"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
	Category   Category          `json:"category"`
	PostTags   []Tag             `json:"post_tags" gorm:"many2many:post_tags;"`
	Comments   []UserCommentPost `json:"comments"`
	Liked      []UserLikedPost   `json:"likes" `
}

func (*Post) TableName() string {
	return "posts"
}

type UserLikedPost struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	PostID   int    `json:"post_id"`
	LikeType string `json:"like_type"`
}

func (*UserLikedPost) TableName() string {
	return "user_liked_posts"
}

type UserCommentPost struct {
	ID          int       `json:"id"`
	PostID      int       `json:"post_id"`
	UserID      int       `json:"user_id"`
	PostComment string    `json:"post_comment" gorm:"column:user_comment"`
	CreatedAt   time.Time `json:"created_at"`
}

func (*UserCommentPost) TableName() string {
	return "user_comment_post"
}

type PostTag struct {
	ID     int `json:"id"`
	PostID int `json:"post_id"`
	TagID  int `json:"tag_id"`
}

func (*PostTag) TableName() string {
	return "post_tags"
}
