package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateigraura/wirebo-api/core/handlers"
	"github.com/mateigraura/wirebo-api/crypto/hashing"
	"github.com/mateigraura/wirebo-api/repository"
)

const (
	hashLength       = 64
	contentTypeImage = "image/%s"
)

func UploadAvatar(c *gin.Context) {
	form, err := c.FormFile("imgFile")
	if err != nil {
		c.JSON(http.StatusBadRequest, responseMsg(badRequestBody))
		return
	}

	file, err := form.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, responseMsg(badRequestBody))
		return
	}
	id, ok := c.Get("id")
	if !ok {
		c.JSON(http.StatusInternalServerError, responseMsg(errMessage))
		return
	}

	avatarHandler := handlers.NewAvatarHandler(
		&repository.AvatarRepository{},
		&repository.UserRepositoryImpl{},
		&hashing.ShaHasher{},
	)
	hash, err := avatarHandler.Save(file, id.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseMsg(errMessage))
		return
	}

	c.JSON(http.StatusOK, gin.H{"hash": hash})
}

func GetAvatar(c *gin.Context) {
	hash := c.Param("hash")
	if hash == "" || len([]byte(hash)) != hashLength {
		c.JSON(http.StatusBadRequest, responseMsg(badRequestBody))
		return
	}
	avatarHandler := handlers.NewAvatarHandler(
		&repository.AvatarRepository{},
		nil,
		nil,
	)
	content, imgType, err := avatarHandler.GetByHash(hash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseMsg(errMessage))
		return
	}

	c.Data(http.StatusOK, fmt.Sprintf(contentTypeImage, imgType), content)
}
