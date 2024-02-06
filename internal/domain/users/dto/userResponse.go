package dto

import "github.com/hudayberdipolat/go-newsApp-backend/internal/models"

type UserResponse struct {
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type AuthUserResponse struct {
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	AccessToken string `json:"access_token"`
}

func NewAuthUserResponse(user *models.User, accessToken string) AuthUserResponse {
	return AuthUserResponse{
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt.Format("01-02-2006"),
		UpdatedAt:   user.UpdatedAt.Format("01-02-2006"),
		AccessToken: accessToken,
	}
}

func NewUserResponse(user *models.User) UserResponse {
	return UserResponse{
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt.Format("01-02-2006"),
		UpdatedAt:   user.UpdatedAt.Format("01-02-2006"),
	}
}
