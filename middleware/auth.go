package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/satriohtmo/go-gin-gorm.git/helpers"
)

func Authentication(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unvalid token"})
		return
	}

	accessToken := strings.Split(authHeader, " ")[1]
	claims, err := helpers.ReadToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.Set("reqID", claims.ID)

	c.Next()
}