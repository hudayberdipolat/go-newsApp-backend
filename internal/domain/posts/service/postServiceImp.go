package service

import (
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	dto "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/posts/dto"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/posts/repository"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/utils"
	"github.com/hudayberdipolat/go-newsApp-backend/pkg/config"
)

type postServiceImp struct {
	postRepo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return postServiceImp{
		postRepo: repo,
	}
}

func (p postServiceImp) FindAll() ([]dto.AllPostResponse, error) {
	posts, err := p.postRepo.GetAll()
	if err != nil {
		return nil, err
	}
	allPostResponses := dto.NewAllPostResponse(posts)
	return allPostResponses, nil
}

func (p postServiceImp) FindOne(postID int) (*dto.OnePostResponse, error) {
	post, err := p.postRepo.GetOne(postID)

	if err != nil {
		return nil, err
	}
	postResponse := dto.NewOnePostResponse(post)
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
			log.Println(errOldImageDelete.Error())
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

func (p postServiceImp) CreateTagForPost(createPostTag dto.CreateTagForPost) error {
	getPost, err := p.postRepo.GetOne(createPostTag.PostID)
	if err != nil {
		return errors.New("post not found")
	}
	getTag, errTag := p.postRepo.GetOneTag(createPostTag.TagID)
	if errTag != nil {
		return errors.New("tag not found")
	}
	tagForPost := models.PostTag{
		PostID: getPost.ID,
		TagID:  getTag.ID,
	}
	if errCreateTagForPost := p.postRepo.CreateTagForPost(tagForPost); errCreateTagForPost != nil {
		return errCreateTagForPost
	}
	return nil
}

func (p postServiceImp) AddLikePost(userID, postID int) error {
	// eger user posta on like goyan bolsa we tazeden like-a bassa onda onki goyan likeni ayyrmaly
	// userin onki we user profile-de userin haysy posta like goyyan bolsa onda sol postlaryn sanawyny select etdirmeli
	likePost := models.UserLikedPost{
		UserID: userID,
		PostID: postID,
	}
	if err := p.postRepo.AddLikePost(likePost); err != nil {
		return err
	}
	return nil
}

func (p postServiceImp) AddCommentPost(postID, userID int, addComment dto.AddCommentPostRequest) error {
	addPostComment := models.UserCommentPost{
		PostID:      postID,
		UserID:      userID,
		PostComment: addComment.PostComment,
	}
	if err := p.postRepo.AddCommentPost(addPostComment); err != nil {
		return err
	}
	return nil
}
