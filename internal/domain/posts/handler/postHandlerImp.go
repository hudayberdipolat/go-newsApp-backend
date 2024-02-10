package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/posts/dto"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/posts/service"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/utils/validate"
	"github.com/hudayberdipolat/go-newsApp-backend/pkg/config"
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

// functions for admin panel

// get all posts

func (p postHandlerImp) GetAll(ctx *fiber.Ctx) error {
	posts, err := p.postService.FindAll()
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "posts not found", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get all posts", posts)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

// get one post

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

// create post

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

	// create post
	if err := p.postService.Create(ctx, p.config, createPostRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "can't created post", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "post created successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

// update post

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

	//update post
	if err := p.postService.Update(ctx, p.config, postID, updatePostRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "can't updated post", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "post updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

// delete post

func (p postHandlerImp) Delete(ctx *fiber.Ctx) error {
	postID, _ := strconv.Atoi(ctx.Params("postID"))

	if err := p.postService.Delete(postID); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "can't deleted post", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "post deleted successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

// functions for add tag in post

func (p postHandlerImp) AddTagForPost(ctx *fiber.Ctx) error {
	var postTagForPostRequest dto.CreateTagForPost
	postID, _ := strconv.Atoi(ctx.Params("postID"))
	// body parser
	if err := ctx.BodyParser(&postTagForPostRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// validate
	if err := validate.ValidateStruct(&postTagForPostRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	postTag := dto.CreateTagForPost{
		PostID: postID,
		TagID:  postTagForPostRequest.TagID,
	}
	if err := p.postService.CreateTagForPost(postTag); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "something wrong", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "post tag created successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

// functions for admin panel end

// functions for frontend

// get all posts for front

func (p postHandlerImp) GetAllPosts(ctx *fiber.Ctx) error {
	allPosts, err := p.postService.GetAllPosts()
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "not found posts", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get All Posts", allPosts)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

// get one post for front

func (p postHandlerImp) GetOnePost(ctx *fiber.Ctx) error {
	postSlug := ctx.Params("postSlug")

	post, err := p.postService.GetOnePost(postSlug)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "not found post", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get one post", post)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

// add comment  post

func (p postHandlerImp) AddComment(ctx *fiber.Ctx) error {
	// eger-de user comment yazjak bolsa onda post_slug bilen request-de gelyan
	// post_id select edilende post den bolmaly we sona gora select etdirmeli
	userID := ctx.Locals("user_id").(int)
	postSlug := ctx.Params("postSlug")
	var addCommentRequest dto.AddCommentPostRequest
	if err := ctx.BodyParser(&addCommentRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// validate
	if err := validate.ValidateStruct(&addCommentRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	if err := p.postService.AddCommentPost(userID, postSlug, addCommentRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "sorry something wrong", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "Thank you for comment!!!", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

// add like post

func (p postHandlerImp) AddUserLikeOfPost(ctx *fiber.Ctx) error {
	var addLike dto.AddLike
	postSlug := ctx.Params("postSlug")
	userID := ctx.Locals("user_id").(int)

	if err := ctx.BodyParser(&addLike); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// validate
	if err := validate.ValidateStruct(&addLike); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	if err := p.postService.AddLikePost(userID, postSlug, addLike); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "something wrong", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "Thank you for like!!!", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
