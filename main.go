package main

import (
	"github.com/mateigraura/wirebo-api/logging"
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

	logging.InitLoggers()
	utils.LoadEnvFile(env)
	storage.CreateSchema()
	storage.CreateRedisClient()
	router.Run()
}
