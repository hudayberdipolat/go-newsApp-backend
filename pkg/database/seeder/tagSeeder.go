package seeder

import "github.com/hudayberdipolat/go-newsApp-backend/internal/models"

func (n newSeederImp) tagSeeder() error {
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
