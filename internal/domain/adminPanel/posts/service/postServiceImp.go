package service

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/posts/dto"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/posts/repository"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/utils"
	"github.com/hudayberdipolat/go-newsApp-backend/pkg/config"
	"time"
)

type postServiceImp struct {
	postRepo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return postServiceImp{
		postRepo: repo,
	}
}

func (p postServiceImp) FindAll() ([]dto.PostResponse, error) {
	posts, err := p.postRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var postResponses []dto.PostResponse
	for _, post := range posts {
		postResponse := dto.NewPostResponse(post)
		postResponses = append(postResponses, postResponse)
	}
	return postResponses, nil
}

func (p postServiceImp) FindOne(postID int) (*dto.PostResponse, error) {
	post, err := p.postRepo.GetOne(postID)
	if err != nil {
		return nil, err
	}

	postResponse := dto.NewPostResponse(*post)
	return &postResponse, nil
}

func (p postServiceImp) Create(ctx *fiber.Ctx, config config.Config, request dto.CreatePostRequest) error {
	// image upload
	path, err := utils.UploadFile(ctx, "image_url", config.PublicPath, "postImages")
	if err != nil {
		return err
	}
	// image upload end

	randString := utils.RandStringRunes(8)
	if request.PostStatus == "" {
		request.PostStatus = "draft"
	}
	createPost := models.Post{
		PostTitle:  request.PostTitle,
		PostSlug:   slug.Make(request.PostTitle) + "-" + randString,
		PostDesc:   request.PostDesc,
		PostStatus: request.PostStatus,
		ImageUrl:   path,
		CategoryID: request.CategoryID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := p.postRepo.Create(createPost); err != nil {
		return err
	}
	return nil
}

func (p postServiceImp) Update(ctx *fiber.Ctx, config config.Config, postID int, request dto.UpdatePostRequest) error {
	updatePost, err := p.postRepo.GetOne(postID)
	if err != nil {
		return errors.New("post not found")
	}

	file, _ := ctx.FormFile("image_url")
	if file != nil {
		//old_image delete
		if errOldImageDelete := utils.DeleteFile(*updatePost.ImageUrl); errOldImageDelete != nil {
			return errOldImageDelete
		}
		//new image upload
		path, errFileUpload := utils.UploadFile(ctx, "image_url", config.PublicPath, "postImages")
		if errFileUpload != nil {
			return errFileUpload
		}
		request.ImageUrl = path
	}

	randString := utils.RandStringRunes(8)
	updatePost.PostTitle = request.PostTitle
	updatePost.PostSlug = slug.Make(request.PostTitle) + "-" + randString
	updatePost.PostDesc = request.PostDesc
	updatePost.ImageUrl = request.ImageUrl
	updatePost.PostStatus = request.PostStatus
	updatePost.CategoryID = request.CategoryID
	updatePost.UpdatedAt = time.Now()

	if errUpdate := p.postRepo.Update(updatePost.ID, *updatePost); errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (p postServiceImp) Delete(postID int) error {
	deletePost, err := p.postRepo.GetOne(postID)
	if err != nil {
		return errors.New("post not found")
	}

	//post image delete
	if errImageDelete := utils.DeleteFile(*deletePost.ImageUrl); err != nil {
		return errImageDelete
	}
	if errDelete := p.postRepo.Delete(deletePost.ID); errDelete != nil {
		return errDelete
	}
	return nil
}
