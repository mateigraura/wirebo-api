package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mateigraura/wirebo-api/controllers"
	"github.com/mateigraura/wirebo-api/middleware"
	"github.com/mateigraura/wirebo-api/utils"
	"github.com/mateigraura/wirebo-api/ws"
)

func Run() {
	wsServer := ws.NewWsServer()
	go wsServer.Run()

	router := gin.Default()

	registerAPIGroup(router)

	router.GET("/ws/", func(c *gin.Context) {
		ws.ServeWs(wsServer, c.Writer, c.Request)
	})

	if err := router.Run(utils.GetEnvFile()[utils.Port]); err != nil {
		log.Fatal(err)
	}
}

func registerAPIGroup(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/login", controllers.Login)
		api.POST("/register", controllers.Register)
		api.POST("/refresh", controllers.Refresh)

		_ = api.Group("/auth").Use(middleware.Authorization())
		{
			api.GET("/rooms")

			api.GET("/get-key", controllers.GetPublicKey)
			api.POST("/add-key", controllers.AddPubKey)
		}
	}
}
