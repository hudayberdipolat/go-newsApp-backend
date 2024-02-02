package service

import "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/front/users/dto"

type UserService interface {
	RegisterUser(request dto.RegisterUserRequest) (*dto.UserResponse, error)
	LoginUser(request dto.LoginUserRequest) (*dto.UserResponse, error)
	GetUserData(userID int, phoneNumber string) (*dto.UserResponse, error)
	UpdateUserData(userID int, data dto.ChangeUserData) error
	UpdateUserPassword(userID int, password dto.ChangeUserPassword) error
	DeleteUser(userID int, phoneNumber string) error
}
