package app

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func NewApp(dependencies *Dependencies) (httpServer *fiber.App) {
	httpServer = fiber.New(fiber.Config{
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
		AppName:      dependencies.Config.HttpConfig.AppName,
		ServerHeader: dependencies.Config.HttpConfig.AppHeader,
		BodyLimit:    35 * 1024 * 1024,
		ReadTimeout:  time.Minute * 1,
		WriteTimeout: time.Minute * 1,
	})

	// get routes routes

	return httpServer
}
