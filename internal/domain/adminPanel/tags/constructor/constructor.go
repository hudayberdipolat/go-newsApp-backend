package constructor

import (
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/tags/handler"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/tags/repository"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/tags/service"
	"gorm.io/gorm"
)

var tagRepo repository.TagRepository
var tagService service.TagService
var TagHandler handler.TagHandler

func TagRequirementCreator(db *gorm.DB) {
	tagRepo = repository.NewTagRepository(db)
	tagService = service.NewTagService(tagRepo)
	TagHandler = handler.NewTagHandler(tagService)
}
