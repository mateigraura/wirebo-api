package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateigraura/wirebo-api/core/handlers"
	"github.com/mateigraura/wirebo-api/models"
	"github.com/mateigraura/wirebo-api/repository"
)

func Login(c *gin.Context) {
	var request models.LoginRequest
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseMsg(badRequestBody))
		return
	}

	handler := handlers.NewAuthHandler(&repository.UserRepositoryImpl{})
	token, err := handler.Login(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseMsg(errMessage))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func Register(c *gin.Context) {
	var request models.RegisterRequest
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseMsg(badRequestBody))
		return
	}

	handler := handlers.NewAuthHandler(&repository.UserRepositoryImpl{})
	ok, _ := handler.Register(request)
	if !ok {
		c.JSON(http.StatusInternalServerError, responseMsg(errMessage))
		return
	}

	c.JSON(http.StatusOK, okMessage)
}

func Refresh(c *gin.Context) {
	var request map[string]interface{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseMsg(badRequestBody))
		return
	}
	if request["token"] == nil {
		c.JSON(http.StatusBadRequest, responseMsg(badRequestBody))
		return
	}
	handler := handlers.NewAuthHandler(&repository.UserRepositoryImpl{})
	token, err := handler.Refresh(request["token"].(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseMsg(errMessage))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
