package photo_controller

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/satriohtmo/go-gin-gorm.git/app"
	"github.com/satriohtmo/go-gin-gorm.git/database"
	"github.com/satriohtmo/go-gin-gorm.git/models"
)

func AddNewPhoto(c *gin.Context) {
	id := c.GetUint("id")

	var body app.PhotoRequestBody
	c.Bind(&body)

	file, err := c.FormFile("photo_url")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	filepath := filepath.Join("uploads", file.Filename)
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newPhoto := models.Photo{Title: body.Title, Caption: body.Caption, PhotoURL: filepath + file.Filename, UserID: int(id)}

	if err := database.DB.Create(&newPhoto).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "New photo created"})
}

func GetPhotos(c *gin.Context) {
	var photos []models.Photo
	database.DB.Find(&photos)

	c.JSON(http.StatusOK, gin.H{"data": photos})
}

func GetPhotoById(c *gin.Context) {
	id := c.Param("id")

	var photo models.Photo
	if err := database.DB.First(&photo, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": photo})
}


func UpdatePhoto(c *gin.Context) {
	id := c.Param("id")
	reqID := c.GetUint("reqID")

	file, err := c.FormFile("photo_url")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	filepath := filepath.Join("uploads", file.Filename)
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	updatePhoto := models.Photo{
		Title:    c.PostForm("title"),
		Caption:  c.PostForm("caption"),
		PhotoURL: filepath + file.Filename,
		UserID:   int(reqID),
	}

	if err := database.DB.Model(&updatePhoto).Where("id = ?", id).Updates(&updatePhoto).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo updated"})
}

func DeletePhoto(c *gin.Context) {
	id := c.Param("id")

	var photo models.Photo
	if err := database.DB.First(&photo, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if err := database.DB.Delete(&photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted"})
}