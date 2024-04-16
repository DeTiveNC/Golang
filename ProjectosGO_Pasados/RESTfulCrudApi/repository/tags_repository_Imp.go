package repository

import (
	"errors"
	"log"

	"github.com/detivenc/restfulcrudapi/data/request"
	"github.com/detivenc/restfulcrudapi/model"
	"gorm.io/gorm"
)

type TagsRepositoryImp struct {
	Db *gorm.DB
}

// Delete
func (t *TagsRepositoryImp) Delete(tagId int) {
	var tags model.Tags
	result := t.Db.Where("id = ?", tagId).Delete(&tags)
	if result.Error != nil {

		log.Fatal(result.Error)
	}
}

// Find All
func (t *TagsRepositoryImp) FindAll() []model.Tags {
	var tags []model.Tags
	result := t.Db.Find(&tags)
	if result.Error != nil {

		log.Fatal(result.Error)
	}
	return tags
}

// FindById
func (t *TagsRepositoryImp) FindById(tagId int) (tags model.Tags, err error) {
	var tag model.Tags
	result := t.Db.Find(&tag, tagId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

// Save
func (t *TagsRepositoryImp) Save(tags model.Tags) {
	result := t.Db.Create(&tags)
	if result.Error != nil {

		log.Fatal(result.Error)
	}
}

// Update
func (t *TagsRepositoryImp) Update(tags model.Tags) {
	var updateTag = request.UpdateTagRequest{
		Id:   tags.Id,
		Name: tags.Name,
	}
	result := t.Db.Model(&tags).Updates(updateTag)
	if result.Error != nil {

		log.Fatal(result.Error)
	}
}

func NewTagsRepositoryImp(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImp{Db: Db}
}
