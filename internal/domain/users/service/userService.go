package service

import (
	dto2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/users/dto"
)

type UserService interface {
	RegisterUser(request dto2.RegisterUserRequest) (*dto2.UserResponse, error)
	LoginUser(request dto2.LoginUserRequest) (*dto2.UserResponse, error)
	GetUserData(userID int, phoneNumber string) (*dto2.UserResponse, error)
	UpdateUserData(userID int, data dto2.ChangeUserData) error
	UpdateUserPassword(userID int, password dto2.ChangeUserPassword) error
	DeleteUser(userID int, phoneNumber string) error
}
