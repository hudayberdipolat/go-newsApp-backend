package service

import (
	dto "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/users/dto"
)

type UserService interface {
	RegisterUser(request dto.RegisterUserRequest) (*dto.AuthUserResponse, error)
	LoginUser(request dto.LoginUserRequest) (*dto.AuthUserResponse, error)
	GetUserData(userID int, phoneNumber string) (*dto.UserResponse, error)
	UpdateUserData(userID int, data dto.ChangeUserData) (*dto.AuthUserResponse, error)
	UpdateUserPassword(userID int, password dto.ChangeUserPassword) error
	DeleteUser(userID int, phoneNumber string) error

	// functions for admin panel

	GetAllUsers() ([]dto.GetAllUsersResponse, error)
	GetOneUser(userID int) (*dto.GetUserResponse, error)
	UpdateUserStatus(userID int, request dto.UpdateUserStatusRequest) error
}
