package migrations

import (
	"fmt"
	"log"

	"github.com/satriohtmo/go-gin-gorm.git/database"
	"github.com/satriohtmo/go-gin-gorm.git/models"
)

func RunMigration()  {
	err := database.DB.AutoMigrate(&models.User{}, &models.Photo{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database migrated")
}