package constructor

import (
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/categories/handler"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/categories/repository"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/categories/service"
	"gorm.io/gorm"
)

var categoryRepo repository.CategoryRepository
var categoryService service.CategoryService
var CategoryHandler handler.CategoryHandler

func CategoryRequirementCreator(db *gorm.DB) {
	categoryRepo = repository.NewCategoryRepository(db)
	categoryService = service.NewCategoryService(categoryRepo)
	CategoryHandler = handler.NewCategoryHandler(categoryService)
}
