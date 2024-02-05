package handler

import "github.com/gofiber/fiber/v2"

type AuthAdminHandler interface {
	Login(ctx *fiber.Ctx) error
}
