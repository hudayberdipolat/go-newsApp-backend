package handler

import "github.com/gofiber/fiber/v2"

type PostHandler interface {
	GetAll(ctx *fiber.Ctx) error
	GetOne(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error

	// functions for frontend
	GetAllPosts(ctx *fiber.Ctx) error
	GetOnePost(ctx *fiber.Ctx) error
	AddTagForPost(ctx *fiber.Ctx) error
	AddUserLikeOfPost(ctx *fiber.Ctx) error
	AddComment(ctx *fiber.Ctx) error
}
