package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mateigraura/wirebo-api/controllers"
	"github.com/mateigraura/wirebo-api/core/utils"
	"github.com/mateigraura/wirebo-api/middleware"
	"github.com/mateigraura/wirebo-api/repository"
	"github.com/mateigraura/wirebo-api/ws"
)

func Run() {
	router := gin.Default()
	router.Use(cors())
	registerWsServer(router)
	registerAPIGroup(router)
	if err := router.Run(utils.GetEnvFile()[utils.Port]); err != nil {
		panic(err)
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
			protected.GET("/get-key/:id", controllers.GetPublicKey)
			protected.POST("/add-key", controllers.AddPublicKey)

			protected.POST("/avatar", controllers.UploadAvatar)
			protected.GET("/search/:query", controllers.Search)
			protected.GET("/user", controllers.GetUser)

			protected.GET("/rooms", controllers.GetRooms)
			protected.GET("/room-messages/:id", controllers.GetRoomMessages)
			protected.GET("/room/private/:id", controllers.GetRoom)
			protected.POST("/room/new", controllers.CreateRoom)
		}
	}
}

func registerWsServer(router *gin.Engine) {
	wsServerArgs := ws.ServerArgs{
		RoomRepository:    &repository.RoomRepositoryImpl{},
		MessageRepository: &repository.MessageRepositoryImpl{},
	}
	wsServer := ws.NewWsServer(wsServerArgs)
	go wsServer.Run()

	router.GET("/ws/:id", func(c *gin.Context) {
		ws.ServeWs(wsServer, c)
	})
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
