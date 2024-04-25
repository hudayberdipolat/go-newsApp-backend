package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/categories/dto"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/categories/service"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/utils/validate"
)

type categoryHandlerImp struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) CategoryHandler {
	return categoryHandlerImp{
		categoryService: service,
	}
}

func (c categoryHandlerImp) Edit(ctx *fiber.Ctx) error {
	categoryID, _ := strconv.Atoi(ctx.Params("categoryID"))

	category, err := c.categoryService.EditCategory(categoryID)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "category not found", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get category", category)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
func (c categoryHandlerImp) GetAll(ctx *fiber.Ctx) error {
	categories, err := c.categoryService.FindAll()
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "categories not found", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get all categories", categories)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (c categoryHandlerImp) GetOne(ctx *fiber.Ctx) error {
	categoryID, _ := strconv.Atoi(ctx.Params("categoryID"))
	category, err := c.categoryService.FindOne(categoryID)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "category not found", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get category", category)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (c categoryHandlerImp) Create(ctx *fiber.Ctx) error {
	var createCategoryRequest dto.CreateCategoryRequest

	// body parser
	if err := ctx.BodyParser(&createCategoryRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate
	if err := validate.ValidateStruct(&createCategoryRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// create category

	if err := c.categoryService.CreateCategory(createCategoryRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "can't created category", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "category created successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (c categoryHandlerImp) Update(ctx *fiber.Ctx) error {
	categoryID, _ := strconv.Atoi(ctx.Params("categoryID"))
	var updateCategoryRequest dto.UpdateCategoryRequest
	// body parser

	if err := ctx.BodyParser(&updateCategoryRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate
	if err := validate.ValidateStruct(&updateCategoryRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// update category

	if err := c.categoryService.UpdateCategory(categoryID, updateCategoryRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "can't updated category", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "category updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (c categoryHandlerImp) Delete(ctx *fiber.Ctx) error {
	categoryID, _ := strconv.Atoi(ctx.Params("categoryID"))
	if err := c.categoryService.DeleteCategory(categoryID); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "can't deleted category", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "category deleted successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

// functions for frontend

func (c categoryHandlerImp) GetAllCategories(ctx *fiber.Ctx) error {
	allCategories, err := c.categoryService.GetAllCategories()
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "get all category error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get all categories", allCategories)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (c categoryHandlerImp) GetOneCategory(ctx *fiber.Ctx) error {
	categorySlug := ctx.Params("categorySlug")
	getCategory, err := c.categoryService.GetOneCategory(categorySlug)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "get one category error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get one category", getCategory)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
