package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satriohtmo/go-gin-gorm.git/app"
	"github.com/satriohtmo/go-gin-gorm.git/database"
	"github.com/satriohtmo/go-gin-gorm.git/helpers"
	"github.com/satriohtmo/go-gin-gorm.git/models"
)

func GetAllUsers(c *gin.Context)  {
	var user []models.User

	database.DB.Find(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UserById(c *gin.Context)  {
	var user models.User
	id := c.Param("id")

	database.DB.First(&user, id)


	c.JSON(http.StatusOK, gin.H{"data": user})
}

func EditUserById(c *gin.Context)  {
	id := c.Param("id")
	var body app.UpdateUser
	c.Bind(&body)
	newUser := models.User{
		Username: body.Username, 
		Email: body.Email, 
		Password: body.Password,
	}

	hashPass, _ := helpers.HashPassword(newUser.Password)

	user := models.User{
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: hashPass,
	}

	if err := database.DB.Model(&user).Where("id = ?", id).Updates(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User has been updated"})
}

func DeleteUserById(c *gin.Context)  {
	id := c.Param("id")

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if err := database.DB.Unscoped().Delete(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User has been deleted"})
}