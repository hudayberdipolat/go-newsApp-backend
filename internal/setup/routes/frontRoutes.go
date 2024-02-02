package routes

import (
	"github.com/gofiber/fiber/v2"
	userConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/front/users/constructor"
)

func FrontRoutes(app *fiber.App) {
	frontApiRoute := app.Group("api/front")

	// userAuth routes
	userAuthRoute := frontApiRoute.Group("/auth")
	userAuthRoute.Post("/register", userConstructor.UserHandler.Register)
	userAuthRoute.Post("/login", userConstructor.UserHandler.Login)

	// user data Routes
	userDataRoutes := frontApiRoute.Group("/user-profile")
	userDataRoutes.Get("/", userConstructor.UserHandler.GetProfileData)
	userDataRoutes.Put("/update-profile", userConstructor.UserHandler.UpdateProfile)
	userDataRoutes.Put("/change-password", userConstructor.UserHandler.ChangePassword)
	userDataRoutes.Delete("/delete-profile", userConstructor.UserHandler.DeleteProfile)
}
