package constructor

import (
	"github.com/hudayberdipolat/go-newsApp-backend/internal/app"
	categoryConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/categories/constructor"
	tagConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/tags/constructor"
)

func Build(dependencies *app.Dependencies) {
	categoryConstructor.CategoryRequirementCreator(dependencies.DB)
	tagConstructor.TagRequirementCreator(dependencies.DB)
}
