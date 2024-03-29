package handler

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	GetProfileData(ctx *fiber.Ctx) error
	UpdateProfile(ctx *fiber.Ctx) error
	ChangePassword(ctx *fiber.Ctx) error
	DeleteProfile(ctx *fiber.Ctx) error

	// functions for admin panel
	GetAllUsers(ctx *fiber.Ctx) error
	GetOneUser(ctx *fiber.Ctx) error
	UpdateUserStatus(ctx *fiber.Ctx) error
}
