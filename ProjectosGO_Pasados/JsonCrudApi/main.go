package main

import (
	"github.com/detivenc/jsoncrudapi/cmd"
	"github.com/detivenc/jsoncrudapi/controllers"
	"github.com/detivenc/jsoncrudapi/initializers/db"
	"github.com/gin-gonic/gin"
)

func init() {
	cmd.LoadEnv()
	db.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.Run() // listen and serve on 0.0.0.0:8080
}
