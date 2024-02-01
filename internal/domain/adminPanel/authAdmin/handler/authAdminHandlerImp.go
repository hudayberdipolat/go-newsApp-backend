package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/authAdmin/dto"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/authAdmin/service"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/utils/validate"
	"net/http"
)

type authAdminHandlerImp struct {
	authAdminService service.AuthAdminService
}

func NewAuthAdminHandler(service service.AuthAdminService) AuthAdminHandler {
	return authAdminHandlerImp{
		authAdminService: service,
	}

}

func (a authAdminHandlerImp) Login(ctx *fiber.Ctx) error {
	var loginAdminRequest dto.AdminLoginRequest

	// body parser
	if err := ctx.BodyParser(&loginAdminRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate
	if err := validate.ValidateStruct(&loginAdminRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	loginUser, err := a.authAdminService.LoginAdmin(loginAdminRequest)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "error login data", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "admin logged successfully", loginUser)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
