package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-newsApp-backend/pkg/config"
)

func SetStaticRoute(app *fiber.App, config *config.Config) {
	app.Static("/public", config.PublicPath)
}
