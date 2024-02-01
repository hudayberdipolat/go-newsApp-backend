package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/posts/dto"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/posts/service"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/utils"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/utils/validate"
	"github.com/hudayberdipolat/go-newsApp-backend/pkg/config"
	"net/http"
	"strconv"
)

type postHandlerImp struct {
	postService service.PostService
	config      config.Config
}

func NewPostHandler(service service.PostService, conf config.Config) PostHandler {
	return postHandlerImp{
		postService: service,
		config:      conf,
	}
}

func (p postHandlerImp) GetAll(ctx *fiber.Ctx) error {
	posts, err := p.postService.FindAll()
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "posts not found", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get all posts", posts)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (p postHandlerImp) GetOne(ctx *fiber.Ctx) error {
	postID, _ := strconv.Atoi(ctx.Params("postID"))
	post, err := p.postService.FindOne(postID)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "posts not found", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get one post", post)
	return ctx.Status(http.StatusOK).JSON(successResponse)

}

func (p postHandlerImp) Create(ctx *fiber.Ctx) error {
	var createPostRequest dto.CreatePostRequest

	// body parser
	if err := ctx.BodyParser(&createPostRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate
	if err := validate.ValidateStruct(&createPostRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	//HANDLE FILE STARTED
	path, err := utils.UploadFile(ctx, "post_image", p.config.PublicPath, "postImages")
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "Error file upload", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	//HANDLE FILE END
	createPostRequest.ImageUrl = path

	// create post
	if err := p.postService.Create(createPostRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "can't created post", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "post created successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (p postHandlerImp) Update(ctx *fiber.Ctx) error {
	postID, _ := strconv.Atoi(ctx.Params("postID"))
	var updatePostRequest dto.UpdatePostRequest
	// body parser

	if err := ctx.BodyParser(&updatePostRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate
	if err := validate.ValidateStruct(&updatePostRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// post update
	// eger input-da post-image bar bolsa onda post-yn onki bar bolan suratyny ocurmeli
	// we taze post-image upload etmeli
	file, _ := ctx.FormFile("post_image")
	if file != nil {
		//old_image delete
		if errOldImageDelete := utils.DeleteFile(updatePostRequest.OldImage); errOldImageDelete != nil {
			errResponse := response.Error(http.StatusBadRequest, "Error not deleted old image", errOldImageDelete.Error(), nil)
			return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
		}
		//new image upload

		path, errFileUpload := utils.UploadFile(ctx, "post_image", p.config.PublicPath, "postImages")
		if errFileUpload != nil {
			errResponse := response.Error(http.StatusBadRequest, "Error file upload", errFileUpload.Error(), nil)
			return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
		}
		updatePostRequest.ImageUrl = path
	} else {
		updatePostRequest.ImageUrl = &updatePostRequest.OldImage
	}

	//update post
	if err := p.postService.Update(postID, updatePostRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "can't updated post", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "post updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (p postHandlerImp) Delete(ctx *fiber.Ctx) error {
	postID, _ := strconv.Atoi(ctx.Params("postID"))

	if err := p.postService.Delete(postID); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "can't deleted post", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "post deleted successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}