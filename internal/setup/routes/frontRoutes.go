package routes

import (
	"github.com/gofiber/fiber/v2"
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
	userDataRoutes.Put("/update-profile", userConstructor.UserHandler.UpdateProfile)
	userDataRoutes.Put("/change-password", userConstructor.UserHandler.ChangePassword)
	userDataRoutes.Delete("/delete-profile", userConstructor.UserHandler.DeleteProfile)
	//  user like and write comment routes
	userDataRoutes.Post("/addPostLike", postConstructor.PostHandler.AddUserLikeOfPost)
	userDataRoutes.Post("/addComment", postConstructor.PostHandler.AddComment)

	//1. user like goyan postlaryny gorup bilmeli ---> GET method
	//3. user comment yazan postlaryny gorup bilmeli ---> GET method

	//front ucin public api
	//
	// categories
	//
	//1. get All categories
	//2. get one category with posts
	//
	//
	// posts
	//
	//1. get all posts (id postTitle, postSlug, createdAt, postImage, ClickCount, like count )  with category name
	//2. get one post with comments, tags, category name

}
