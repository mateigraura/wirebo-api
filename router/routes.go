package router

import (
	"github.com/mateigraura/wirebo-api/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mateigraura/wirebo-api/controllers"
	"github.com/mateigraura/wirebo-api/core"
	"github.com/mateigraura/wirebo-api/utils"
)

func Run() {
	wsServer := core.NewWsServer()
	go wsServer.Run()

	router := gin.Default()

	registerAPIGroup(router)

	router.GET("/ws/", func(c *gin.Context) {
		core.ServeWs(wsServer, c.Writer, c.Request)
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

		_ = api.Group("").Use(middleware.Authorization())
		{
			api.GET("/rooms")
		}
	}
}
