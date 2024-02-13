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
	// user data-lary get edilip alnanada user-in yazan commentlerinin sanyny we goyyan like-larynyn sanyny hem yazdyrmaly
	userDataRoutes.Get("/", userConstructor.UserHandler.GetProfileData)
	// user profile data update functions
	userDataRoutes.Put("/update-profile", userConstructor.UserHandler.UpdateProfile)
	userDataRoutes.Put("/change-password", userConstructor.UserHandler.ChangePassword)
	// delete user profile.
	// eger-de user profile-ni pozan yagdayynda user-in goyyan like-lary we yazan comment-leri hem delete edilmeli
	userDataRoutes.Delete("/delete-profile", userConstructor.UserHandler.DeleteProfile)

	// user like goyyan post-larynyn ahlisini gorup bilmeli
	//userDataRoutes.Get("/user-like-post", postConstructor.PostHandler.GetAllLikePostOfUser)

	// user comment yazan postlaryny gorup bilmeli. we yazan comment-ni hem gormeli

	//userDataRoutes.Get("/user-comment-post", postConstructor.PostHandler.GetAllCommentOfUser)

	// categories
	categories := frontApiRoute.Group("/categories")
	categories.Get("/", categoryConstructor.CategoryHandler.GetAllCategories)
	categories.Get("/:categorySlug", categoryConstructor.CategoryHandler.GetOneCategory)

	// posts

	posts := frontApiRoute.Group("/posts")
	// get all posts
	// get all posts edilende like we dislike comment count sanyny hem aldyrmaly
	posts.Get("/", postConstructor.PostHandler.GetAllPosts)
	// get one post with slug .
	// post get edilip alynanda goylan goylan like we dislike sanyny get etdirip aldyrmaly
	// post get edilende click count update edilmeli
	posts.Get("/:postSlug", postConstructor.PostHandler.GetOnePost)

	// user comment yazmak we like goymak ucin  hokmany suratda ulgama giren bolmaly
	posts.Use(middleware.UserMiddleware)
	// user comment yazan-da yzyna return edilende yzyna comment sanyny yazdyrmaly
	posts.Post("/:postSlug/add-comment", postConstructor.PostHandler.AddComment)

	// egerde user post-a like yada dislike goysa onda yzyna like we dislike sanyny return etdirmeli

	posts.Post("/:postSlug/add-like", postConstructor.PostHandler.AddUserLikeOfPost)

	//1. get all posts (id postTitle, postSlug, createdAt, postImage, ClickCount, like count )  with category name
	//2. get one post with comments, tags, category name
}
