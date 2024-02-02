package constructor

import (
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/front/users/handler"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/front/users/repository"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/front/users/service"
	"gorm.io/gorm"
)

var userRepo repository.UserRepository
var userService service.UserService
var UserHandler handler.UserHandler

func UserRequirementCreator(db *gorm.DB) {
	userRepo = repository.NewUserRepository(db)
	userService = service.NewUserService(userRepo)
	UserHandler = handler.NewUserHandler(userService)
}
