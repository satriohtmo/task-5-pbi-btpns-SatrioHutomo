package database

import (
	"fmt"
	"log"

	"github.com/satriohtmo/go-gin-gorm.git/config/database_config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb()  {
	var errConnection  error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", database_config.DB_HOST, database_config.DB_USER, database_config.DB_PASSWORD, database_config.DB_NAME, database_config.DB_PORT)
    DB, errConnection = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if errConnection != nil {
		log.Fatal("error connecting to DB")
	}
	

	log.Println("connected database")


}