package models

import "time"

type Admin struct {
	ID          int       `json:"id"`
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`
	AdminRole   string    `json:"admin_role"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (*Admin) TableName() string {
	return "admins"
}
