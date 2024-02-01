package constructor

import (
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/authAdmin/handler"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/authAdmin/repository"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/authAdmin/service"
	"gorm.io/gorm"
)

var authAdminRepo repository.AuthAdminRepository
var authAdminService service.AuthAdminService
var AuthAdminHandler handler.AuthAdminHandler

func AuthAdminRequirementCreator(db *gorm.DB) {
	authAdminRepo = repository.NewAuthAdminRepository(db)
	authAdminService = service.NewAuthAdminService(authAdminRepo)
	AuthAdminHandler = handler.NewAuthAdminHandler(authAdminService)
}
