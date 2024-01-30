package models

type Permission struct {
	ID             int    `json:"id"`
	PermissionName string `json:"permission_name"`
}

func (*Permission) TableName() string {
	return "permissions"
}
