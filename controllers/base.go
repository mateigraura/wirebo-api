package controllers

import "github.com/gin-gonic/gin"

const badRequestBody = "Bad request body"

const errMessage = "Failure"
const okMessage = "Success"

var responseMsg = func(message string) gin.H {
	return gin.H{
		"message": message,
	}
}
