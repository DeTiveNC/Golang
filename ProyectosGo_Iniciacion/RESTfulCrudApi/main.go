package main

import (
	"github.com/detivenc/restfulcrudapi/config"
	"github.com/detivenc/restfulcrudapi/controller"
	"github.com/detivenc/restfulcrudapi/initializers"
	"github.com/detivenc/restfulcrudapi/model"
	"github.com/detivenc/restfulcrudapi/repository"
	"github.com/detivenc/restfulcrudapi/routes"
	"github.com/detivenc/restfulcrudapi/service"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func init() {
	initializers.CollectEnv()
}

func main() {
	log.Info().Msg("Started Server")
	//Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})
	// Repository
	tagRepository := repository.NewTagsRepositoryImp(db)

	// Service
	tagService := service.NewTagsServiceImp(tagRepository, validate)

	// Controller
	tagsController := controller.NewTagsController(tagService)

	routerGin := routes.NewRouter(tagsController)

	err := routerGin.Run()
	if err != nil {
		log.Err(err)
	}
}
