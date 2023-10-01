package auth_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satriohtmo/go-gin-gorm.git/app"
	"github.com/satriohtmo/go-gin-gorm.git/database"
	"github.com/satriohtmo/go-gin-gorm.git/helpers"
	"github.com/satriohtmo/go-gin-gorm.git/models"
)

func SignUp(c *gin.Context)  {
	var body app.Register

	c.Bind(&body)

	var user models.User

	database.DB.First(&user, "email = ?", body.Email)

	if user.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "this email already exists",
		})
		return 
	}

	hash, err := helpers.HashPassword(body.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"message": "hasing Pass failed",
		})
		return
	}

	newUser := models.User{Username: body.Username, Email: body.Email, Password: string(hash)}

	result := database.DB.Create(&newUser)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User could not be created",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "New User created",
	})
}

func Login(c *gin.Context)  {
	var body app.Login

	c.Bind(&body)

	var user models.User
	if err := database.DB.First(&user, "email = ?", body.Email).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Invalid email"})
		return
	}

	if err := helpers.ComparePassword(body.Password, user.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid password"})
		return
	}

	accessToken, _ := helpers.GenerateToken(user.ID)

	c.JSON(http.StatusCreated, gin.H{"message": "Login Success", "userID": user.ID, "token": accessToken})
}