package seeder

import (
	"time"

	"github.com/gosimple/slug"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
)

func (n newSeederImp) categorySeeder() error {
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
	return nil
}
