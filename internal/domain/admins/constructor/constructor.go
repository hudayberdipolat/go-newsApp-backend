package constructor

import (
	handler2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/admins/handler"
	repository2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/admins/repository"
	service2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/admins/service"
	"gorm.io/gorm"
)

var adminRepo repository2.AdminRepository
var adminService service2.AdminService
var AdminHandler handler2.AdminHandler

func AdminRequirementCreator(db *gorm.DB) {
	adminRepo = repository2.NewAdminRepository(db)
	adminService = service2.NewAdminService(adminRepo)
	AdminHandler = handler2.NewAdminHandler(adminService)
}
