package repository

import (
	"errors"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
	"gorm.io/gorm"
)

type roleRepositoryImp struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return roleRepositoryImp{
		db: db,
	}
}

func (r roleRepositoryImp) GetAll() ([]models.Role, error) {
	var roles []models.Role
	if err := r.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r roleRepositoryImp) GetOne(roleID int) (*models.Role, error) {
	var role models.Role

	if err := r.db.Where("id", roleID).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r roleRepositoryImp) Create(role models.Role) error {
	if err := r.db.Create(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("Bu role ady eýýäm ulanylýar !!!")
		}
		return err
	}
	return nil
}

func (r roleRepositoryImp) Update(roleID int, role models.Role) error {
	var roleModel models.Role
	err := r.db.Model(&roleModel).Where("id =?", roleID).Updates(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("Bu role ady eýýäm ulanylýar !!!")
		}
		return err
	}
	return nil
}

func (r roleRepositoryImp) Delete(roleID int) error {
	var role models.Role
	if err := r.db.Where("id=?", roleID).Unscoped().Delete(&role).Error; err != nil {
		return err
	}
	return nil
}
