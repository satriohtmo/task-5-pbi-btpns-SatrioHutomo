package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/satriohtmo/go-gin-gorm.git/database"
	"github.com/satriohtmo/go-gin-gorm.git/models"
)

func GetAllUsers(c *gin.Context)  {
	var user []models.User

	database.DB.Find(&user)

	c.JSON(200, gin.H{"data": user})
}

func UserById(c *gin.Context)  {
	var user models.User
	id := c.Param("id")

	database.DB.First(&user, id)


	c.JSON(200, gin.H{"data": user})
}

func EditUserById(c *gin.Context)  {
	var user models.User

	var body struct {
		Username string
		Email string
		Password string
	}

	c.Bind(&body)

	id := c.Param("id")

	database.DB.First(&user, id)

	database.DB.Model(&user).Updates(models.User{Username: body.Username, Email: body.Email, Password: body.Password})

	c.JSON(201, gin.H{"data": user})
}

func DeleteUserById(c *gin.Context)  {
	var user models.User
	id := c.Param("id")

	database.DB.Delete(&user, id)


	c.JSON(200, gin.H{"message": "User has been deleted"})
}