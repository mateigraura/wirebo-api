package controllers

import (
	"github.com/mateigraura/wirebo-api/crypto/hashing"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateigraura/wirebo-api/core/handlers"
	"github.com/mateigraura/wirebo-api/models"
	"github.com/mateigraura/wirebo-api/repository"
)

func CreateRoom(c *gin.Context) {
	var createRoomRequest models.CreateRoomRequest
	err := c.Bind(&createRoomRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseMsg(badRequestBody))
		return
	}

	if !(len(createRoomRequest.UsersRefs) >= 2) {
		c.JSON(http.StatusBadRequest, responseMsg(badRequestBody))
		return
	}

	roomHandler := handlers.NewRoomHandler(
		&repository.RoomRepositoryImpl{},
		&hashing.ShaHasher{},
	)
	createdRoom, err := roomHandler.CreateRoom(createRoomRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseMsg(errMessage))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"room": createdRoom,
	})
}

func GetRooms(c *gin.Context) {
	id, ok := c.Get("id")
	if !ok {
		c.JSON(http.StatusInternalServerError, responseMsg(errMessage))
		return
	}

	roomHandler := handlers.NewRoomHandler(&repository.RoomRepositoryImpl{}, nil)
	rooms, err := roomHandler.GetRoomsForUser(id.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseMsg(errMessage))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"rooms": rooms,
	})
}
