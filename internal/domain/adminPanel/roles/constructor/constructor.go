package constructor

import (
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/roles/handler"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/roles/repository"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/roles/service"
	"gorm.io/gorm"
)

var roleRepo repository.RoleRepository
var roleService service.RoleService
var RoleHandler handler.RoleHandler

func RoleRequirementCreator(db *gorm.DB) {
	roleRepo = repository.NewRoleRepository(db)
	roleService = service.NewRoleService(roleRepo)
	RoleHandler = handler.NewRoleHandler(roleService)
}
