package repository

import (
	"errors"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
	"gorm.io/gorm"
)

type userRepositoryImp struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepositoryImp{
		db: db,
	}
}

func (u userRepositoryImp) GetUserByPhoneNumber(phoneNumber string) (*models.User, error) {
	var user models.User
	if err := u.db.Where("phone_number=?", phoneNumber).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (u userRepositoryImp) GetUserByID(userID int) (*models.User, error) {
	var user models.User
	if err := u.db.Where("id=?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (u userRepositoryImp) GetUserData(userID int, phoneNumber string) (*models.User, error) {
	var user models.User
	err := u.db.Where("id =?", userID).Where("phone_number=?", phoneNumber).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (u userRepositoryImp) Create(user models.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("Bu Telefon belgisi öň hasaba alynan!!!")
		}
		return err
	}
	return nil
}

func (u userRepositoryImp) Update(userID int, user models.User) error {
	var updateUser models.User
	if err := u.db.Model(&updateUser).Where("id=?", userID).Updates(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("Bu Telefon belgisi öň hem ulanylýar !!!")
		}
		return err
	}
	return nil
}

func (u userRepositoryImp) Delete(userID int, phoneNumber string) error {
	var user models.User
	err := u.db.Where("id =?", userID).Where("phone_number=?", phoneNumber).Unscoped().Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u userRepositoryImp) ChangeUserPassword(userID int, password string) error {
	var user models.User
	err := u.db.Model(&user).Where("id =?", userID).Updates(models.User{Password: password}).Error
	if err != nil {
		return err
	}
	return nil
}

// functions for admin panel

func (u userRepositoryImp) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
