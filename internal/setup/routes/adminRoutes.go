package routes

import (
	"github.com/gofiber/fiber/v2"
	adminConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/admins/constructor"
	authAdminconstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/authAdmin/constructor"
	categoryConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/categories/constructor"
	postConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/posts/constructor"
	roleConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/roles/constructor"
	tagConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/tags/constructor"
)

func AdminRoutes(app *fiber.App) {
	adminApiRoute := app.Group("/api/admin")

	// auth Admin routes

	authAdminRoute := adminApiRoute.Group("/auth")
	authAdminRoute.Post("/login", authAdminconstructor.AuthAdminHandler.Login)

	//category routes
	categoryRoute := adminApiRoute.Group("/categories")
	categoryRoute.Get("/", categoryConstructor.CategoryHandler.GetAll)
	categoryRoute.Get("/:categoryID", categoryConstructor.CategoryHandler.GetOne)
	categoryRoute.Post("/create", categoryConstructor.CategoryHandler.Create)
	categoryRoute.Put("/:categoryID/update", categoryConstructor.CategoryHandler.Update)
	categoryRoute.Delete("/:categoryID/delete", categoryConstructor.CategoryHandler.Delete)

	// tag routes
	tagRoute := adminApiRoute.Group("/tags")
	tagRoute.Get("/", tagConstructor.TagHandler.GetAll)
	tagRoute.Get("/:tagID", tagConstructor.TagHandler.GetOne)
	tagRoute.Post("/create", tagConstructor.TagHandler.Create)
	tagRoute.Put("/:tagID/update", tagConstructor.TagHandler.Update)
	tagRoute.Delete("/:tagID/delete", tagConstructor.TagHandler.Delete)

	// post routes
	postRoute := adminApiRoute.Group("/posts")
	postRoute.Get("/", postConstructor.PostHandler.GetAll)
	postRoute.Get("/:postID", postConstructor.PostHandler.GetOne)
	postRoute.Post("/create", postConstructor.PostHandler.Create)
	postRoute.Put("/:postID/update", postConstructor.PostHandler.Update)
	postRoute.Delete("/:postID/delete", postConstructor.PostHandler.Delete)

	// admin routes
	adminRoute := adminApiRoute.Group("/admins")
	adminRoute.Get("/", adminConstructor.AdminHandler.GetAll)
	adminRoute.Get("/:adminID", adminConstructor.AdminHandler.GetOne)
	adminRoute.Post("/create", adminConstructor.AdminHandler.Create)
	adminRoute.Put("/:adminID/update", adminConstructor.AdminHandler.Update)
	adminRoute.Delete("/:adminID/delete", adminConstructor.AdminHandler.Delete)

	// role routes

	roleRoute := adminApiRoute.Group("/roles")
	roleRoute.Get("/", roleConstructor.RoleHandler.GetAll)
	roleRoute.Get("/:roleID", roleConstructor.RoleHandler.GetOne)
	roleRoute.Post("/create", roleConstructor.RoleHandler.Create)
	roleRoute.Put("/:roleID/update", roleConstructor.RoleHandler.Update)
	roleRoute.Delete("/:roleID/delete", roleConstructor.RoleHandler.Delete)

	// role permission

}
