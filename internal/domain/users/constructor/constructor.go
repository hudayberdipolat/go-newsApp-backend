package constructor

import (
	handler2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/users/handler"
	repository2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/users/repository"
	service2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/users/service"
	"gorm.io/gorm"
)

var userRepo repository2.UserRepository
var userService service2.UserService
var UserHandler handler2.UserHandler

func UserRequirementCreator(db *gorm.DB) {
	userRepo = repository2.NewUserRepository(db)
	userService = service2.NewUserService(userRepo)
	UserHandler = handler2.NewUserHandler(userService)
}
