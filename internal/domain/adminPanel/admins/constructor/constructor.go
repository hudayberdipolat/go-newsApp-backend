package constructor

import (
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/admins/handler"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/admins/repository"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/admins/service"
	"gorm.io/gorm"
)

var adminRepo repository.AdminRepository
var adminService service.AdminService
var AdminHandler handler.AdminHandler

func AdminRequirementCreator(db *gorm.DB) {
	adminRepo = repository.NewAdminRepository(db)
	adminService = service.NewAdminService(adminRepo)
	AdminHandler = handler.NewAdminHandler(adminService)
}
