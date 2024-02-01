package seeder

import (
	"github.com/gosimple/slug"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type NewSeederInterface interface {
	Seeder() error
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
	adminDefaultPassword, _ := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
	// superAdmin
	superAdmin := models.Admin{
		FullName:    "Hudayberdi Polatov",
		PhoneNumber: "99365097512",
		AdminRole:   "super_admin",
		Password:    string(adminDefaultPassword),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := n.db.Create(&superAdmin).Error; err != nil {
		return err
	}

	// admins
	admins := []models.Admin{
		{
			FullName:    "Gulnara Polatova",
			PhoneNumber: "99365021065",
			AdminRole:   "admin",
			Password:    string(adminDefaultPassword),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			FullName:    "Nurgeldi Polatov",
			PhoneNumber: "99365277904",
			AdminRole:   "admin",
			Password:    string(adminDefaultPassword),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	if err := n.db.Create(&admins).Error; err != nil {
		return err
	}

	// user

	userDefaultPassword, _ := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
	users := []models.User{
		{
			FullName:    "Yrysgal Garajayew",
			PhoneNumber: "99365010203",
			UserStatus:  "active",
			Password:    string(userDefaultPassword),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			FullName:    "Nuryagdy Jumayew",
			PhoneNumber: "99365777777",
			UserStatus:  "active",
			Password:    string(userDefaultPassword),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	if err := n.db.Create(&users).Error; err != nil {
		return err
	}

	// categories
	categoryNames := []string{"sport", "ylym bilim", "medeniyet", "saglyk"}
	var categories []models.Category
	for _, categoryName := range categoryNames {
		category := models.Category{
			CategoryName:   categoryName,
			CategoryStatus: "passive",
			CategorySlug:   slug.Make(categoryName),
			CreatedAt:      time.Now(),
		}

		categories = append(categories, category)

	}
	if err := n.db.Create(&categories).Error; err != nil {
		return err
	}

	// tags
	tagNames := []string{"sport", "ylym-bilim", "medeniyet", "saglyk"}

	for _, tagName := range tagNames {
		tag := models.Tag{
			TagName: tagName,
		}
		if err := n.db.Create(&tag).Error; err != nil {
			return err
		}
	}
	return nil
}
