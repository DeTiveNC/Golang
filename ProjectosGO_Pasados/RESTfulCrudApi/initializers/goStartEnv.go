package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func CollectEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[Error] Not charging the env")
	}
}
