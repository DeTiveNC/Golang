package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/detivenc/restfulcrudapi/data/request"
	"github.com/detivenc/restfulcrudapi/data/response"
	"github.com/detivenc/restfulcrudapi/service"
	"github.com/gin-gonic/gin"
)

type TagsController struct {
	tagsService service.TagsService
}

func NewTagsController(service service.TagsService) *TagsController {
	return &TagsController{
		tagsService: service,
	}
}

// Create
func (controller *TagsController) Create(ctx *gin.Context) {
	createTagsRequest := request.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&createTagsRequest)
	if err != nil {
		log.Fatal(err)
	}
	controller.tagsService.Create(createTagsRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// Update
func (controller *TagsController) Update(ctx *gin.Context) {
	updateTagsRequest := request.UpdateTagRequest{}
	err := ctx.ShouldBindJSON(&updateTagsRequest)
	if err != nil {
		log.Fatal(err)
	}
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	if err != nil {
		log.Fatal(err)
	}
	updateTagsRequest.Id = id

	controller.tagsService.Update(updateTagsRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// Delete
func (controller *TagsController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	if err != nil {
		log.Fatal(err)
	}
	controller.tagsService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// FindById
func (controller *TagsController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	if err != nil {
		log.Fatal(err)
	}
	tagResponse := controller.tagsService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagResponse,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAll
func (controller *TagsController) FindAll(ctx *gin.Context) {
	tagResponse := controller.tagsService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
