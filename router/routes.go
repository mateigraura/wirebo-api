package router

import (
	"log"
	"net/http"

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
		auth := api.Group("/auth")
		{
			auth.POST("/login", controllers.Login)
			auth.POST("/register", controllers.Register)
			auth.POST("/refresh", controllers.Refresh)
		}

		protected := api.Group("/p").Use(middleware.Authorization())
		{
			protected.GET("/rooms", func(c *gin.Context) {
				id, ok := c.Get("id")
				if !ok {
					c.JSON(http.StatusInternalServerError, "Failure")
				}
				c.JSON(http.StatusOK, id)
			})

			protected.GET("/get-key", controllers.GetPublicKey)
			protected.POST("/add-key", controllers.AddPublicKey)
		}
	}
}
