package models

import "time"

type User struct {
	ID          int
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`
	UserStatus  string    `json:"user_status" gorm:"column:status"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (*User) TableName() string {
	return "users"
}
