package seeder

import (
	"time"

	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (n newSeederImp) userSeeder() error {
	defaultPassword, _ := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
	defaultStatus := "active"
	users := []models.User{
		{
			FullName:    "Yrysgal Garajayew",
			PhoneNumber: "99365010203",
			UserStatus:  defaultStatus,
			Password:    string(defaultPassword),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			FullName:    "Nuryagdy Jumayew",
			PhoneNumber: "99365777777",
			UserStatus:  defaultStatus,
			Password:    string(defaultPassword),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	if err := n.db.Create(&users).Error; err != nil {
		return err
	}
	return nil
}
