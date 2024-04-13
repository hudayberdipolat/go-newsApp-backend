package routes

import (
	"github.com/gofiber/fiber/v2"
	categoryConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/categories/constructor"
	postConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/posts/constructor"
	userConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/users/constructor"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/middleware"
)

func FrontRoutes(app *fiber.App) {
	frontApiRoute := app.Group("api/front")

	// userAuth routes
	userAuthRoute := frontApiRoute.Group("/auth")
	userAuthRoute.Post("/register", userConstructor.UserHandler.Register)
	userAuthRoute.Post("/login", userConstructor.UserHandler.Login)

	// user data Routes
	userDataRoutes := frontApiRoute.Group("/user-profile")
	userDataRoutes.Use(middleware.UserMiddleware)
	userDataRoutes.Get("/", userConstructor.UserHandler.GetProfileData)
	// user profile data update functions
	userDataRoutes.Put("/update-profile", userConstructor.UserHandler.UpdateProfile)
	userDataRoutes.Put("/change-password", userConstructor.UserHandler.ChangePassword)

	userDataRoutes.Delete("/delete-profile", userConstructor.UserHandler.DeleteProfile)

	// categories
	categories := frontApiRoute.Group("/categories")
	categories.Get("/", categoryConstructor.CategoryHandler.GetAllCategories)
	categories.Get("/:categorySlug", categoryConstructor.CategoryHandler.GetOneCategory)

	// posts
	posts := frontApiRoute.Group("/posts")
	posts.Get("/:page/:page_size", postConstructor.PostHandler.GetAllPosts)
	posts.Get("/:postSlug", postConstructor.PostHandler.GetOnePost)
	posts.Use(middleware.UserMiddleware)
	// post add comment
	posts.Post("/:postSlug/add-comment", postConstructor.PostHandler.AddComment)
	// post add like or dislike
	posts.Post("/:postSlug/add-like", postConstructor.PostHandler.AddUserLikeOfPost)
}
