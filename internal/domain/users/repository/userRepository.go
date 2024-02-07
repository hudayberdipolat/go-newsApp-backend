package repository

import "github.com/hudayberdipolat/go-newsApp-backend/internal/models"

type UserRepository interface {
	GetUserByPhoneNumber(phoneNumber string) (*models.User, error)
	GetUserByID(userID int) (*models.User, error)
	GetUserData(userID int, phoneNumber string) (*models.User, error)
	Create(user models.User) error
	Update(userID int, user models.User) error
	Delete(userID int, phoneNumber string) error
	ChangeUserPassword(userID int, password string) error

	// function for admin panel
	GetAllUsers() ([]models.User, error)
	GetOneUser(userID int) (*models.User, error)
}
