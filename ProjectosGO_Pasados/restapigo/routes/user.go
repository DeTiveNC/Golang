package routes

import (
	"github.com/detivenc/restapigo/controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {

	router.GET("/", controller.GetUsers)
	router.POST("/", controller.PostUser)
	router.DELETE("/:id", controller.DeleteUser)
	router.PATCH("/:id", controller.UpdateUser)
}
