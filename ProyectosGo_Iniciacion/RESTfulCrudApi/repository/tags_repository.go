package repository

import (
	"github.com/detivenc/restfulcrudapi/model"
)

type TagsRepository interface {
	Save(tags model.Tags)
	Update(tags model.Tags)
	Delete(tagId int)
	FindById(tagId int) (tags model.Tags, err error)
	FindAll() []model.Tags
}
