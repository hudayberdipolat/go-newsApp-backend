package seeder

import (
	"gorm.io/gorm"
)

type NewSeederInterface interface {
	Seeder() error
	userSeeder() error
	adminSeeder() error
	categorySeeder() error
	tagSeeder() error
}

type newSeederImp struct {
	db *gorm.DB
}

func NewSeeder(db *gorm.DB) NewSeederInterface {
	return newSeederImp{
		db: db,
	}
}

func (n newSeederImp) Seeder() error {
	// admin seeder
	if errAdminSeeder := n.userSeeder(); errAdminSeeder != nil {
		return errAdminSeeder
	}
	// user seeder
	if errUserSeeder := n.userSeeder(); errUserSeeder != nil {
		return errUserSeeder
	}
	// category seeder
	if errCategenderSeeder := n.categorySeeder(); errCategenderSeeder != nil {
		return errCategenderSeeder
	}
	// tag seeder
	if errTagSeeder := n.tagSeeder(); errTagSeeder != nil {
		return errTagSeeder
	}
	// product seeder
	if errProductSeeder := n.productSeeder(); errProductSeeder != nil {
		return errProductSeeder
	}
	return nil
}
