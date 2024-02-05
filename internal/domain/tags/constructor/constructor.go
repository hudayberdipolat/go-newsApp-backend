package constructor

import (
	handler2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/tags/handler"
	repository2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/tags/repository"
	service2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/tags/service"
	"gorm.io/gorm"
)

var tagRepo repository2.TagRepository
var tagService service2.TagService
var TagHandler handler2.TagHandler

func TagRequirementCreator(db *gorm.DB) {
	tagRepo = repository2.NewTagRepository(db)
	tagService = service2.NewTagService(tagRepo)
	TagHandler = handler2.NewTagHandler(tagService)
}
