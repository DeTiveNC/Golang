package main

import (
	"github.com/detivenc/restapigo/initializers"
	"github.com/detivenc/restapigo/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.InitializerDB()
}

func main() {
	router := gin.Default()
	routes.UserRouter(router)
	err := router.Run()
	if err != nil {
		return
	}
}
