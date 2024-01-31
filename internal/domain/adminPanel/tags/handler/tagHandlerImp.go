package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/tags/dto"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/tags/service"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/utils/validate"
	"net/http"
	"strconv"
)

type tagHandlerImp struct {
	tagService service.TagService
}

func NewTagHandler(service service.TagService) TagHandler {
	return tagHandlerImp{
		tagService: service,
	}
}

func (t tagHandlerImp) GetAll(ctx *fiber.Ctx) error {
	tags, err := t.tagService.FindAll()
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "tags not found", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get all tags", tags)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (t tagHandlerImp) GetOne(ctx *fiber.Ctx) error {
	tagID, _ := strconv.Atoi(ctx.Params("tagID"))
	tags, err := t.tagService.FindOne(tagID)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "tag not found", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get tag", tags)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (t tagHandlerImp) Create(ctx *fiber.Ctx) error {
	var createTagRequest dto.CreateTagRequest

	// body parser
	if err := ctx.BodyParser(&createTagRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate
	if err := validate.ValidateStruct(&createTagRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// create tag
	if err := t.tagService.Create(createTagRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "can't created tag", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "tag created successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (t tagHandlerImp) Update(ctx *fiber.Ctx) error {
	tagID, _ := strconv.Atoi(ctx.Params("tagID"))
	var updateTagRequest dto.UpdateTagRequest
	// body parser

	if err := ctx.BodyParser(&updateTagRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate
	if err := validate.ValidateStruct(&updateTagRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// update category

	if err := t.tagService.Update(tagID, updateTagRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "can't updated tag", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "tag updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (t tagHandlerImp) Delete(ctx *fiber.Ctx) error {
	tagID, _ := strconv.Atoi(ctx.Params("tagID"))

	if err := t.tagService.Delete(tagID); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "can't deleted tag", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "tag deleted successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
