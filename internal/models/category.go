package models

import "time"

type Category struct {
	ID             int       `json:"id"`
	CategoryName   string    `json:"category_name"`
	CategorySlug   string    `json:"category_slug"`
	CreatedAt      time.Time `json:"created_at"`
	CategoryStatus string    `json:"category_status"`
	Posts          []Post    `json:"posts"`
}

func (*Category) TableName() string {
	return "categories"
}
