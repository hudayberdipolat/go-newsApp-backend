package seeder

import (
	"time"

	"github.com/gosimple/slug"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

	// user

	users := []models.User{
		{
			FullName:    "Yrysgal Garajayew",
			PhoneNumber: "99365010203",
			UserStatus:  "active",
			Password:    string(defaultPassword),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			FullName:    "Nuryagdy Jumayew",
			PhoneNumber: "99365777777",
			UserStatus:  "active",
			Password:    string(defaultPassword),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	if err := n.db.Omit("UserCommentPost.*,UserLikedPost.*").Create(&users).Error; err != nil {
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
	var tags []models.Tag
	for _, tagName := range tagNames {
		tag := models.Tag{
			TagName: tagName,
		}
		tags = append(tags, tag)
	}
	if err := n.db.Create(&tags).Error; err != nil {
		return err
	}
	return nil
}
