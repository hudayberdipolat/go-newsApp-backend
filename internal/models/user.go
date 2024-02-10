package models

import "time"

type User struct {
	ID              int
	FullName        string            `json:"full_name"`
	PhoneNumber     string            `json:"phone_number"`
	UserStatus      string            `json:"user_status" gorm:"column:status"`
	Password        string            `json:"password"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	UserCommentPost []UserCommentPost `json:"user_comment" gorm:"many2many:user_comment_post"`
	UserLikedPost   []UserLikedPost   `json:"user_liked_post" gorm:"many2many:user_liked_posts"`
}

func (*User) TableName() string {
	return "users"
}
