package main

import (
	"os"

	"github.com/mateigraura/wirebo-api/router"
	"github.com/mateigraura/wirebo-api/storage"
	"github.com/mateigraura/wirebo-api/utils"
)

func main() {
	env := os.Args[1]
	utils.LoadEnvFile(env)

	storage.CreateSchema()

	router.Run()
}
