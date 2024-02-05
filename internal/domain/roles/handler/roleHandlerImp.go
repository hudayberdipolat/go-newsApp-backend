package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/roles/dto"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/roles/service"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/utils/validate"
	"net/http"
	"strconv"
)

type roleHandlerImp struct {
	roleService service.RoleService
}

func NewRoleHandler(service service.RoleService) RoleHandler {
	return roleHandlerImp{
		roleService: service,
	}
}

func (r roleHandlerImp) GetAll(ctx *fiber.Ctx) error {
	roles, err := r.roleService.FindAll()
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "roles not found", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get all roles", roles)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (r roleHandlerImp) GetOne(ctx *fiber.Ctx) error {
	roleID, _ := strconv.Atoi(ctx.Params("roleID"))
	role, err := r.roleService.FindOne(roleID)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "role not found", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get role data", role)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (r roleHandlerImp) Create(ctx *fiber.Ctx) error {
	var createRoleRequest dto.CreateRoleRequest
	// body parser
	if err := ctx.BodyParser(&createRoleRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate
	if err := validate.ValidateStruct(&createRoleRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// create role
	if err := r.roleService.Create(createRoleRequest); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "can't role created", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "role created successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (r roleHandlerImp) Update(ctx *fiber.Ctx) error {
	roleID, _ := strconv.Atoi(ctx.Params("roleID"))
	var updateRoleRequest dto.UpdateRoleRequest
	// body parser
	if err := ctx.BodyParser(&updateRoleRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate
	if err := validate.ValidateStruct(&updateRoleRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// update role data

	if err := r.roleService.Update(roleID, updateRoleRequest); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "role can't updated", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "role updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (r roleHandlerImp) Delete(ctx *fiber.Ctx) error {
	roleID, _ := strconv.Atoi(ctx.Params("roleID"))
	if err := r.roleService.Delete(roleID); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "role can't deleted", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "role deleted successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
