package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/users/dto"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/users/service"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/utils/validate"
	"net/http"
)

type userHandlerImp struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return userHandlerImp{
		userService: userService,
	}
}

func (u userHandlerImp) Register(ctx *fiber.Ctx) error {
	var registerUserRequest dto.RegisterUserRequest

	// body parser
	if err := ctx.BodyParser(&registerUserRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate
	if err := validate.ValidateStruct(registerUserRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// phone number validate
	validatePhoneNumber := validate.ValidatePhoneNumber(registerUserRequest.PhoneNumber)
	if !validatePhoneNumber {
		errResponse := response.Error(http.StatusBadRequest, "Nädogry telefon belgi", "Nädogry telefon belgi", nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// register user
	userResponse, err := u.userService.RegisterUser(registerUserRequest)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "user can't registered", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "user registered successfully", userResponse)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (u userHandlerImp) Login(ctx *fiber.Ctx) error {
	var loginUserRequest dto.LoginUserRequest
	// body parser
	if err := ctx.BodyParser(&loginUserRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate
	if err := validate.ValidateStruct(loginUserRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// phone number validate
	validatePhoneNumber := validate.ValidatePhoneNumber(loginUserRequest.PhoneNumber)
	if !validatePhoneNumber {
		errResponse := response.Error(http.StatusBadRequest, "Nädogry telefon belgi", "Nädogry telefon belgi", nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	//login user

	userResponse, err := u.userService.LoginUser(loginUserRequest)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "user can't login", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "user login successfully", userResponse)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (u userHandlerImp) GetProfileData(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(int)
	phoneNumber := ctx.Locals("phone_number").(string)

	userResponse, err := u.userService.GetUserData(userID, phoneNumber)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "user not found", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "user profile data ", userResponse)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (u userHandlerImp) UpdateProfile(ctx *fiber.Ctx) error {
	var updateUserRequest dto.ChangeUserData
	userID := ctx.Locals("user_id").(int)
	// body parser
	if err := ctx.BodyParser(&updateUserRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate
	if err := validate.ValidateStruct(updateUserRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// phone number validate
	validatePhoneNumber := validate.ValidatePhoneNumber(updateUserRequest.PhoneNumber)
	if !validatePhoneNumber {
		errResponse := response.Error(http.StatusBadRequest, "Nädogry telefon belgi", "Nädogry telefon belgi", nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// update user Data

	if err := u.userService.UpdateUserData(userID, updateUserRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "user can't updated", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "user updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (u userHandlerImp) ChangePassword(ctx *fiber.Ctx) error {
	var changeUserPassword dto.ChangeUserPassword
	userID := ctx.Locals("user_id").(int)
	// body parser
	if err := ctx.BodyParser(&changeUserPassword); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate
	if err := validate.ValidateStruct(changeUserPassword); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	if err := u.userService.UpdateUserPassword(userID, changeUserPassword); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "user can't change password", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "user password changed successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (u userHandlerImp) DeleteProfile(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(int)
	phoneNumber := ctx.Locals("phone_number").(string)

	// delete user account
	if err := u.userService.DeleteUser(userID, phoneNumber); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "user account can't deleted", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "Account deleted successfully ", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
