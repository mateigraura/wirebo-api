package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mateigraura/wirebo-api/crypto/authorization"
)

const (
	noBearerPresent = "No authorization bearer provided"
	incorrectBearer = "Incorrect bearer provided"
	invalidJwtToken = "Invalid or expired token"
	bearerSplitOn   = "Bearer "
)

var returnUnauthorized = func(c *gin.Context, errMessage string) {
	c.JSON(http.StatusUnauthorized, gin.H{"message": errMessage})
}

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		if bearer == "" {
			returnUnauthorized(c, noBearerPresent)
			c.Abort()
			return
		}

		ok, token := parseBearer(bearer)
		if !ok {
			returnUnauthorized(c, incorrectBearer)
			c.Abort()
			return
		}

		claims, err := authorization.ValidateJwt(token)
		if err != nil {
			returnUnauthorized(c, invalidJwtToken)
			c.Abort()
			return
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
