package routes

import (
	"net/http"

	"github.com/detivenc/restfulcrudapi/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(tagsController *controller.TagsController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome here")
	})
	baseRouter := router.Group("/api")
	tagRouter := baseRouter.Group("/tags")
	tagRouter.GET("", tagsController.FindAll)
	tagRouter.GET("/:tagId", tagsController.FindById)
	tagRouter.POST("", tagsController.Create)
	tagRouter.PATCH("/:tagId", tagsController.Update)
	tagRouter.DELETE("/:tagId", tagsController.Delete)

	return router
}
