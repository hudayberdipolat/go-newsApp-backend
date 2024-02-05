package constructor

import (
	"github.com/hudayberdipolat/go-newsApp-backend/internal/app"
	adminConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/admins/constructor"
	authAdminconstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/authAdmin/constructor"
	categoryConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/categories/constructor"
	postConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/posts/constructor"
	roleConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/roles/constructor"
	tagConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/tags/constructor"
	userConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/users/constructor"
)

func Build(dependencies *app.Dependencies) {

	authAdminconstructor.AuthAdminRequirementCreator(dependencies.DB)
	categoryConstructor.CategoryRequirementCreator(dependencies.DB)
	tagConstructor.TagRequirementCreator(dependencies.DB)
	postConstructor.PostRequirementCreator(dependencies.DB, dependencies.Config)
	adminConstructor.AdminRequirementCreator(dependencies.DB)
	roleConstructor.RoleRequirementCreator(dependencies.DB)
	userConstructor.UserRequirementCreator(dependencies.DB)
}
