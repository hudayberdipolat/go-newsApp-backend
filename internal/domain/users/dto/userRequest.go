package dto

type RegisterUserRequest struct {
	FullName        string `json:"full_name" validate:"required,min=5"`
	PhoneNumber     string `json:"phone_number" validate:"required"`
	Password        string `json:"password" validate:"required,min=5"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type LoginUserRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type ChangeUserData struct {
	FullName    string `json:"full_name" validate:"required,min=5"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type ChangeUserPassword struct {
	OldPassword     string `json:"old_password" validate:"required"`
	Password        string `json:"password" validate:"required,min=5"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type UpdateUserStatusRequest struct {
	UserStatus string `json:"user_status" validate:"required"`
}
