package seeder

import (
	"math/rand/v2"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/gosimple/slug"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/utils"
)

func (n newSeederImp) productSeeder() error {
	for i := 1; i <= 200; i++ {
		postTitle := faker.Word()
		randString := utils.RandStringRunes(8)
		imageUrl := "public/postImages/image.jpg"
		post := models.Post{
			PostTitle:  postTitle,
			PostDesc:   faker.Paragraph(),
			PostSlug:   slug.Make(postTitle) + "-" + randString,
			PostStatus: "active",
			ImageUrl:   &imageUrl,
			CategoryID: rand.IntN(4),
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		if err := n.db.Create(&post).Error; err != nil {
			return err
		}
	}
	return nil
}
