package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/satriohtmo/go-gin-gorm.git/config"
	"github.com/satriohtmo/go-gin-gorm.git/config/app_config"
	"github.com/satriohtmo/go-gin-gorm.git/controllers/auth_controller"
	"github.com/satriohtmo/go-gin-gorm.git/controllers/user_controller"
	"github.com/satriohtmo/go-gin-gorm.git/database"
	"github.com/satriohtmo/go-gin-gorm.git/database/migrations"
)


func main()  {
	err := godotenv.Load()
	
	if err != nil{
		log.Println("error loading .env file")
	}

	config.InitConfig()
	database.ConnectDb()
	migrations.RunMigration()

	 r := gin.Default()

	 r.GET("/", user_controller.GetAllUsers)
	 r.GET("/user/:id", user_controller.UserById)
	 r.POST("/register", auth_controller.SignUp)
	 r.POST("/login", auth_controller.Login)
	 r.PUT("/user/:id", user_controller.EditUserById)
	 r.DELETE("/user/:id", user_controller.DeleteUserById)

	 r.Run(app_config.PORT)
}
