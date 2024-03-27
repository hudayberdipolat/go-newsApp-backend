package routes

import (
	"github.com/gofiber/fiber/v2"
	adminConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/admins/constructor"
	authAdminconstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/authAdmin/constructor"
	categoryConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/categories/constructor"
	postConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/posts/constructor"
	tagConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/tags/constructor"
	userConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/users/constructor"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/middleware"
)

func AdminRoutes(app *fiber.App) {
	adminApiRoute := app.Group("/api/admin")

	// auth Admin routes
	authAdminRoute := adminApiRoute.Group("/auth")
	authAdminRoute.Post("/login", authAdminconstructor.AuthAdminHandler.Login)

	// admin routes
	adminRoute := adminApiRoute.Group("/admins")
	adminRoute.Use(middleware.SuperAdminMiddleware)
	adminRoute.Get("/", adminConstructor.AdminHandler.GetAll)
	adminRoute.Get("/:adminID", adminConstructor.AdminHandler.GetOne)
	adminRoute.Post("/create", adminConstructor.AdminHandler.Create)
	adminRoute.Put("/:adminID/update", adminConstructor.AdminHandler.Update)
	adminRoute.Delete("/:adminID/delete", adminConstructor.AdminHandler.Delete)

	// functions for user
	userRoutes := adminApiRoute.Group("/users")
	userRoutes.Use(middleware.SuperAdminMiddleware)
	userRoutes.Get("/", userConstructor.UserHandler.GetAllUsers)
	userRoutes.Get("/:userID", userConstructor.UserHandler.GetOneUser)
	userRoutes.Put("/:userID/update-status", userConstructor.UserHandler.UpdateUserStatus)

	//category routes
	categoryRoute := adminApiRoute.Group("/categories")
	// categoryRoute.Use(middleware.AdminMiddleware)
	categoryRoute.Get("/", categoryConstructor.CategoryHandler.GetAll)
	categoryRoute.Get("/:categoryID", categoryConstructor.CategoryHandler.GetOne)
	categoryRoute.Post("/create", categoryConstructor.CategoryHandler.Create)
	categoryRoute.Get("/:categoryID/edit", categoryConstructor.CategoryHandler.Edit)
	categoryRoute.Put("/:categoryID/update", categoryConstructor.CategoryHandler.Update)
	categoryRoute.Delete("/:categoryID/delete", categoryConstructor.CategoryHandler.Delete)

	// tag routes
	tagRoute := adminApiRoute.Group("/tags")
	// categoryRoute.Use(middleware.AdminMiddleware)
	tagRoute.Get("/", tagConstructor.TagHandler.GetAll)
	tagRoute.Get("/:tagID", tagConstructor.TagHandler.GetOne)
	tagRoute.Post("/create", tagConstructor.TagHandler.Create)
	tagRoute.Put("/:tagID/update", tagConstructor.TagHandler.Update)
	tagRoute.Delete("/:tagID/delete", tagConstructor.TagHandler.Delete)

	// post routes
	postRoute := adminApiRoute.Group("/posts")
	// postRoute.Use(middleware.AdminMiddleware)
	postRoute.Get("/", postConstructor.PostHandler.GetAll)
	postRoute.Get("/:postID", postConstructor.PostHandler.GetOne)
	postRoute.Post("/create", postConstructor.PostHandler.Create)
	postRoute.Put("/:postID/update", postConstructor.PostHandler.Update)
	postRoute.Delete("/:postID/delete", postConstructor.PostHandler.Delete)
	// add tag for post
	postRoute.Post("/:postID/tags/create", postConstructor.PostHandler.AddTagForPost)

}
