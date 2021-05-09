package main

import (
	"os"

	"github.com/mateigraura/wirebo-api/core/utils"
	"github.com/mateigraura/wirebo-api/router"
	"github.com/mateigraura/wirebo-api/storage"
)

func main() {
	var env string
	if len(os.Args) > 1 && os.Args[1] != "" {
		env = os.Args[1]
	} else {
		env = "dev"
	}

	utils.LoadEnvFile(env)
	storage.CreateSchema()
	storage.CreateRedisClient()
	router.Run()
}
