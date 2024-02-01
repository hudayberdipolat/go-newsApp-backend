package repository

import (
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
	"gorm.io/gorm"
)

type adminRepositoryImp struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return adminRepositoryImp{
		db: db,
	}
}

func (a adminRepositoryImp) GetAll() ([]models.Admin, error) {
	var admins []models.Admin
	if err := a.db.Find(&admins).Error; err != nil {
		return nil, err
	}
	return admins, nil
}

func (a adminRepositoryImp) GetOne(adminID int) (*models.Admin, error) {
	var admin models.Admin
	if err := a.db.Where("id=?", adminID).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (a adminRepositoryImp) Create(admin models.Admin) error {
	if err := a.db.Create(&admin).Error; err != nil {
		return err
	}
	return nil
}

func (a adminRepositoryImp) Update(adminID int, admin models.Admin) error {
	var adminModel models.Admin
	if err := a.db.Model(&adminModel).Where("id=?", adminID).Updates(&admin).Error; err != nil {
		return err
	}
	return nil
}

func (a adminRepositoryImp) Delete(adminID int) error {
	var admin models.Admin
	if err := a.db.Where("id =?", adminID).Unscoped().Delete(&admin).Error; err != nil {
		return err
	}
	return nil
}
func (a adminRepositoryImp) ChangePassword(adminID int, password string) error {
	var adminModel models.Admin
	err := a.db.Model(&adminModel).Where("id =?", adminID).Updates(models.Admin{Password: password}).Error
	if err != nil {
		return err
	}
	return nil
}
