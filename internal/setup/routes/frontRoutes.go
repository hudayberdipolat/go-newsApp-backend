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
	userDataRoutes.Put("/update-profile", userConstructor.UserHandler.UpdateProfile)
	userDataRoutes.Put("/change-password", userConstructor.UserHandler.ChangePassword)
	userDataRoutes.Delete("/delete-profile", userConstructor.UserHandler.DeleteProfile)
	//  user like and write comment routes
	userDataRoutes.Post("/addPostLike", postConstructor.PostHandler.AddUserLikeOfPost)
	userDataRoutes.Post("/addComment", postConstructor.PostHandler.AddComment)

	//userDataRoutes.Get("/user-like-post", postConstructor.PostHandler.GetAllLikePostOfUser)
	//userDataRoutes.Get("/user-comment-post", postConstructor.PostHandler.GetAllCommentOfUser)

	// categories

	categories := frontApiRoute.Group("/categories")
	categories.Get("/", categoryConstructor.CategoryHandler.GetAllCategories)
	categories.Get("/:categorySlug", categoryConstructor.CategoryHandler.GetOneCategory)

	// posts

	posts := frontApiRoute.Group("/posts")
	posts.Get("/", postConstructor.PostHandler.GetAllPosts)
	posts.Get("/:postSlug", postConstructor.PostHandler.GetOnePost)
	// comment yazmak we like goymak ucin user hokmany suratda ulgama giren bolmaly

	posts.Use(middleware.UserMiddleware)
	posts.Post("/:postSlug/add-comment", postConstructor.PostHandler.AddComment)
	posts.Post("/:postSlug/add-like", postConstructor.PostHandler.AddUserLikeOfPost)

	//1. get all posts (id postTitle, postSlug, createdAt, postImage, ClickCount, like count )  with category name
	//2. get one post with comments, tags, category name

}
