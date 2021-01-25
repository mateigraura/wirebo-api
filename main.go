package main

import (
	"os"

	"github.com/mateigraura/wirebo-api/router"
	"github.com/mateigraura/wirebo-api/storage"
	"github.com/mateigraura/wirebo-api/utils"
)

func main() {
	var env string
	if len(os.Args) > 0 && os.Args[1] != "" {
		env = os.Args[1]
	} else {
		env = "dev"
	}

	utils.LoadEnvFile(env)
	storage.CreateSchema()
	router.Run()
}
