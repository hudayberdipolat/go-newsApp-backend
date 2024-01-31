package routes

import (
	"github.com/gofiber/fiber/v2"
	categoryConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/categories/constructor"
	tagConstructor "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/tags/constructor"
)

func AdminRoutes(app *fiber.App) {
	adminApiRoute := app.Group("/api/admin")

	//category routes
	categoryRoute := adminApiRoute.Group("/categories")
	categoryRoute.Get("/", categoryConstructor.CategoryHandler.GetAll)
	categoryRoute.Get("/:categoryID", categoryConstructor.CategoryHandler.GetOne)
	categoryRoute.Post("/create", categoryConstructor.CategoryHandler.Create)
	categoryRoute.Put("/:categoryID/update", categoryConstructor.CategoryHandler.Update)
	categoryRoute.Get("/", categoryConstructor.CategoryHandler.Delete)

	// tag routes
	tagRoute := adminApiRoute.Group("/tags")
	tagRoute.Get("/", tagConstructor.TagHandler.GetAll)
	tagRoute.Get("/:tagID", tagConstructor.TagHandler.GetOne)
	tagRoute.Post("/create", tagConstructor.TagHandler.Create)
	tagRoute.Put("/:tagID/update", tagConstructor.TagHandler.Update)
	tagRoute.Delete("/:tagID/delete", tagConstructor.TagHandler.Delete)
}
