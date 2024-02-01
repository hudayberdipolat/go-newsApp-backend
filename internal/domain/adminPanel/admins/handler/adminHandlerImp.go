package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/admins/dto"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/admins/service"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/utils/validate"
	"net/http"
	"strconv"
)

type adminHandlerImp struct {
	adminService service.AdminService
}

func NewAdminHandler(service service.AdminService) AdminHandler {
	return adminHandlerImp{
		adminService: service,
	}
}

func (a adminHandlerImp) GetAll(ctx *fiber.Ctx) error {
	admins, err := a.adminService.FindAll()
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "admins not found", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get all admins", admins)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (a adminHandlerImp) GetOne(ctx *fiber.Ctx) error {
	adminID, _ := strconv.Atoi(ctx.Params("adminID"))
	admin, err := a.adminService.FindOne(adminID)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "admin not found", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get admin data", admin)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (a adminHandlerImp) Create(ctx *fiber.Ctx) error {
	var createAdminRequest dto.CreateAdminRequest

	// body parser
	if err := ctx.BodyParser(&createAdminRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate
	if err := validate.ValidateStruct(createAdminRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	if err := a.adminService.Create(createAdminRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "can't created admin", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "admin created successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (a adminHandlerImp) Update(ctx *fiber.Ctx) error {
	var updateAdminData dto.UpdateAdminRequest
	adminID, _ := strconv.Atoi(ctx.Params("adminID"))
	// body parser
	if err := ctx.BodyParser(&updateAdminData); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate
	if err := validate.ValidateStruct(updateAdminData); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// update Admin data

	if err := a.adminService.Update(adminID, updateAdminData); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "can't updated admin", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "admin updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)

}

func (a adminHandlerImp) Delete(ctx *fiber.Ctx) error {
	adminID, _ := strconv.Atoi(ctx.Params("adminID"))
	if err := a.adminService.Delete(adminID); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "can't deleted admin", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "admin deleted successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
