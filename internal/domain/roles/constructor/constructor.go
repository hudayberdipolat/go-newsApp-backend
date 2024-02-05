package constructor

import (
	handler2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/roles/handler"
	repository2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/roles/repository"
	service2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/roles/service"
	"gorm.io/gorm"
)

var roleRepo repository2.RoleRepository
var roleService service2.RoleService
var RoleHandler handler2.RoleHandler

func RoleRequirementCreator(db *gorm.DB) {
	roleRepo = repository2.NewRoleRepository(db)
	roleService = service2.NewRoleService(roleRepo)
	RoleHandler = handler2.NewRoleHandler(roleService)
}
