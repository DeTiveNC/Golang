package service

import (
	"github.com/detivenc/restfulcrudapi/data/request"
	"github.com/detivenc/restfulcrudapi/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}
