package router

import (
	"github.com/gin-gonic/gin"
	"github.com/satriohtmo/go-gin-gorm.git/controllers/auth_controller"
	"github.com/satriohtmo/go-gin-gorm.git/controllers/photo_controller"
	"github.com/satriohtmo/go-gin-gorm.git/controllers/user_controller"
	"github.com/satriohtmo/go-gin-gorm.git/middleware"
)

func Router() *gin.Engine {
	router := gin.Default()

	 router.POST("/register", auth_controller.SignUp)
	 router.POST("/login", auth_controller.Login)

	 router.Use(middleware.Authentication)

	 users := router.Group("/api/users")
	{
		users.GET("/", user_controller.GetAllUsers)
		users.GET("/:id", user_controller.UserById)
		users.PUT("/:id", middleware.AuthUser, user_controller.EditUserById)
		users.DELETE("/:id", middleware.AuthUser, user_controller.DeleteUserById)
	}

	photos := router.Group("/api/photos")
	{
		photos.GET("/", photo_controller.GetPhotos)
		photos.POST("/", photo_controller.AddNewPhoto)
		photos.GET("/:id", photo_controller.GetPhotoById)
		photos.PUT("/:id", photo_controller.UpdatePhoto)
		photos.DELETE("/:id", photo_controller.DeletePhoto)
	}

	 return router
}