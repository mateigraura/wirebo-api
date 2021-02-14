package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mateigraura/wirebo-api/crypto"
)

const noBearerPresent = "No authorization bearer provided"
const incorrectBearer = "Incorrect bearer provided"
const invalidJwtToken = "Invalid or expired token"
const bearerSplitOn = "Bearer "

var returnUnauthorized = func(c *gin.Context, errMessage string) {
	c.JSON(http.StatusUnauthorized, gin.H{"message": errMessage})
	c.Abort()
	return
}

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		if bearer == "" {
			returnUnauthorized(c, noBearerPresent)
		}

		ok, token := parseBearer(bearer)
		if !ok {
			returnUnauthorized(c, incorrectBearer)
		}

		claims, err := crypto.ValidateJwt(token)
		if err != nil {
			returnUnauthorized(c, invalidJwtToken)
		} else {
			c.Set("id", claims.Id)
		}
		c.Next()
	}
}

func parseBearer(bearer string) (bool, string) {
	splitBearer := strings.Split(bearer, bearerSplitOn)

	if len(splitBearer) != 2 {
		return false, ""
	}

	return true, strings.TrimSpace(splitBearer[1])
}
