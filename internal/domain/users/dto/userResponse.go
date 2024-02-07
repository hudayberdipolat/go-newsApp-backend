package dto

import "github.com/hudayberdipolat/go-newsApp-backend/internal/models"

type UserResponse struct {
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func NewUserResponse(user *models.User) UserResponse {
	return UserResponse{
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt.Format("01-02-2006"),
		UpdatedAt:   user.UpdatedAt.Format("01-02-2006"),
	}
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

type GetAllUsersResponse struct {
	ID               int    `json:"id"`
	FullName         string `json:"full_name"`
	PhoneNumber      string `json:"phone_number"`
	UserStatus       string `json:"user_status"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	UserLikedCount   int    `json:"user_liked_count"`
	UserCommentCount int    `json:"user_comment_count"`
}

func NewGetAllUsersResponse(users []models.User) []GetAllUsersResponse {
	var getAllUserResponses []GetAllUsersResponse
	for _, user := range users {
		getUserResponse := GetAllUsersResponse{
			ID:               user.ID,
			FullName:         user.FullName,
			PhoneNumber:      user.PhoneNumber,
			UserStatus:       user.UserStatus,
			CreatedAt:        user.CreatedAt.Format("01-02-2006"),
			UpdatedAt:        user.UpdatedAt.Format("01-02-2006"),
			UserLikedCount:   0,
			UserCommentCount: 0,
		}
		getAllUserResponses = append(getAllUserResponses, getUserResponse)
	}
	return getAllUserResponses
}

type GetUserResponse struct {
	ID               int                `json:"id"`
	FullName         string             `json:"full_name"`
	PhoneNumber      string             `json:"phone_number"`
	UserStatus       string             `json:"user_status"`
	CreatedAt        string             `json:"created_at"`
	UpdatedAt        string             `json:"updated_at"`
	UserLikedCount   int                `json:"user_liked_count"`
	UserCommentCount int                `json:"user_comment_count"`
	UserLikedPosts   []userLikedPost    `json:"user_liked_posts"`
	UserWroteComment []userWroteComment `json:"user_wrote_comment"`
}

type userLikedPost struct {
	ID        int
	PostTitle string
	PostSlug  string
	ImageUrl  *string
}

type userWroteComment struct {
	ID               int     `json:"id"`
	PostTitle        string  `json:"post_title"`
	PostSlug         string  `json:"post_slug,omitempty"`
	ImageUrl         *string `json:"image_url,omitempty"`
	UserWroteComment string  `json:"user_wrote_comment"`
}

func NewGetUserResponse(user models.User) GetUserResponse {
	return GetUserResponse{
		ID:               user.ID,
		FullName:         user.FullName,
		PhoneNumber:      user.PhoneNumber,
		UserStatus:       user.UserStatus,
		CreatedAt:        user.CreatedAt.Format("01-02-2006"),
		UpdatedAt:        user.UpdatedAt.Format("01-02-2006"),
		UserLikedCount:   0,
		UserCommentCount: 0,
		UserLikedPosts:   nil,
		UserWroteComment: nil,
	}
}
