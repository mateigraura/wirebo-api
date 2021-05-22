package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateigraura/wirebo-api/core/handlers"
	"github.com/mateigraura/wirebo-api/repository"
)

func GetPublicKey(c *gin.Context) {
	id := c.Param("id")
	handler := handlers.NewKeyMapHandler(&repository.KeyMapRepositoryImpl{})
	pubKey, err := handler.GetKey(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseMsg(errMessage))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"publicKey": pubKey,
	})
}

func AddPublicKey(c *gin.Context) {
	var request map[string]interface{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseMsg(badRequestBody))
		return
	}

	id, ok := c.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, responseMsg(badRequestBody))
		return
	}

	handler := handlers.NewKeyMapHandler(&repository.KeyMapRepositoryImpl{})
	pubKey, err := handler.InsertKey(id.(string), request["publicKey"].(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseMsg(errMessage))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"publicKey": pubKey,
	})
}
