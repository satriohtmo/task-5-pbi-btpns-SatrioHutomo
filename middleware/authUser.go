package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satriohtmo/go-gin-gorm.git/database"
	"github.com/satriohtmo/go-gin-gorm.git/models"
)

func AuthUser(c *gin.Context) {
	id := c.Param("id")
	userId := c.GetUint("reqID")

	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	ID := user.ID

	if userId != ID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You dont have access"})
		return
	}

	c.Next()
}