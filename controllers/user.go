package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateigraura/wirebo-api/core/handlers"
	"github.com/mateigraura/wirebo-api/repository"
)

func Search(c *gin.Context) {
	query := c.Param("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, responseMsg(badRequestBody))
		return
	}

	userHandler := handlers.NewUserHandler(&repository.UserRepositoryImpl{})
	results, err := userHandler.Search(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseMsg(errMessage))
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": results})
}
