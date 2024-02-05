package constructor

import (
	handler2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/categories/handler"
	repository2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/categories/repository"
	service2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/categories/service"
	"gorm.io/gorm"
)

var categoryRepo repository2.CategoryRepository
var categoryService service2.CategoryService
var CategoryHandler handler2.CategoryHandler

func CategoryRequirementCreator(db *gorm.DB) {
	categoryRepo = repository2.NewCategoryRepository(db)
	categoryService = service2.NewCategoryService(categoryRepo)
	CategoryHandler = handler2.NewCategoryHandler(categoryService)
}
