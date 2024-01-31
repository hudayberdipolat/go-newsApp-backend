package service

import (
	"errors"
	"github.com/gosimple/slug"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/posts/dto"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/posts/repository"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
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

func (p postServiceImp) Create(request dto.CreatePostRequest) error {
	createPost := models.Post{
		PostTitle:  request.PostTitle,
		PostSlug:   slug.Make(request.PostTitle),
		PostDesc:   request.PostDesc,
		PostStatus: request.PostStatus,
		ImageUrl:   request.ImageUrl,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := p.postRepo.Create(createPost); err != nil {
		return err
	}
	return nil
}

func (p postServiceImp) Update(postID int, request dto.UpdatePostRequest) error {
	updatePost, err := p.postRepo.GetOne(postID)
	if err != nil {
		return errors.New("post not found")
	}
	updatePost.PostTitle = request.PostTitle
	updatePost.PostSlug = slug.Make(request.PostTitle)
	updatePost.PostDesc = request.PostDesc
	updatePost.ImageUrl = request.ImageUrl
	updatePost.PostStatus = request.PostStatus
	updatePost.UpdatedAt = time.Now()

	if errUpdate := p.postRepo.Update(postID, *updatePost); errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (p postServiceImp) Delete(postID int) error {
	_, err := p.postRepo.GetOne(postID)
	if err != nil {
		return errors.New("post not found")
	}
	if errDelete := p.postRepo.Delete(postID); errDelete != nil {
		return errDelete
	}
	return nil
}
