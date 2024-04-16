package service

import (
	"log"

	"github.com/detivenc/restfulcrudapi/data/request"
	"github.com/detivenc/restfulcrudapi/data/response"
	"github.com/detivenc/restfulcrudapi/model"
	"github.com/detivenc/restfulcrudapi/repository"
	"github.com/go-playground/validator/v10"
)

type TagsServiceImp struct {
	TagsRepository repository.TagsRepository
	validate       *validator.Validate
}

// Create implements TagsService.
func (t *TagsServiceImp) Create(tags request.CreateTagsRequest) {
	err := t.validate.Struct(tags)
	if err != nil {

		log.Fatal(err)
	}
	tagModel := model.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)
}

// Delete implements TagsService.
func (t *TagsServiceImp) Delete(tagsId int) {
	t.TagsRepository.Delete(tagsId)
}

// FindAll implements TagsService.
func (t *TagsServiceImp) FindAll() []response.TagsResponse {
	result := t.TagsRepository.FindAll()
	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}
	return tags
}

// FindById implements TagsService.
func (t *TagsServiceImp) FindById(tagsId int) response.TagsResponse {
	tagData, err := t.TagsRepository.FindById(tagsId)
	if err != nil {

		log.Fatal(err)
	}

	tagResponse := response.TagsResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}
	return tagResponse
}

// Update implements TagsService.
func (t *TagsServiceImp) Update(tags request.UpdateTagRequest) {
	tagData, err := t.TagsRepository.FindById(tags.Id)
	if err != nil {

		log.Fatal(err)
	}
	tagData.Name = tags.Name
	t.TagsRepository.Update(tagData)
}

func NewTagsServiceImp(tagRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImp{
		TagsRepository: tagRepository,
		validate:       validate,
	}
}
