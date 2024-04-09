package seeder

import (
	"time"

	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (n newSeederImp) adminSeeder() error {

	defaultPassword, _ := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
	defaultStatus := "active"

	// admins
	admins := []models.Admin{
		{
			FullName:    "Hudayberdi Polatov",
			PhoneNumber: "99365010203",
			AdminRole:   "super_admin",
			AdminStatus: defaultStatus,
			Password:    string(defaultPassword),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			FullName:    "Gulnara Polatova",
			PhoneNumber: "99365040506",
			AdminRole:   "admin",
			AdminStatus: defaultStatus,
			Password:    string(defaultPassword),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			FullName:    "Nurgeldi Polatov",
			PhoneNumber: "99365070809",
			AdminRole:   "admin",
			AdminStatus: defaultStatus,
			Password:    string(defaultPassword),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	if err := n.db.Create(&admins).Error; err != nil {
		return err
	}
	return nil
}
