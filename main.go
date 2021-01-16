package main

import (
	"github.com/mateigraura/wirebo-api/storage"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mateigraura/wirebo-api/core"
	"github.com/mateigraura/wirebo-api/utils"
)

func main() {
	env := os.Args[1]
	utils.LoadEnvFile(env)

	storage.CreateSchema(storage.Connection(false))

	wsServer := core.NewWsServer()
	go wsServer.Run()

	router := gin.Default()
	router.GET("/", hello)
	router.GET("/ws/", func(c *gin.Context) {
		log.Println(c.Request)
		core.ServeWs(wsServer, c.Writer, c.Request)
	})

	if err := router.Run(utils.GetEnvFile()[utils.Port]); err != nil {
		log.Fatal(err)
	}
}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}
