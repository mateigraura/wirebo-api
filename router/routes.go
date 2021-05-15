package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateigraura/wirebo-api/controllers"
	"github.com/mateigraura/wirebo-api/core/utils"
	"github.com/mateigraura/wirebo-api/middleware"
	"github.com/mateigraura/wirebo-api/repository"
	"github.com/mateigraura/wirebo-api/ws"
)

func Run() {
	router := gin.Default()
	registerWsServer(router)
	registerAPIGroup(router)
	if err := router.Run(utils.GetEnvFile()[utils.Port]); err != nil {
		log.Fatal(err)
	}
}

func registerAPIGroup(router *gin.Engine) {
	router.MaxMultipartMemory = 8 << 20
	api := router.Group("/api")
	{
		api.GET("/avatar/:hash", controllers.GetAvatar)
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
			protected.POST("/avatar", controllers.UploadAvatar)
			protected.GET("/search/:query", controllers.Search)
			protected.GET("/user", controllers.GetUser)
		}
	}
}

func registerWsServer(router *gin.Engine) {
	wsServer := ws.NewWsServer(&repository.RoomRepositoryImpl{})
	go wsServer.Run()

	router.GET("/ws/:id/:key", func(c *gin.Context) {
		ws.ServeWs(wsServer, c)
	})
}
