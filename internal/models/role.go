package models

type Role struct {
	ID       int    `json:"id"`
	RoleName string `json:"role_name"`
}

func (*Role) TableName() string {
	return "roles"
}
