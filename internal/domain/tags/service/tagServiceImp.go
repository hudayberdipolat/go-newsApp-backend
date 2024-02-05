package service

import (
	"errors"
	dto2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/tags/dto"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/tags/repository"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
)

type tagServiceImp struct {
	tagRepo repository.TagRepository
}

func NewTagService(repo repository.TagRepository) TagService {
	return tagServiceImp{
		tagRepo: repo,
	}
}

func (t tagServiceImp) FindAll() ([]dto2.TagResponse, error) {
	tags, err := t.tagRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var tagResponses []dto2.TagResponse

	for _, tag := range tags {
		tagResponse := dto2.NewTagResponse(tag)
		tagResponses = append(tagResponses, tagResponse)
	}
	return tagResponses, nil
}

func (t tagServiceImp) FindOne(tagID int) (*dto2.TagResponse, error) {
	tag, err := t.tagRepo.GetOne(tagID)
	if err != nil {
		return nil, err
	}
	tagResponse := dto2.NewTagResponse(*tag)
	return &tagResponse, nil
}

func (t tagServiceImp) Create(request dto2.CreateTagRequest) error {
	createTag := models.Tag{
		TagName: request.TagName,
	}
	if err := t.tagRepo.Create(createTag); err != nil {
		return err
	}
	return nil
}

func (t tagServiceImp) Update(tagID int, request dto2.UpdateTagRequest) error {

	updateTag, err := t.tagRepo.GetOne(tagID)
	if err != nil {
		return errors.New("tag not found")
	}
	updateTag.TagName = request.TagName

	if err := t.tagRepo.Update(tagID, *updateTag); err != nil {
		return err
	}
	return nil
}

func (t tagServiceImp) Delete(tagID int) error {
	_, err := t.tagRepo.GetOne(tagID)
	if err != nil {
		return errors.New("tag not found")
	}
	if errDelete := t.tagRepo.Delete(tagID); errDelete != nil {
		return errDelete
	}
	return nil
}
