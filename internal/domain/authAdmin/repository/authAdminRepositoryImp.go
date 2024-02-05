package repository

import (
	"errors"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
	"gorm.io/gorm"
)

type authAdminRepositoryImp struct {
	db *gorm.DB
}

func NewAuthAdminRepository(db *gorm.DB) AuthAdminRepository {
	return authAdminRepositoryImp{
		db: db,
	}
}

func (a authAdminRepositoryImp) GetAdmin(phoneNumber string) (*models.Admin, error) {
	var admin models.Admin
	if err := a.db.Where("phone_number=?", phoneNumber).First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Phone number ya-da password nadogry!!!")
		}
		return nil, err
	}
	return &admin, nil
}
