package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/satriohtmo/go-gin-gorm.git/config"
	"github.com/satriohtmo/go-gin-gorm.git/config/app_config"
	"github.com/satriohtmo/go-gin-gorm.git/database"
	"github.com/satriohtmo/go-gin-gorm.git/database/migrations"
	"github.com/satriohtmo/go-gin-gorm.git/router"
)


func main()  {
	err := godotenv.Load()
	
	if err != nil{
		log.Println("error loading .env file")
	}

	config.InitConfig()
	database.ConnectDb()
	migrations.RunMigration()

	router := router.Router()


	 router.Run(app_config.PORT)
}
