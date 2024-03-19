package app

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/setup/routes"
)

func NewApp(dependencies *Dependencies) (httpServer *fiber.App) {
	httpServer = fiber.New(fiber.Config{
		ServerHeader: dependencies.Config.HttpConfig.AppHeader,
		BodyLimit:    35 * 1024 * 1024,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return ctx.Status(code).JSON(fiber.Map{
				"status":  code,
				"message": "Näsazlyk ýüze çykdy, Sonrak synanysyn!!!",
			})
		},
		AppName: dependencies.Config.HttpConfig.AppName,
	})
	httpServer.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, API-KEY, Language, Platform, Authorization, Access-Control-Allow-Origin",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))
	// get routes routes
	routes.AdminRoutes(httpServer)
	routes.FrontRoutes(httpServer)
	routes.SetStaticRoute(httpServer, dependencies.Config)
	return httpServer
}
