package constructor

import (
	"github.com/hudayberdipolat/go-newsApp-backend/internal/app"
	adminConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/admins/constructor"
	categoryConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/categories/constructor"
	postConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/posts/constructor"
	tagConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/tags/constructor"
)

func Build(dependencies *app.Dependencies) {
	categoryConstructor.CategoryRequirementCreator(dependencies.DB)
	tagConstructor.TagRequirementCreator(dependencies.DB)
	postConstructor.PostRequirementCreator(dependencies.DB, dependencies.Config)
	adminConstructor.AdminRequirementCreator(dependencies.DB)
}
